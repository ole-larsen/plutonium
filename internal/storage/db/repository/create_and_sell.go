package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/models"
)

type CreateAndSell struct {
	Created     strfmt.Date `db:"created"`
	Updated     strfmt.Date `db:"updated"`
	Deleted     strfmt.Date `db:"deleted"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	Link        string      `db:"link"`
	ID          int64       `db:"id"`
	ImageID     int64       `db:"image_id"`
	OrderBy     int64       `db:"order_by"`
	CreatedBy   int64       `db:"created_by"`
	UpdatedBy   int64       `db:"updated_by"`
	Enabled     bool        `db:"enabled"`
}

type CreateAndSellRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, createAndSellMap map[string]interface{}) error
	Update(ctx context.Context, createAndSellMap map[string]interface{}) ([]*models.CreateAndSell, error)
	GetCreateAndSell(ctx context.Context) ([]*models.CreateAndSell, error)
	GetCreateAndSellByID(ctx context.Context, id int64) (*models.CreateAndSell, error)
	GetPublicCreateAndSell(ctx context.Context) ([]*models.PublicCreateAndSellItem, error)
}

// CreateAndSellRepository - repository to store users.
type CreateAndSellRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewCreateAndSellRepository(db *sqlx.DB, tbl string) *CreateAndSellRepository {
	if db == nil {
		return nil
	}

	return &CreateAndSellRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *CreateAndSellRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *CreateAndSellRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *CreateAndSellRepository) Create(ctx context.Context, createAndSellMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO create_and_sell (title, description, image_id, link, enabled, order_by, created_by_id, updated_by_id)
VALUES (:title, :description, :image_id, :link, :enabled, :order_by, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING`, createAndSellMap)

	return err
}

func (r *CreateAndSellRepository) Update(ctx context.Context, createAndSellMap map[string]interface{}) ([]*models.CreateAndSell, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx,
		`UPDATE create_and_sell SET
	title=:title,
	link=:link,
	description=:description,
	image_id=:image_id,
	enabled=:enabled,
	order_by=:order_by,
	updated_by_id=:updated_by_id WHERE id =:id`, createAndSellMap)
	if err != nil {
		return nil, err
	}

	return r.GetCreateAndSell(ctx)
}

func (r *CreateAndSellRepository) GetCreateAndSell(ctx context.Context) ([]*models.CreateAndSell, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		items []*models.CreateAndSell
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT 
	id, 
	title, 
	link,
	description, 
	image_id, 
	enabled, 
	order_by, 
	created_by_id, 
	updated_by_id
	FROM create_and_sell;`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item CreateAndSell
		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Link,
			&item.Description,
			&item.ImageID,
			&item.Enabled,
			&item.OrderBy,
			&item.CreatedBy,
			&item.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, &models.CreateAndSell{
			ID:          item.ID,
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			ImageID:     item.ImageID,
			OrderBy:     item.OrderBy,
			Enabled:     item.Enabled,
			CreatedByID: item.CreatedBy,
			UpdatedByID: item.UpdatedBy,
		})
	}

	defer rows.Close()

	return items, nil
}

func (r *CreateAndSellRepository) GetCreateAndSellByID(ctx context.Context, id int64) (*models.CreateAndSell, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var item CreateAndSell

	sqlStatement :=
		`SELECT 
	c.id, 
	c.title, 
	c.description, 
	c.link,
	c.image_id, 
	c.enabled, 
	c.order_by, 
	c.created_by_id, 
	c.updated_by_id
FROM create_and_sell c where c.id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	if err := row.Scan(&item.ID, &item.Title, &item.Description, &item.Link, &item.ImageID,
		&item.Enabled, &item.OrderBy, &item.CreatedBy, &item.UpdatedBy); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("create_and_sell not found")
		}

		return nil, NewError(err)
	}

	return &models.CreateAndSell{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		ImageID:     item.ImageID,
		OrderBy:     item.OrderBy,
		Enabled:     item.Enabled,
		CreatedByID: item.CreatedBy,
		UpdatedByID: item.UpdatedBy,
	}, nil
}

func (r *CreateAndSellRepository) GetPublicCreateAndSell(ctx context.Context) ([]*models.PublicCreateAndSellItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		items    []*models.PublicCreateAndSellItem
	)

	rows, err := r.DB.QueryxContext(ctx, `
SELECT c.id, 
	c.title, 
	c.description, 
	c.link,
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
						) FROM files f WHERE f.id = c.image_id)
		) FROM files f WHERE f.id = c.image_id) as image
FROM create_and_sell c
WHERE 
	c.enabled = true AND c.deleted isNULL  GROUP BY c.id ORDER BY c.order_by ASC;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item CreateAndSell

		var image AggregatedImageJSON

		err = rows.Scan(&item.ID, &item.Title, &item.Description, &item.Link, &image)
		if err != nil {
			return nil, err
		}

		items = append(items, &models.PublicCreateAndSellItem{
			ID: item.ID,
			Attributes: &models.PublicCreateAndSellItemAttributes{
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
				Image: &models.PublicFile{
					Attributes: image.Attributes,
					ID:         image.ID,
				},
			},
		})
	}

	defer rows.Close()

	return items, multierr.ErrorOrNil()
}
