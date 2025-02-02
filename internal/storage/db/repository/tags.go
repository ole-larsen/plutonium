package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ole-larsen/plutonium/models"
)

func ConvertSlice[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))

	for _, v := range in {
		if converted, ok := v.(E); ok {
			out = append(out, converted)
		}
	}

	return
}
func ConvertInterfaceToSlice(input interface{}) []interface{} {
	var out []interface{}

	rv := reflect.ValueOf(input)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			out = append(out, rv.Index(i).Interface())
		}
	}

	return out
}

type Tag struct {
	Created   strfmt.Date   `db:"created"`
	Updated   strfmt.Date   `db:"updated"`
	Deleted   strfmt.Date   `db:"deleted"`
	Title     string        `db:"title"`
	Slug      string        `db:"slug"`
	BlogID    pq.Int64Array `db:"blog_id"`
	PageID    pq.Int64Array `db:"page_id"`
	ID        int64         `db:"id"`
	ParentID  int64         `db:"parent_id"`
	CreatedBy int64         `db:"created_by"`
	UpdatedBy int64         `db:"updated_by"`
	Enabled   bool          `db:"enabled"`
}

type TagsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, tagMap map[string]interface{}) error
	Update(ctx context.Context, tagMap map[string]interface{}) ([]*models.Tag, error)
	GetTags(ctx context.Context) ([]*models.Tag, error)
	GetTagByID(ctx context.Context, id int64) (*models.Tag, error)
}

// TagsRepository - repository to store frontend tags.
type TagsRepository struct {
	Blogs BlogsRepositoryInterface
	Pages PagesRepositoryInterface
	DB    sqlx.DB
	TBL   string
}

func NewTagsRepository(
	db *sqlx.DB,
	tbl string,
	blogs BlogsRepositoryInterface,
	pages PagesRepositoryInterface,
) *TagsRepository {
	if db == nil {
		return nil
	}

	return &TagsRepository{
		DB:    *db,
		TBL:   tbl,
		Blogs: blogs,
		Pages: pages,
	}
}

