package repository

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
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
	MigrateContext(ctx context.Context) error
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

func (r *PagesRepository) MigrateContext(ctx context.Context) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
				id                     SERIAL PRIMARY KEY,
				category_id            INTEGER NOT NULL,
				title                  varchar(255) UNIQUE NOT NULL,
				slug                   varchar(255),
			    description            text,
			    content                text,
			    image_id               INTEGER,
				enabled                bool NOT NULL DEFAULT TRUE,
			    is_menu                bool NOT NULL DEFAULT FALSE,
			    order_by               integer,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			
			ALTER TABLE pages ADD CONSTRAINT pages_category_id_foreign
				FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE;
			ALTER TABLE pages ADD CONSTRAINT pages_image_id_foreign
				FOREIGN KEY (image_id) REFERENCES files(id);
			ALTER TABLE pages ADD CONSTRAINT pages_created_by_id_foreign
				FOREIGN KEY (created_by_id) REFERENCES users(id);
			ALTER TABLE pages ADD CONSTRAINT pages_updated_by_id_foreign
				FOREIGN KEY (updated_by_id) REFERENCES users(id);`, r.TBL))

	return err
}
