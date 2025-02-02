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

type Faq struct {
	Created   strfmt.Date `db:"created"`
	Updated   strfmt.Date `db:"updated"`
	Deleted   strfmt.Date `db:"deleted"`
	Question  string      `db:"question"`
	Answer    string      `db:"answer"`
	ID        int64       `db:"id"`
	OrderBy   int64       `db:"order_by"`
	CreatedBy int64       `db:"created_by"`
	UpdatedBy int64       `db:"updated_by"`
	Enabled   bool        `db:"enabled"`
}

type FaqsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, faqMap map[string]interface{}) error
	Update(ctx context.Context, faqMap map[string]interface{}) ([]*models.Faq, error)
	GetFaqs(ctx context.Context) ([]*models.Faq, error)
	GetFaqByID(ctx context.Context, id int64) (*models.Faq, error)
	GetPublicFaqs(ctx context.Context) ([]*models.PublicFaqItem, error)
}

// FaqsRepository - repository to store users.
type FaqsRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewFaqsRepository(db *sqlx.DB, tbl string) *FaqsRepository {
	if db == nil {
		return nil
	}

	return &FaqsRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *FaqsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *FaqsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *FaqsRepository) Create(ctx context.Context, faqMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO faqs (question, answer, enabled, order_by, created_by_id, updated_by_id)
		VALUES (:question, :answer, :enabled, :order_by, :created_by_id, :updated_by_id)
		ON CONFLICT DO NOTHING`, faqMap)

	return err
}

func (r *FaqsRepository) Update(ctx context.Context, faqMap map[string]interface{}) ([]*models.Faq, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE faqs SET
                question=:question,
                answer=:answer,
                enabled=:enabled,
                order_by=:order_by,
                updated_by_id=:updated_by_id WHERE id =:id`, faqMap)
	if err != nil {
		return nil, err
	}

	return r.GetFaqs(ctx)
}

func (r *FaqsRepository) GetFaqs(ctx context.Context) ([]*models.Faq, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		faqs     []*models.Faq
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT id, question, answer, enabled, order_by, created_by_id, updated_by_id from faqs;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var faq Faq

		err = rows.Scan(&faq.ID, &faq.Question, &faq.Answer,
			&faq.Enabled, &faq.OrderBy, &faq.CreatedBy, &faq.UpdatedBy)
		if err != nil {
			return nil, err
		}

		faqs = append(faqs, &models.Faq{
			ID:          faq.ID,
			Question:    faq.Question,
			Answer:      faq.Answer,
			Enabled:     faq.Enabled,
			OrderBy:     faq.OrderBy,
			CreatedByID: faq.CreatedBy,
			UpdatedByID: faq.UpdatedBy,
		})
	}

	defer rows.Close()

	return faqs, multierr.ErrorOrNil()
}

func (r *FaqsRepository) GetFaqByID(ctx context.Context, id int64) (*models.Faq, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var faq Faq

	sqlStatement := `SELECT id, question, answer, enabled, order_by, created_by_id, updated_by_id from faqs where id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&faq.ID, &faq.Question, &faq.Answer,
		&faq.Enabled, &faq.OrderBy, &faq.CreatedBy, &faq.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("faq not found"))
		}

		return nil, NewError(err)
	}

	return &models.Faq{
		ID:          faq.ID,
		Question:    faq.Question,
		Answer:      faq.Answer,
		Enabled:     faq.Enabled,
		OrderBy:     faq.OrderBy,
		CreatedByID: faq.CreatedBy,
		UpdatedByID: faq.UpdatedBy,
	}, nil
}

func (r *FaqsRepository) GetPublicFaqs(ctx context.Context) ([]*models.PublicFaqItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		faqs     []*models.PublicFaqItem
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT id, question, answer from faqs WHERE enabled = true AND deleted isNULL GROUP BY id ORDER BY order_by ASC;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var faq Faq

		if err := rows.Scan(&faq.ID, &faq.Question, &faq.Answer); err != nil {
			return nil, err
		}

		faqs = append(faqs, &models.PublicFaqItem{
			ID: faq.ID,
			Attributes: &models.PublicFaqItemAttributes{
				Question: faq.Question,
				Answer:   faq.Answer,
			},
		})
	}

	defer rows.Close()

	return faqs, multierr.ErrorOrNil()
}
