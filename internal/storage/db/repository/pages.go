package repository

import (
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
