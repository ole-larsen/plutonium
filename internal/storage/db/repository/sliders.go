package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/models"
)

type Slider struct {
	Created     strfmt.Date `db:"created"`
	Updated     strfmt.Date `db:"updated"`
	Deleted     strfmt.Date `db:"deleted"`
	Provider    string      `db:"provider"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	ID          int64       `db:"id"`
	CreatedBy   int64       `db:"created_by"`
	UpdatedBy   int64       `db:"updated_by"`
	Enabled     bool        `db:"enabled"`
}

type SlidersRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, sliderMap map[string]interface{}) error
	Update(ctx context.Context, sliderMap map[string]interface{}) ([]*models.Slider, error)
	GetSliders(ctx context.Context) ([]*models.Slider, error)
	GetSliderByTitle(ctx context.Context, title string) (*models.Slider, error)
	GetSliderByID(ctx context.Context, id int64) (*models.Slider, error)
	GetSliderByProvider(ctx context.Context, provider string) (*models.PublicSlider, error)
}

// SlidersRepository - repository to store frontend sliders.
type SlidersRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewSlidersRepository(db *sqlx.DB, tbl string) *SlidersRepository {
	if db == nil {
		return nil
	}

	return &SlidersRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *SlidersRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *SlidersRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *SlidersRepository) Create(ctx context.Context, sliderMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO sliders (provider, title, description, enabled, created_by_id, updated_by_id)
VALUES (:provider, :title, :description, :enabled, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING`, sliderMap)

	return err
}

func (r *SlidersRepository) Update(ctx context.Context, sliderMap map[string]interface{}) ([]*models.Slider, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
UPDATE sliders SET
	provider=:provider,
	title=:title,
	description=:description,
	enabled=:enabled,
	updated_by_id=:updated_by_id 
WHERE id =:id`, sliderMap)

	if err != nil {
		return nil, err
	}

	return r.GetSliders(ctx)
}

func (r *SlidersRepository) GetSliders(ctx context.Context) ([]*models.Slider, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		sliders []*models.Slider
	)

	rows, err := r.DB.QueryxContext(ctx, `SELECT id, provider, title, description, enabled, created_by_id, updated_by_id from sliders;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var slider Slider

		err = rows.Scan(&slider.ID, &slider.Provider, &slider.Title, &slider.Description, &slider.Enabled, &slider.CreatedBy, &slider.UpdatedBy)
		if err != nil {
			return nil, err
		}

		sliders = append(sliders, &models.Slider{
			ID:          slider.ID,
			Provider:    slider.Provider,
			Title:       slider.Title,
			Description: slider.Description,
			Enabled:     slider.Enabled,
			CreatedByID: slider.CreatedBy,
			UpdatedByID: slider.UpdatedBy,
		})
	}

	defer rows.Close()

	return sliders, nil
}

func (r *SlidersRepository) GetSliderByTitle(ctx context.Context, title string) (*models.Slider, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var slider Slider

	sqlStatement := `SELECT id, provider, title, description, enabled, created_by_id, updated_by_id from sliders where title=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, title)

	err := row.Scan(&slider.ID, &slider.Provider, &slider.Title, &slider.Description, &slider.Enabled, &slider.CreatedBy, &slider.UpdatedBy)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("slider not found")
	case nil:
		return &models.Slider{
			ID:          slider.ID,
			Provider:    slider.Provider,
			Title:       slider.Title,
			Description: slider.Description,
			Enabled:     slider.Enabled,
			CreatedByID: slider.CreatedBy,
			UpdatedByID: slider.UpdatedBy,
		}, nil
	default:
		return nil, err
	}
}

func (r *SlidersRepository) GetSliderByID(ctx context.Context, id int64) (*models.Slider, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var slider Slider

	sqlStatement := "SELECT id, provider, title, description, enabled, created_by_id, updated_by_id FROM sliders WHERE id=$1;"
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&slider.ID, &slider.Provider, &slider.Title, &slider.Description, &slider.Enabled, &slider.CreatedBy, &slider.UpdatedBy)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("slider not found")
	case nil:
		return &models.Slider{
			ID:          slider.ID,
			Provider:    slider.Provider,
			Title:       slider.Title,
			Description: slider.Description,
			Enabled:     slider.Enabled,
			CreatedByID: slider.CreatedBy,
			UpdatedByID: slider.UpdatedBy,
		}, nil
	default:
		return nil, err
	}
}

func (r *SlidersRepository) GetSliderByProvider(ctx context.Context, provider string) (*models.PublicSlider, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		slider      models.PublicSlider
		sliderItems AggregatedSliderItemJSON
	)
	// home-01
	sqlStatement := `
SELECT
	s.id,
	(SELECT JSON_AGG(JSON_BUILD_OBJECT(
		'heading',     i.heading,
		'description', i.description,
		'btnLink1',    i.btn_link_1,
		'btnText1',    i.btn_text_1,
		'btnLink2',    i.btn_link_2,
		'btnText2',    i.btn_text_2,
		'image',       (SELECT JSON_BUILD_OBJECT(
			'id', f.id,
			'attributes', (SELECT JSON_BUILD_OBJECT(
				'name',            f.name,
				'alt', f.alt,
				'caption',         f.caption,
				'ext',             f.ext,
				'provider',        f.provider,
				'width',           f.width,
				'height',          f.height,
				'size',            f.size,
				'url',             f.url
				) FROM files f WHERE f.id = i.image_id)
			) FROM files f WHERE f.id = i.image_id),
		'bg',          (SELECT JSON_BUILD_OBJECT(
			'id', f.id,
			'attributes', (SELECT JSON_BUILD_OBJECT(
				'name',            f.name,
				'alt', f.alt,
				'caption',         f.caption,
				'ext',             f.ext,
				'provider',        f.provider,
				'width',           f.width,
				'height',          f.height,
				'size',            f.size,
				'url',             f.url
				) FROM files f WHERE f.id = i.bg_image_id)
			) FROM files f WHERE f.id = i.bg_image_id) 
	))
FROM sliders_items i
WHERE i.slider_id = s.id AND
	i.enabled = true AND
	i.deleted IS NULL) as attributes
FROM sliders s
WHERE
	s.enabled = true AND
	s.deleted IS NULL AND
	s.provider = $1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, provider)
	err := row.Scan(&slider.ID, &sliderItems)

	switch err {
	case sql.ErrNoRows:
		return nil, NewError(errors.New("slider not found"))
	case nil:
		slider.Attributes = &models.PublicSliderAttributes{SliderItems: sliderItems}
		return &slider, nil
	default:
		return nil, err
	}
}
