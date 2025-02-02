package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ole-larsen/plutonium/models"
)

type BlogTag struct {
	Link  string `db:"link"`
	Title string `db:"title"`
}

type Blog struct {
	Deleted     strfmt.Date   `db:"deleted"`
	Updated     strfmt.Date   `db:"updated"`
	Created     strfmt.Date   `db:"created"`
	PublicDate  strfmt.Date   `db:"public_date"`
	Image1ID    *int64        `db:"image_1_id"`
	Image2ID    *int64        `db:"image_2_id"`
	Image2      *int64        `db:"image2"`
	Image3      *int64        `db:"image3"`
	Image1      *int64        `db:"image1"`
	Image3ID    *int64        `db:"image_3_id"`
	Title       string        `db:"title"`
	Link        string        `db:"link"`
	Description string        `db:"description"`
	Slug        string        `db:"slug"`
	Content     string        `db:"content"`
	TagID       pq.Int64Array `db:"tag_id"`
	Tags        []BlogTag     `db:"tags"`
	Author      Author        `db:"author"`
	AuthorID    int64         `db:"author_id"`
	OrderBy     int64         `db:"order_by"`
	CreatedBy   int64         `db:"created_by"`
	UpdatedBy   int64         `db:"updated_by"`
	ImageID     int64         `db:"image_id"`
	Image       int64         `db:"image"`
	ID          int64         `db:"id"`
	Enabled     bool          `db:"enabled"`
}

type BlogsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, blogMap map[string]interface{}) error
	Update(ctx context.Context, blogMap map[string]interface{}) ([]*models.Blog, error)
	GetBlogs(ctx context.Context) ([]*models.Blog, error)
	GetBlogByID(ctx context.Context, id int64) (*models.Blog, error)
	GetPublicBlogs(ctx context.Context) ([]*models.PublicBlogItem, error)
	GetPublicBlogItem(ctx context.Context, slug string) (*models.PublicBlogItem, error)
}

// BlogsRepository - repository to store frontend blogs.
type BlogsRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewBlogsRepository(db *sqlx.DB, tbl string) *BlogsRepository {
	if db == nil {
		return nil
	}

	return &BlogsRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *BlogsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *BlogsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *BlogsRepository) Create(ctx context.Context, blogMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO blogs (title, slug, description, content,
							author_id, image_id, image_1_id, image_2_id, image_3_id, enabled, order_by, created_by_id, updated_by_id)
VALUES (:title, :slug, :description, :content, :author_id,
		:image_id, :image_1_id, :image_2_id, :image_3_id, :enabled, :order_by, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING`, blogMap)

	return err
}

func (r *BlogsRepository) Update(ctx context.Context, blogMap map[string]interface{}) ([]*models.Blog, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE blogs SET
title=:title,
slug=:slug,
description=:description,
content=:content,
image_id=:image_id,
image_1_id=:image_1_id,
image_2_id=:image_2_id,
image_3_id=:image_3_id,
author_id=:author_id,
enabled=:enabled,
order_by=:order_by,
updated_by_id=:updated_by_id WHERE id =:id`, blogMap)
	if err != nil {
		return nil, err
	}

	return r.GetBlogs(ctx)
}

func (r *BlogsRepository) GetBlogs(ctx context.Context) ([]*models.Blog, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		blogs    []*models.Blog
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT
	b.id,
   	b.title,
   	b.slug,
   	b.description,
   	b.content,
   	b.author_id,
   	b.image_id,
   	b.image_1_id,
   	b.image_2_id,
   	b.image_3_id,
   	b.enabled,
   	b.order_by,
   	b.public_date,
   	b.created_by_id,
   	b.updated_by_id,
	array_remove(ARRAY_AGG(tags_blogs.tag_id), NULL) blog_id
FROM blogs b
LEFT JOIN tags_blogs ON tags_blogs.blog_id = b.id
GROUP BY b.id;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blog Blog

		err = rows.Scan(&blog.ID, &blog.Title, &blog.Slug, &blog.Description, &blog.Content, &blog.AuthorID,
			&blog.ImageID, &blog.Image1ID, &blog.Image2ID, &blog.Image3ID, &blog.Enabled, &blog.OrderBy, &blog.PublicDate,
			&blog.CreatedBy, &blog.UpdatedBy, &blog.TagID)
		if err != nil {
			return nil, err
		}

		var image1ID int64
		if blog.Image1ID != nil {
			image1ID = *blog.Image1ID
		}

		var image2ID int64
		if blog.Image2ID != nil {
			image2ID = *blog.Image2ID
		}

		var image3ID int64
		if blog.Image3ID != nil {
			image3ID = *blog.Image3ID
		}

		blogs = append(blogs, &models.Blog{
			ID:          blog.ID,
			TagID:       blog.TagID,
			Title:       blog.Title,
			Slug:        blog.Slug,
			Description: blog.Description,
			Content:     blog.Content,
			AuthorID:    blog.AuthorID,
			ImageID:     blog.ImageID,
			Image1ID:    image1ID,
			Image2ID:    image2ID,
			Image3ID:    image3ID,
			Enabled:     blog.Enabled,
			OrderBy:     blog.OrderBy,
			PublicDate:  blog.PublicDate,
			CreatedByID: blog.CreatedBy,
			UpdatedByID: blog.UpdatedBy,
		})
	}

	defer rows.Close()

	return blogs, multierr.ErrorOrNil()
}

