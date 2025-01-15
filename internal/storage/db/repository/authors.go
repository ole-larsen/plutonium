package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ole-larsen/plutonium/models"
)

type Social struct {
	ID       int64  `db:"id"`
	AuthorID int64  `db:"author_id"`
	Name     string `db:"name"`
	Link     string `db:"link"`
	Icon     string `db:"icon"`
}
type Author struct {
	ID          int64       `db:"id"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	Name        string      `db:"name"`
	Slug        string      `db:"slug"`
	ImageID     int64       `db:"image_id"`
	Enabled     bool        `db:"enabled"`
	OrderBy     int64       `db:"order_by"`
	CreatedBy   int64       `db:"created_by"`
	UpdatedBy   int64       `db:"updated_by"`
	Created     strfmt.Date `db:"created"`
	Updated     strfmt.Date `db:"updated"`
	Deleted     strfmt.Date `db:"deleted"`
}

type AuthorsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, authorMap map[string]interface{}, socials []*models.Social, wallets []*models.Wallet) error
	Update(ctx context.Context, authorMap map[string]interface{}, socials []*models.Social, wallets []*models.Wallet) ([]*models.Author, error)
	GetAuthors(ctx context.Context) ([]*models.Author, error)
	GetAuthorByID(ctx context.Context, id int64) (*models.Author, error)
	GetPublicAuthors(ctx context.Context) ([]*models.PublicAuthorItem, error)
	GetPublicAuthor(ctx context.Context, slug string) (*models.PublicAuthorItem, error)
	BindWallet(ctx context.Context, walletMap map[string]interface{}) error
	BindSocial(ctx context.Context, socialMap map[string]interface{}) error
}

// AuthorsRepository - repository to store frontendauthors.
type AuthorsRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewAuthorsRepository(db *sqlx.DB, tbl string) *AuthorsRepository {
	if db == nil {
		return nil
	}

	return &AuthorsRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *AuthorsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *AuthorsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *AuthorsRepository) Create(ctx context.Context, authorMap map[string]interface{}, socials []*models.Social, wallets []*models.Wallet) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO authors (title, description, name, slug,
		                           image_id, enabled, order_by, created_by_id, updated_by_id)
		VALUES (:title, :description, :name, :slug,
		        :image_id, :enabled, :order_by, :created_by_id, :updated_by_id)
		ON CONFLICT DO NOTHING`, authorMap)

	if socials != nil {
		var lastInsertId int64
		sqlStatement := `SELECT id FROM authors WHERE slug=$1;`
		row := r.DB.QueryRow(sqlStatement, authorMap["slug"])
		err = row.Scan(&lastInsertId)
		if err != nil {
			return err
		}
		authorMap["id"] = lastInsertId
		for _, social := range socials {
			socialMap := make(map[string]interface{})
			socialMap["author_id"] = authorMap["id"]
			socialMap["name"] = social.Name
			socialMap["link"] = social.Link
			socialMap["icon"] = social.Icon
			socialMap["updated_by_id"] = authorMap["updated_by_id"]
			socialMap["created_by_id"] = authorMap["created_by_id"]

			err = r.BindSocial(ctx, socialMap)
			if err != nil {
				return err
			}
		}
		for _, wallet := range wallets {
			walletMap := make(map[string]interface{})
			walletMap["author_id"] = authorMap["id"]
			walletMap["name"] = wallet.Name
			walletMap["address"] = wallet.Address
			walletMap["updated_by_id"] = authorMap["updated_by_id"]
			walletMap["created_by_id"] = authorMap["created_by_id"]

			err = r.BindWallet(ctx, walletMap)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (r *AuthorsRepository) Update(ctx context.Context, authorMap map[string]interface{}, socials []*models.Social, wallets []*models.Wallet) ([]*models.Author, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}
	_, err := r.DB.NamedExecContext(ctx, `UPDATE authors SET
                title=:title,
                description=:description,
                name=:name,
                slug=:slug,
                image_id=:image_id,
                enabled=:enabled,
                order_by=:order_by,
                updated_by_id=:updated_by_id WHERE id =:id`, authorMap)
	if err != nil {
		return nil, err
	}

	if socials != nil {
		for _, social := range socials {
			socialMap := make(map[string]interface{})
			socialMap["author_id"] = authorMap["id"]
			socialMap["name"] = social.Name
			socialMap["link"] = social.Link
			socialMap["icon"] = social.Icon
			socialMap["updated_by_id"] = authorMap["updated_by_id"]
			socialMap["created_by_id"] = authorMap["created_by_id"]

			err = r.BindSocial(ctx, socialMap)
			if err != nil {
				return nil, err
			}
		}
	}
	if wallets != nil {
		for _, wallet := range wallets {
			walletMap := make(map[string]interface{})
			walletMap["author_id"] = authorMap["id"]
			walletMap["name"] = wallet.Name
			walletMap["address"] = wallet.Address
			walletMap["updated_by_id"] = authorMap["updated_by_id"]
			walletMap["created_by_id"] = authorMap["created_by_id"]

			err = r.BindWallet(ctx, walletMap)
			if err != nil {
				return nil, err
			}
		}
	}
	return r.GetAuthors(ctx)
}

func (r *AuthorsRepository) GetAuthors(ctx context.Context) ([]*models.Author, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}
	var (
		multierr multierror.Error
		authors  []*models.Author
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT 
    		id, 
			title, 
			description, 
			name, 
			slug, 
			image_id, 
			enabled, 
			order_by, 
			created_by_id, 
			updated_by_id,
			(SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
				'name', s.name,
				'link', s.link,
			    'icon', s.icon
			))) FROM authors_socials s WHERE s.author_id = a.id) as socials,
    		(SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
				'name', w.name,
				'address', w.address
			))) FROM authors_wallets w WHERE w.author_id = a.id) as wallets
			FROM authors a;`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author Author
		var socials AggregatedSocial
		var wallets AggregatedWallet
		err = rows.Scan(
			&author.ID,
			&author.Title,
			&author.Description,
			&author.Name,
			&author.Slug,
			&author.ImageID,
			&author.Enabled,
			&author.OrderBy,
			&author.CreatedBy,
			&author.UpdatedBy,
			&socials,
			&wallets)

		if err != nil {
			return nil, err
		}
		authorSocials := make([]*models.Social, 0)
		for _, social := range socials {
			authorSocials = append(authorSocials, &models.Social{
				Link: social.Link,
				Name: social.Name,
				Icon: social.Icon,
			})
		}
		authorWallets := make([]*models.Wallet, 0)
		for _, wallet := range wallets {
			authorWallets = append(authorWallets, &models.Wallet{
				Name:    wallet.Name,
				Address: wallet.Address,
			})
		}
		authors = append(authors, &models.Author{
			ID:          author.ID,
			Title:       author.Title,
			Description: author.Description,
			Name:        author.Name,
			Slug:        author.Slug,
			ImageID:     author.ImageID,
			OrderBy:     author.OrderBy,
			Enabled:     author.Enabled,
			CreatedByID: author.CreatedBy,
			UpdatedByID: author.UpdatedBy,
			Socials:     authorSocials,
			Wallets:     authorWallets,
		})
	}
	defer rows.Close()

	return authors, multierr.ErrorOrNil()
}

func (r *AuthorsRepository) GetAuthorByID(ctx context.Context, id int64) (*models.Author, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}
	var author Author
	var socials AggregatedSocial
	var wallets AggregatedWallet
	sqlStatement := `SELECT 
    	a.id, 
    	a.title, 
    	a.description, 
    	a.name, 
    	a.slug, 
    	a.image_id, 
    	a.enabled, 
    	a.order_by, 
    	a.created_by_id, 
    	a.updated_by_id,
    	(SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
				'name', s.name,
				'link', s.link,
    	        'icon', s.icon
			))) FROM authors_socials s WHERE s.author_id = a.id) as socials,
    	(SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
				'name', w.name,
				'address', w.address
		))) FROM authors_wallets w WHERE w.author_id = a.id) as wallets
	FROM authors a where a.id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)
	err := row.Scan(&author.ID, &author.Title, &author.Description,
		&author.Name, &author.Slug, &author.ImageID, &author.Enabled, &author.OrderBy, &author.CreatedBy, &author.UpdatedBy, &socials, &wallets)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("slider author not found")
	case nil:
		authorSocials := make([]*models.Social, 0)
		for _, social := range socials {
			authorSocials = append(authorSocials, &models.Social{
				Link: social.Link,
				Name: social.Name,
				Icon: social.Icon,
			})
		}
		authorWallets := make([]*models.Wallet, 0)
		for _, wallet := range wallets {
			authorWallets = append(authorWallets, &models.Wallet{
				Name:    wallet.Name,
				Address: wallet.Address,
			})
		}
		return &models.Author{
			ID:          author.ID,
			Title:       author.Title,
			Description: author.Description,
			Name:        author.Name,
			Slug:        author.Slug,
			ImageID:     author.ImageID,
			OrderBy:     author.OrderBy,
			Enabled:     author.Enabled,
			CreatedByID: author.CreatedBy,
			UpdatedByID: author.UpdatedBy,
			Socials:     authorSocials,
			Wallets:     authorWallets,
		}, err
	default:
		return nil, err
	}
}

