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

type Page struct {
	Created     strfmt.Date `db:"created"`
	Deleted     strfmt.Date `db:"deleted"`
	Updated     strfmt.Date `db:"updated"`
	Title       string      `db:"title"`
	Slug        string      `db:"slug"`
	Description string      `db:"description"`
	Content     string      `db:"content"`
	ImageID     int64       `db:"image_id"`
	OrderBy     int64       `db:"order_by"`
	CreatedBy   int64       `db:"created_by"`
	UpdatedBy   int64       `db:"updated_by"`
	ID          int64       `db:"id"`
	CategoryID  int64       `db:"category_id"`
	Enabled     bool        `db:"enabled"`
}

type PagesRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, pageMap map[string]interface{}) error
	Update(ctx context.Context, pageMap map[string]interface{}) ([]*models.Page, error)
	GetPages(ctx context.Context) ([]*models.Page, error)
	GetPageByID(ctx context.Context, id int64) (*models.Page, error)
	GetPageBySlug(ctx context.Context, provider string) (*models.PublicPage, error)
}

// PagesRepository - repository to store frontend pages.
type PagesRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewPagesRepository(db *sqlx.DB, tbl string) *PagesRepository {
	if db == nil {
		return nil
	}

	return &PagesRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *PagesRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *PagesRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *PagesRepository) Create(ctx context.Context, pageMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO pages (category_id, title, slug, description, content,
							image_id, enabled, order_by, created_by_id, updated_by_id)
VALUES (:category_id, :title, :slug, :description, :content,
		:image_id, :enabled, :order_by, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING`, pageMap)

	return err
}

func (r *PagesRepository) Update(ctx context.Context, pageMap map[string]interface{}) ([]*models.Page, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE pages SET
category_id=:category_id,
title=:title,
slug=:slug,
description=:description,
content=:content,
image_id=:image_id,
enabled=:enabled,
order_by=:order_by,
updated_by_id=:updated_by_id WHERE id =:id`, pageMap)
	if err != nil {
		return nil, err
	}

	return r.GetPages(ctx)
}

func (r *PagesRepository) GetPages(ctx context.Context) ([]*models.Page, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		pages    []*models.Page
	)

	rows, err := r.DB.QueryxContext(ctx,
		"SELECT id, category_id, title, slug, description, content, image_id, enabled, order_by, created_by_id, updated_by_id from pages;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var page Page

		err = rows.Scan(&page.ID, &page.CategoryID, &page.Title, &page.Slug, &page.Description, &page.Content,
			&page.ImageID, &page.Enabled, &page.OrderBy, &page.CreatedBy, &page.UpdatedBy)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &models.Page{
			ID:          page.ID,
			CategoryID:  page.CategoryID,
			Title:       page.Title,
			Slug:        page.Slug,
			Description: page.Description,
			Content:     page.Content,
			ImageID:     page.ImageID,
			Enabled:     page.Enabled,
			OrderBy:     page.OrderBy,
			CreatedByID: page.CreatedBy,
			UpdatedByID: page.UpdatedBy,
		})
	}

	defer rows.Close()

	return pages, multierr.ErrorOrNil()
}

func (r *PagesRepository) GetPageByID(ctx context.Context, id int64) (*models.Page, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var page Page

	sqlStatement := `SELECT id, category_id, title, slug, description, content, image_id, enabled, order_by, created_by_id, updated_by_id from pages where id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&page.ID, &page.CategoryID, &page.Title, &page.Slug, &page.Description, &page.Content,
		&page.ImageID, &page.Enabled, &page.OrderBy, &page.CreatedBy, &page.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("page not found"))
		}

		return nil, NewError(err)
	}

	return &models.Page{
		ID:          page.ID,
		CategoryID:  page.CategoryID,
		Title:       page.Title,
		Slug:        page.Slug,
		Description: page.Description,
		Content:     page.Content,
		ImageID:     page.ImageID,
		Enabled:     page.Enabled,
		OrderBy:     page.OrderBy,
		CreatedByID: page.CreatedBy,
		UpdatedByID: page.UpdatedBy,
	}, nil
}

func (r *PagesRepository) GetPageBySlug(ctx context.Context, provider string) (*models.PublicPage, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var page models.PublicPage

	var attributes AggregatedPageAttributes
	// home-01
	sqlStatement := `
SELECT
	p.id,
	(SELECT JSON_BUILD_OBJECT(
	   'title',       p.title,
	   'description', p.description,
	   'category',    (SELECT title FROM categories WHERE id = p.category_id),
	   'content',     p.content,
	   'link',        p.slug,
	   'image',       (SELECT JSON_BUILD_OBJECT(
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
						  ) FROM files f WHERE f.id = p.image_id)
	   				  ) FROM files f WHERE f.id = p.image_id)
	)) as attributes
   FROM pages p
   WHERE
      p.enabled = true AND
      p.deleted isNULL AND
	  p.slug=$1;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, provider)
	if err := row.Scan(&page.ID, &attributes); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("page not found"))
		}

		return nil, err
	}

	page.Attributes = &models.PublicPageAttributes{
		Title:       attributes.Title,
		Description: attributes.Description,
		Content:     attributes.Content,
		Link:        attributes.Link,
		Image:       attributes.Image,
		Category:    attributes.Category,
	}

	return &page, nil
}