func (r *TagsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *TagsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *TagsRepository) SetBlogsRelations(ctx context.Context, tagMap map[string]interface{}) error {
	if tagMap["blog_id"] != nil {
		blogIDs := ConvertSlice[int64](ConvertInterfaceToSlice(tagMap["blog_id"]))
		for _, blogID := range blogIDs {
			blog, err := r.Blogs.GetBlogByID(ctx, blogID)
			if err != nil {
				return err
			}

			if blog != nil {
				arg := map[string]interface{}{
					"tag_id":  tagMap["id"],
					"blog_id": blogID,
				}

				_, err = r.DB.NamedExecContext(ctx, `INSERT INTO tags_blogs (tag_id, blog_id) VALUES (:tag_id, :blog_id) ON CONFLICT DO NOTHING`, arg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *TagsRepository) SetPagesRelations(ctx context.Context, tagMap map[string]interface{}) error {
	if tagMap["page_id"] != nil {
		pageIDs := ConvertSlice[int64](ConvertInterfaceToSlice(tagMap["page_id"]))
		for _, pageID := range pageIDs {
			page, err := r.Pages.GetPageByID(ctx, pageID)
			if err != nil {
				return err
			}

			if page != nil {
				arg := map[string]interface{}{
					"tag_id":  tagMap["id"],
					"page_id": pageID,
				}

				_, err = r.DB.NamedExecContext(ctx, `INSERT INTO tags_pages (tag_id, page_id) VALUES (:tag_id, :page_id) ON CONFLICT DO NOTHING`, arg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *TagsRepository) Create(ctx context.Context, tagMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	sqlResult, err := r.DB.NamedExecContext(ctx, `
INSERT INTO tags (parent_id, title, slug, enabled, created_by_id, updated_by_id)
VALUES (:parent_id, :title, :slug, :enabled, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING RETURNING id`, tagMap)
	if err != nil {
		return err
	}

	tagID, err := sqlResult.LastInsertId()

	if err != nil {
		return err
	}

	tagMap["id"] = tagID
	if err := r.SetPagesRelations(ctx, tagMap); err != nil {
		return err
	}

	if err := r.SetBlogsRelations(ctx, tagMap); err != nil {
		return err
	}

	return nil
}

func (r *TagsRepository) Update(ctx context.Context, tagMap map[string]interface{}) ([]*models.Tag, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE tags SET
parent_id=:parent_id,
title=:title,
slug=:slug,
enabled=:enabled,
updated_by_id=:updated_by_id WHERE id =:id`, tagMap)
	if err != nil {
		return nil, err
	}

	tagID, ok := tagMap["id"].(int64)

	if !ok {
		return nil, NewError(fmt.Errorf("could not get tag id"))
	}

	tagMap["id"] = tagID
	if err := r.SetPagesRelations(ctx, tagMap); err != nil {
		return nil, err
	}

	if err := r.SetBlogsRelations(ctx, tagMap); err != nil {
		return nil, err
	}

	return r.GetTags(ctx)
}

func (r *TagsRepository) GetTags(ctx context.Context) ([]*models.Tag, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		tags []*models.Tag
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT
	t.id,
	t.parent_id,
	t.title,
	t.slug,
	t.enabled,
	t.created_by_id,
	t.updated_by_id,
	ARRAY(SELECT DISTINCT blog_id FROM unnest(array_remove(ARRAY_AGG(tags_blogs.blog_id), NULL)) blog_id) blogs_id,
	ARRAY(SELECT DISTINCT page_id FROM unnest(array_remove(ARRAY_AGG(tags_pages.page_id), NULL)) page_id) page_id
FROM tags t
	LEFT JOIN tags_blogs ON tags_blogs.tag_id = t.id
	LEFT JOIN tags_pages ON tags_pages.tag_id = t.id
GROUP BY t.id;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.ParentID, &tag.Title, &tag.Slug, &tag.Enabled, &tag.CreatedBy, &tag.UpdatedBy, &tag.BlogID, &tag.PageID); err != nil {
			return nil, err
		}

		tags = append(tags, &models.Tag{
			ID:          tag.ID,
			ParentID:    tag.ParentID,
			BlogID:      tag.BlogID,
			PageID:      tag.PageID,
			Title:       tag.Title,
			Slug:        tag.Slug,
			Enabled:     tag.Enabled,
			CreatedByID: tag.CreatedBy,
			UpdatedByID: tag.UpdatedBy,
		})
	}

	defer rows.Close()

	return tags, nil
}

func (r *TagsRepository) GetTagByID(ctx context.Context, id int64) (*models.Tag, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var tag Tag

	sqlStatement := `
SELECT
	t.id,
	t.parent_id,
	t.title,
	t.slug,
	t.enabled,
	t.created_by_id,
	t.updated_by_id,
	ARRAY(SELECT DISTINCT blog_id FROM unnest(array_remove(ARRAY_AGG(tags_blogs.blog_id), NULL)) blog_id) blog_id,
	ARRAY(SELECT DISTINCT page_id FROM unnest(array_remove(ARRAY_AGG(tags_pages.page_id), NULL)) page_id) page_id
FROM tags t
	LEFT JOIN tags_blogs ON tags_blogs.tag_id = t.id
	LEFT JOIN tags_pages ON tags_pages.tag_id = t.id
WHERE t.id=$1 GROUP BY t.id;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	if err := row.Scan(
		&tag.ID,
		&tag.ParentID,
		&tag.Title,
		&tag.Slug,
		&tag.Enabled,
		&tag.CreatedBy,
		&tag.UpdatedBy,
		&tag.BlogID,
		&tag.PageID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("tag not found"))
		}

		return nil, NewError(err)
	}

	return &models.Tag{
		ID:          tag.ID,
		ParentID:    tag.ParentID,
		BlogID:      tag.BlogID,
		PageID:      tag.PageID,
		Title:       tag.Title,
		Slug:        tag.Slug,
		Enabled:     tag.Enabled,
		CreatedByID: tag.CreatedBy,
		UpdatedByID: tag.UpdatedBy,
	}, nil
}
