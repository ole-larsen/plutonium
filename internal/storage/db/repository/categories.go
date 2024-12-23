package repository

import (
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
