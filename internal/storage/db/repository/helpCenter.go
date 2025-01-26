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

type HelpCenter struct {
	Created     strfmt.Date `db:"created"`
	Deleted     strfmt.Date `db:"deleted"`
	Updated     strfmt.Date `db:"updated"`
	Title       string      `db:"title"`
	Slug        string      `db:"slug"`
	Link        string      `db:"link"`
	Description string      `db:"description"`
	UpdatedBy   int64       `db:"updated_by"`
	CreatedBy   int64       `db:"created_by"`
	ID          int64       `db:"id"`
	OrderBy     int64       `db:"order_by"`
	ImageID     int64       `db:"image_id"`
	Enabled     bool        `db:"enabled"`
}

type HelpCenterRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, helpCenterMap map[string]interface{}) error
	Update(ctx context.Context, helpCenterMap map[string]interface{}) ([]*models.HelpCenter, error)
	GetHelpCenter(ctx context.Context) ([]*models.HelpCenter, error)
	GetHelpCenterByID(ctx context.Context, id int64) (*models.HelpCenter, error)
	GetPublicHelpCenter(ctx context.Context) ([]*models.PublicHelpCenterItem, error)
}

// HelpCenterRepository - repository to store users.
type HelpCenterRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewHelpCenterRepository(db *sqlx.DB, tbl string) *HelpCenterRepository {
	if db == nil {
		return nil
	}

	return &HelpCenterRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *HelpCenterRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *HelpCenterRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *HelpCenterRepository) Create(ctx context.Context, helpCenterMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO help_center (title, slug, description, image_id, enabled, order_by, created_by_id, updated_by_id)
		VALUES (:title, :slug, :description,:image_id, :enabled, :order_by, :created_by_id, :updated_by_id)
		ON CONFLICT DO NOTHING`, helpCenterMap)

	return err
}

func (r *HelpCenterRepository) Update(ctx context.Context, helpCenterMap map[string]interface{}) ([]*models.HelpCenter, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE help_center SET
                title=:title,
                slug=:slug,
                description=:description,
                image_id=:image_id,
                enabled=:enabled,
                order_by:=order_by,
                updated_by_id=:updated_by_id WHERE id =:id`, helpCenterMap)
	if err != nil {
		return nil, err
	}

	return r.GetHelpCenter(ctx)
}

func (r *HelpCenterRepository) GetHelpCenter(ctx context.Context) ([]*models.HelpCenter, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr    multierror.Error
		helpCenters []*models.HelpCenter
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT id, title, slug, description, image_id, enabled, order_by, created_by_id, updated_by_id from help_center;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var helpCenter HelpCenter

		err = rows.Scan(&helpCenter.ID, &helpCenter.Title, &helpCenter.Slug, &helpCenter.Description,
			&helpCenter.ImageID, &helpCenter.Enabled, &helpCenter.OrderBy, &helpCenter.CreatedBy, &helpCenter.UpdatedBy)
		if err != nil {
			return nil, err
		}

		helpCenters = append(helpCenters, &models.HelpCenter{
			ID:          helpCenter.ID,
			Title:       helpCenter.Title,
			Slug:        helpCenter.Slug,
			Description: helpCenter.Description,
			ImageID:     helpCenter.ImageID,
			Enabled:     helpCenter.Enabled,
			OrderBy:     helpCenter.OrderBy,
			CreatedByID: helpCenter.CreatedBy,
			UpdatedByID: helpCenter.UpdatedBy,
		})
	}

	defer rows.Close()

	return helpCenters, multierr.ErrorOrNil()
}

func (r *HelpCenterRepository) GetHelpCenterByID(ctx context.Context, id int64) (*models.HelpCenter, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var helpCenter HelpCenter

	sqlStatement := `SELECT id, title, slug, description, image_id, enabled, order_by, created_by_id, updated_by_id from help_center where id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&helpCenter.ID, &helpCenter.Title, &helpCenter.Slug, &helpCenter.Description,
		&helpCenter.ImageID, &helpCenter.Enabled, &helpCenter.OrderBy, &helpCenter.CreatedBy, &helpCenter.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("helpCenter not found"))
		}

		return nil, NewError(err)
	}

	return &models.HelpCenter{
		ID:          helpCenter.ID,
		Title:       helpCenter.Title,
		Slug:        helpCenter.Slug,
		Description: helpCenter.Description,
		ImageID:     helpCenter.ImageID,
		Enabled:     helpCenter.Enabled,
		OrderBy:     helpCenter.OrderBy,
		CreatedByID: helpCenter.CreatedBy,
		UpdatedByID: helpCenter.UpdatedBy,
	}, nil
}

func (r *HelpCenterRepository) GetPublicHelpCenter(ctx context.Context) ([]*models.PublicHelpCenterItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr    multierror.Error
		helpCenters []*models.PublicHelpCenterItem
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT 
    h.id, 
    h.title, 
    h.slug as link, 
    h.description, 
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
					  ) FROM files f WHERE f.id = h.image_id)
	) FROM files f WHERE f.id = h.image_id) as image
    FROM help_center h WHERE 
	h.enabled = true AND h.deleted isNULL  GROUP BY h.id ORDER BY order_by ASC;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var helpCenter HelpCenter

		var image AggregatedImageJSON

		err = rows.Scan(&helpCenter.ID, &helpCenter.Title, &helpCenter.Link, &helpCenter.Description, &image)
		if err != nil {
			return nil, err
		}

		helpCenters = append(helpCenters, &models.PublicHelpCenterItem{
			ID:          helpCenter.ID,
			Title:       helpCenter.Title,
			Link:        helpCenter.Link,
			Description: helpCenter.Description,
			Image: &models.PublicFile{
				Attributes: image.Attributes,
				ID:         image.ID,
			},
		})
	}

	defer rows.Close()

	return helpCenters, multierr.ErrorOrNil()
}