func (r *AuthorsRepository) GetPublicAuthors(ctx context.Context) ([]*models.PublicAuthorItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}
	var (
		multierr multierror.Error
		authors  []*models.PublicAuthorItem
	)

	rows, err := r.DB.QueryxContext(ctx, `
			SELECT a.id, 
			       a.title, 
			       a.description, 
			       a.name, 
			       a.slug, 
			       (SELECT JSON_BUILD_OBJECT(
						'id', f.id,
						'attributes', (SELECT JSON_BUILD_OBJECT(
										'name',            f.name,
										'alt',             f.alt,
										'caption',         f.caption,
										'ext',             f.ext,
										'provider',        f.provider,
										'width',           f.width,
										'height',          f.height,
										'size',            f.size,
										'url',             f.url
									  ) FROM files f WHERE f.id = a.image_id)
					  ) FROM files f WHERE f.id = a.image_id) as image,
			       (SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
					    'name', s.name,
				        'link', s.link,
			            'icon', s.icon
			       ))) FROM authors_socials s WHERE s.author_id = a.id) as socials,
			       (SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
						'name', w.name,
						'address', w.address
					))) FROM authors_wallets w WHERE w.author_id = a.id) as wallets
		    FROM authors a
		    WHERE 
			    a.enabled = true AND a.deleted isNULL  GROUP BY a.id ORDER BY a.order_by ASC;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author Author
		var image AggregatedImageJSON
		var socials AggregatedSocial
		var wallets AggregatedWallet
		err = rows.Scan(&author.ID, &author.Title, &author.Description,
			&author.Name, &author.Slug, &image, &socials, &wallets)
		if err != nil {
			return nil, err
		}

		authorSocials := make([]*models.PublicSocial, 0)
		for _, social := range socials {
			authorSocials = append(authorSocials, &models.PublicSocial{
				Link: social.Link,
				Name: social.Name,
				Icon: social.Icon,
			})
		}

		authorWallets := make([]*models.PublicWallet, 0)
		for _, wallet := range wallets {
			authorWallets = append(authorWallets, &models.PublicWallet{
				Name:    wallet.Name,
				Address: wallet.Address,
			})
		}

		authors = append(authors, &models.PublicAuthorItem{
			ID:          author.ID,
			Title:       author.Title,
			Description: author.Description,
			Name:        author.Name,
			Link:        author.Slug,
			Image: &models.PublicFile{
				Attributes: image.Attributes,
				ID:         image.ID,
			},
			Socials: authorSocials,
			Wallets: authorWallets,
		})
	}
	defer rows.Close()

	return authors, multierr.ErrorOrNil()
}

func (r *AuthorsRepository) GetPublicAuthor(ctx context.Context, slug string) (*models.PublicAuthorItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	sqlStatement := `SELECT a.id, 
			       a.title, 
			       a.description, 
			       a.name, 
			       a.slug, 
			       (SELECT JSON_BUILD_OBJECT(
						'id', f.id,
						'attributes', (SELECT JSON_BUILD_OBJECT(
										'name',            f.name,
										'alt',             f.alt,
										'caption',         f.caption,
										'ext',             f.ext,
										'provider',        f.provider,
										'width',           f.width,
										'height',          f.height,
										'size',            f.size,
										'url',             f.url
									  ) FROM files f WHERE f.id = a.image_id)
					  ) FROM files f WHERE f.id = a.image_id) as image,
    			   (SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
					    'name', s.name,
				        'link', s.link,
    			        'icon', s.icon
			       ))) FROM authors_socials s WHERE s.author_id = a.id) as socials,
    			   (SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
						'name', w.name,
						'address', w.address
					))) FROM authors_wallets w WHERE w.author_id = a.id) as wallets
		    FROM authors a
		    WHERE 
		        a.slug = $1 AND
			    a.enabled = true AND a.deleted isNULL  GROUP BY a.id ORDER BY a.order_by ASC;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, slug)
	var author Author
	var image AggregatedImageJSON
	var socials AggregatedSocial
	var wallets AggregatedWallet
	err := row.Scan(&author.ID, &author.Title, &author.Description,
		&author.Name, &author.Slug, &image, &socials, &wallets)

	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("author not found")
	case nil:
		authorSocials := make([]*models.PublicSocial, 0)
		for _, social := range socials {
			authorSocials = append(authorSocials, &models.PublicSocial{
				Link: social.Link,
				Name: social.Name,
				Icon: social.Icon,
			})
		}
		authorWallets := make([]*models.PublicWallet, 0)
		for _, wallet := range wallets {
			authorWallets = append(authorWallets, &models.PublicWallet{
				Name:    wallet.Name,
				Address: wallet.Address,
			})
		}
		return &models.PublicAuthorItem{
			ID:          author.ID,
			Title:       author.Title,
			Description: author.Description,
			Name:        author.Name,
			Link:        author.Slug,
			Image: &models.PublicFile{
				Attributes: image.Attributes,
				ID:         image.ID,
			},
			Socials: authorSocials,
			Wallets: authorWallets,
		}, err
	default:
		return nil, err
	}
}

