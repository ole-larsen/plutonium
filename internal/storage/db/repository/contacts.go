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

type Contact struct {
	Created    strfmt.Date `db:"created"`
	Deleted    strfmt.Date `db:"deleted"`
	Updated    strfmt.Date `db:"updated"`
	Heading    string      `db:"heading"`
	SubHeading string      `db:"sub_heading"`
	BtnLink    string      `db:"btn_link"`
	BtnText    string      `db:"btn_text"`
	ImageID    int64       `db:"image_id"`
	CreatedBy  int64       `db:"created_by"`
	UpdatedBy  int64       `db:"updated_by"`
	ID         int64       `db:"id"`
	PageID     int64       `db:"page_id"`
	Enabled    bool        `db:"enabled"`
}

type ContactsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, contactMap map[string]interface{}) error
	Update(ctx context.Context, contactMap map[string]interface{}) ([]*models.Contact, error)
	GetContacts(ctx context.Context) ([]*models.Contact, error)
	GetContactByID(ctx context.Context, id int64) (*models.Contact, error)
	GetContactByPageID(ctx context.Context, pageID int64) (*models.PublicContact, error)
}

// ContactsRepository - repository to store users.
type ContactsRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewContactsRepository(db *sqlx.DB, tbl string) *ContactsRepository {
	if db == nil {
		return nil
	}

	return &ContactsRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *ContactsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *ContactsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *ContactsRepository) Create(ctx context.Context, contactMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
		INSERT INTO contacts (page_id, heading, sub_heading,
		                           image_id, enabled, created_by_id, updated_by_id)
		VALUES (:page_id, :heading, :sub_heading, :btn_link, :btn_text,
		        :image_id, :enabled, :created_by_id, :updated_by_id)
		ON CONFLICT DO NOTHING`, contactMap)

	return err
}

func (r *ContactsRepository) Update(ctx context.Context, contactMap map[string]interface{}) ([]*models.Contact, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE contacts SET
				page_id=:page_id,
                heading=:heading,
                sub_heading=:sub_heading,
                image_id=:image_id,
                enabled=:enabled,
                updated_by_id=:updated_by_id WHERE id =:id`, contactMap)
	if err != nil {
		return nil, err
	}

	return r.GetContacts(ctx)
}

func (r *ContactsRepository) GetContacts(ctx context.Context) ([]*models.Contact, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		contacts []*models.Contact
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT id, page_id, heading, sub_heading, btn_link, btn_text, image_id, enabled, created_by_id, updated_by_id from contacts;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var contact Contact

		err = rows.Scan(&contact.ID, &contact.PageID, &contact.Heading, &contact.SubHeading, &contact.BtnLink, &contact.BtnText,
			&contact.ImageID, &contact.Enabled, &contact.CreatedBy, &contact.UpdatedBy)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, &models.Contact{
			ID:          contact.ID,
			PageID:      contact.PageID,
			Heading:     contact.Heading,
			SubHeading:  contact.SubHeading,
			BtnLink:     contact.BtnLink,
			BtnText:     contact.BtnText,
			ImageID:     contact.ImageID,
			Enabled:     contact.Enabled,
			CreatedByID: contact.CreatedBy,
			UpdatedByID: contact.UpdatedBy,
		})
	}

	defer rows.Close()

	return contacts, multierr.ErrorOrNil()
}

func (r *ContactsRepository) GetContactByID(ctx context.Context, id int64) (*models.Contact, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var contact Contact

	sqlStatement := `SELECT 
    id, page_id, heading, sub_heading, image_id,
      					enabled, created_by_id, updated_by_id from contacts where id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(
		&contact.ID,
		&contact.PageID,
		&contact.Heading,
		&contact.SubHeading,
		&contact.BtnLink,
		&contact.BtnText,
		&contact.ImageID,
		&contact.Enabled,
		&contact.CreatedBy,
		&contact.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("contact not found"))
		}

		return nil, NewError(err)
	}

	return &models.Contact{
		ID:          contact.ID,
		PageID:      contact.PageID,
		Heading:     contact.Heading,
		SubHeading:  contact.SubHeading,
		BtnLink:     contact.BtnLink,
		BtnText:     contact.BtnText,
		ImageID:     contact.ImageID,
		Enabled:     contact.Enabled,
		CreatedByID: contact.CreatedBy,
		UpdatedByID: contact.UpdatedBy,
	}, nil
}

func (r *ContactsRepository) GetContactByPageID(ctx context.Context, pageID int64) (*models.PublicContact, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var contact models.PublicContact

	var attributes AggregatedContactAttributes
	// home-01
	sqlStatement := `SELECT
	c.id,
	(SELECT JSON_BUILD_OBJECT(
	   'heading',     c.heading,
	   'subHeading',  c.sub_heading,
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
	   				  				  ) FROM files f WHERE
	                  	                f.id = c.image_id AND f.provider = 'contact')
	   				  ) FROM files f WHERE
	                  	f.id = c.image_id AND f.provider = 'contact')
	)) as attributes
   FROM contacts c
   WHERE
      c.enabled = true AND
	  c.deleted isNULL AND
      c.page_id = $1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, pageID)

	err := row.Scan(&contact.ID, &attributes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("contact not found")
		}

		return nil, err
	}

	contact.Attributes = &models.PublicContactAttributes{
		Heading:    attributes.Heading,
		Image:      attributes.Image,
		SubHeading: attributes.SubHeading,
	}

	return &contact, nil
}
