package repository

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
)

type Category struct {
	Created strfmt.Date `db:"created"`
	Updated strfmt.Date `db:"updated"`
	Deleted strfmt.Date `db:"deleted"`
}

type CategoriesRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	MigrateContext(ctx context.Context) error
}

// CategoriesRepository - repository to store frontend categories.
type CategoriesRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewCategoriesRepository(db *sqlx.DB, tbl string) *CategoriesRepository {
	if db == nil {
		return nil
	}

	return &CategoriesRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *CategoriesRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *CategoriesRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *CategoriesRepository) MigrateContext(ctx context.Context) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.ExecContext(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id                     SERIAL PRIMARY KEY,
	parent_id              INTEGER,
	provider               varchar(255),
	title                  varchar(255) UNIQUE NOT NULL,
	slug                   varchar(255),
	description            text,
	content                text,
	image_id               INTEGER,
	enabled                bool NOT NULL DEFAULT TRUE,
	order_by               integer,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE categories ADD CONSTRAINT categories_image_id_foreign
	FOREIGN KEY (image_id) REFERENCES files(id);
ALTER TABLE categories ADD CONSTRAINT categories_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE categories ADD CONSTRAINT categories_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);`, r.TBL))

	return err
}