func (r *BlogsRepository) GetBlogByID(ctx context.Context, id int64) (*models.Blog, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var blog Blog

	sqlStatement := `
SELECT
	b.id,
	b.title,
	b.slug,
	b.description,
	b.content,
	b.author_id,
	b.image_id,
	b.image_1_id,
	b.image_2_id,
	b.image_3_id,
	b.enabled,
	b.order_by,
	b.public_date,
	b.created_by_id,
	b.updated_by_id,
	array_remove(ARRAY_AGG(tags_blogs.tag_id), NULL) blog_id
FROM blogs b
LEFT JOIN tags_blogs ON tags_blogs.blog_id = b.id
where b.id=$1 GROUP BY b.id;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(
		&blog.ID, &blog.Title, &blog.Slug, &blog.Description,
		&blog.Content, &blog.AuthorID,
		&blog.ImageID, &blog.Image1ID, &blog.Image2ID, &blog.Image3ID, &blog.Enabled, &blog.OrderBy, &blog.PublicDate,
		&blog.CreatedBy, &blog.UpdatedBy, &blog.TagID)
	switch err {
	case sql.ErrNoRows:
		return nil, NewError(fmt.Errorf("blog not found"))
	case nil:
		var image1ID int64
		if blog.Image1ID != nil {
			image1ID = *blog.Image1ID
		}

		var image2ID int64
		if blog.Image2ID != nil {
			image2ID = *blog.Image2ID
		}

		var image3ID int64
		if blog.Image3ID != nil {
			image3ID = *blog.Image3ID
		}

		return &models.Blog{
			ID:          blog.ID,
			TagID:       blog.TagID,
			Title:       blog.Title,
			Slug:        blog.Slug,
			Description: blog.Description,
			Content:     blog.Content,
			AuthorID:    blog.AuthorID,
			ImageID:     blog.ImageID,
			Image1ID:    image1ID,
			Image2ID:    image2ID,
			Image3ID:    image3ID,
			Enabled:     blog.Enabled,
			OrderBy:     blog.OrderBy,
			PublicDate:  blog.PublicDate,
			CreatedByID: blog.CreatedBy,
			UpdatedByID: blog.UpdatedBy,
		}, err
	default:
		return nil, err
	}
}

func (r *BlogsRepository) GetPublicBlogs(ctx context.Context) ([]*models.PublicBlogItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		blogs    []*models.PublicBlogItem
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT
	b.id,
	b.title,
	b.slug as link,
	b.description,
	b.content,
	(SELECT JSON_BUILD_OBJECT(
		'id', a.id,
		'attributes', (SELECT JSON_BUILD_OBJECT(
			'title',           a.title,
			'description',     a.description,
			'name',            a.name,
			'link',            a.slug,
			'image',           (SELECT JSON_BUILD_OBJECT(
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
								) FROM files f WHERE f.id = a.image_id)
		) FROM authors a WHERE a.id = b.author_id)
		) FROM authors a WHERE a.id = b.author_id) as author,
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
						) FROM files f WHERE f.id = b.image_id)
		) FROM files f WHERE f.id = b.image_id) as image,
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
						) FROM files f WHERE f.id = b.image_1_id)
		) FROM files f WHERE f.id = b.image_1_id) as image1,
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
						) FROM files f WHERE f.id = b.image_2_id)
		) FROM files f WHERE f.id = b.image_2_id) as image2,
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
						) FROM files f WHERE f.id = b.image_3_id)
		) FROM files f WHERE f.id = b.image_3_id) as image3,
	b.public_date,
	(SELECT (SELECT ARRAY_AGG(JSON_BUILD_OBJECT(
		'title', t.title,
		'link', t.slug
	))) FROM blogs b2
		LEFT JOIN tags_blogs ON tags_blogs.blog_id = b2.id
		LEFT JOIN tags t
		ON t.id = tags_blogs.tag_id AND t.enabled = true AND t.deleted isNULL
		WHERE b.id = b2.id) as tags
FROM blogs b
WHERE 
	b.enabled = true AND b.deleted isNULL  GROUP BY b.id ORDER BY b.order_by ASC;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blog Blog

		var author AggregatedAuthorJSON

		var image AggregatedImageJSON

		var image1 AggregatedImageJSON

		var image2 AggregatedImageJSON

		var image3 AggregatedImageJSON

		var tags AggregatedTagJSON

		err = rows.Scan(&blog.ID, &blog.Title, &blog.Link, &blog.Description, &blog.Content, &author,
			&image, &image1, &image2, &image3, &blog.PublicDate, &tags)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &models.PublicBlogItem{
			ID: blog.ID,
			Attributes: &models.PublicBlogItemAttributes{
				Title:       blog.Title,
				Link:        blog.Link,
				Description: blog.Description,
				Content:     blog.Content,
				Author: &models.PublicAuthor{
					Attributes: author.Attributes,
					ID:         author.ID,
				},
				Image: &models.PublicFile{
					Attributes: image.Attributes,
					ID:         image.ID,
				},
				Image1: &models.PublicFile{
					Attributes: image1.Attributes,
					ID:         image1.ID,
				},
				Image2: &models.PublicFile{
					Attributes: image2.Attributes,
					ID:         image2.ID,
				},
				Image3: &models.PublicFile{
					Attributes: image3.Attributes,
					ID:         image3.ID,
				},
				Date: blog.PublicDate,
			},
		})
	}

	defer rows.Close()

	return blogs, multierr.ErrorOrNil()
}

func (r *BlogsRepository) GetPublicBlogItem(ctx context.Context, slug string) (*models.PublicBlogItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var blog models.PublicBlogItem

	sqlStatement := `SELECT
b.id,
b.title,
to_char( current_timestamp, 'dd-mm-yyyy' ) as date,
b.slug as link,
b.description,
b.order_by,
(SELECT JSON_BUILD_OBJECT(
	'id', a.id,
	'attributes', (SELECT JSON_BUILD_OBJECT(
		'title',           a.title,
		'description',     a.description,
		'name',            a.name,
		'link',            a.slug,
		'image',           (SELECT JSON_BUILD_OBJECT(
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
							) FROM files f WHERE f.id = a.image_id)
	) FROM authors a WHERE a.id = b.author_id)
	) FROM authors a WHERE a.id = b.author_id) as author,
(SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
	'title', t2.title,
	'link',  t2.slug
	))) FROM blogs b2
		LEFT JOIN tags_blogs ON tags_blogs.blog_id = b2.id
		LEFT JOIN tags t2
		ON t2.id = tags_blogs.tag_id AND t2.enabled = true AND t2.deleted isNULL
) popularTags,
(SELECT JSON_AGG(JSON_BUILD_OBJECT(
	'title', t.title,
	'link',  t.slug
) ORDER BY t.slug ASC)) tags,
(SELECT JSON_BUILD_OBJECT(
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
					) FROM files f WHERE f.id = b.image_id)
	) FROM files f WHERE f.id = b.image_id) as image,
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
									) FROM files f WHERE f.id = b.image_1_id)
					) FROM files f WHERE f.id = b.image_1_id) as image1,
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
									) FROM files f WHERE f.id = b.image_2_id)
					) FROM files f WHERE f.id = b.image_2_id) as image2,
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
									) FROM files f WHERE f.id = b.image_3_id)
					) FROM files f WHERE f.id = b.image_3_id) as image3
