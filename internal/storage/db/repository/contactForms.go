package repository

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
)

type ContactForm struct {
	Created   strfmt.Date `db:"created"`
	Updated   strfmt.Date `db:"updated"`
	Deleted   strfmt.Date `db:"deleted"`
	Provider  string      `db:"provider"`
	Name      string      `db:"name"`
	Email     string      `db:"email"`
	Subject   string      `db:"subject"`
	Message   string      `db:"message"`
	ID        int64       `db:"id"`
	PageID    int64       `db:"page_id"`
	CreatedBy int64       `db:"created_by"`
	UpdatedBy int64       `db:"updated_by"`
}

type ContactFormRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, contactFormMap map[string]interface{}) error
	CreateSubscribe(ctx context.Context, contactFormMap map[string]interface{}) error
}

// ContactFormRepository - repository to store users.
type ContactFormRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewContactFormRepository(db *sqlx.DB, tbl string) *ContactFormRepository {
	if db == nil {
		return nil
	}

	return &ContactFormRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *ContactFormRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *ContactFormRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *ContactFormRepository) Create(ctx context.Context, contactFormMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO contacts_forms (provider, page_id, name, email, subject, message)
		VALUES (:provider, :page_id, :name, :email, :subject, :message)
		ON CONFLICT DO NOTHING`, contactFormMap)

	return err
}

func (r *ContactFormRepository) CreateSubscribe(ctx context.Context, contactFormMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO subscribe_forms (provider, email)
		VALUES (:provider, :email)
		ON CONFLICT DO NOTHING`, contactFormMap)

	return err
}