func (r *AuthorsRepository) BindSocial(ctx context.Context, socialMap map[string]interface{}) error {
	var lastInsertId int64
	sqlStatement := `SELECT a.id FROM authors_socials a
		    WHERE a.author_id = $1 AND a.name = $2;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, socialMap["author_id"], socialMap["name"])

	err := row.Scan(&lastInsertId)

	switch err {
	case sql.ErrNoRows:
		_, err = r.DB.NamedExecContext(ctx, `
					INSERT INTO authors_socials (author_id, name, link, icon, created_by_id, updated_by_id)
					VALUES (:author_id, :name, :link, :icon, :created_by_id, :updated_by_id)
					ON CONFLICT DO NOTHING`, socialMap)
		if err != nil {
			return err
		}
	case nil:
		_, err = r.DB.NamedExecContext(ctx, `UPDATE authors_socials SET
				   	name=:name,
				   	link=:link,
				   	icon=:icon,
				   	updated_by_id=:updated_by_id WHERE author_id=:author_id AND name=:name`, socialMap)
		if err != nil {
			return err
		}
	}
	return err
}

func (r *AuthorsRepository) BindWallet(ctx context.Context, walletMap map[string]interface{}) error {
	var lastInsertId int64
	sqlStatement := `SELECT a.id FROM authors_wallets a
		    WHERE a.author_id = $1 AND a.name = $2;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, walletMap["author_id"], walletMap["name"])

	err := row.Scan(&lastInsertId)

	switch err {
	case sql.ErrNoRows:
		_, err = r.DB.NamedExecContext(ctx, `
					INSERT INTO authors_wallets(author_id, name, address, created_by_id, updated_by_id)
					VALUES (:author_id, :name, :address, :created_by_id, :updated_by_id)
					ON CONFLICT DO NOTHING`, walletMap)
		if err != nil {
			return err
		}
	case nil:
		_, err = r.DB.NamedExecContext(ctx, `UPDATE authors_wallets SET
				   	name=:name,
				   	address=:address,
				   	updated_by_id=:updated_by_id WHERE author_id=:author_id AND name=:name`, walletMap)
		if err != nil {
			return err
		}
	}
	return err
}