FROM blogs b
LEFT JOIN tags_blogs ON tags_blogs.blog_id = b.id
LEFT JOIN tags t ON t.id = tags_blogs.tag_id AND t.enabled = true AND t.deleted isNULL
WHERE
	b.slug = $1 AND
	b.enabled = true AND
	b.deleted isNULL
GROUP BY b.id;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, slug)

	var author AggregatedAuthorJSON

	var image AggregatedImageJSON

	var image1 AggregatedImageJSON

	var image2 AggregatedImageJSON

	var image3 AggregatedImageJSON

	var tags AggregatedTagJSON

	var popularTags AggregatedTagJSON
	err := row.Scan(&blog.ID, &blog.Attributes.Title, &blog.Attributes.Date, &blog.Attributes.Link, &blog.Attributes.Description,
		&blog.Attributes.OrderBy, &author, &popularTags, &tags, &image, &image1, &image2, &image3)

	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("blog not found")
	case nil:
		for _, tag := range popularTags {
			blog.Attributes.PopularTags = append(blog.Attributes.PopularTags, &models.PublicTag{
				Link:  "/tags/" + tag.Link,
				Title: tag.Title,
			})
		}
		// prevent nil panic
		blog.Attributes.Tags = make([]*models.PublicTag, 0)
		for _, tag := range tags {
			blog.Attributes.Tags = append(blog.Attributes.Tags, &models.PublicTag{
				Link:  "/tags/" + tag.Link,
				Title: tag.Title,
			})
		}

		blog.Attributes.Image = &models.PublicFile{
			Attributes: image.Attributes,
			ID:         image.ID,
		}
		blog.Attributes.Image1 = &models.PublicFile{
			Attributes: image1.Attributes,
			ID:         image1.ID,
		}
		blog.Attributes.Image2 = &models.PublicFile{
			Attributes: image2.Attributes,
			ID:         image2.ID,
		}
		blog.Attributes.Image3 = &models.PublicFile{
			Attributes: image3.Attributes,
			ID:         image3.ID,
		}
		blog.Attributes.Author = &models.PublicAuthor{
			Attributes: author.Attributes,
			ID:         author.ID,
		}

		return &blog, err
	default:
		return nil, err
	}
}
