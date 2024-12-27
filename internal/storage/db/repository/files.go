package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/models"
)

const PublicDir = "/api/v1/files/"

type File struct {
	Updated     strfmt.Date `db:"updated"`
	Created     strfmt.Date `db:"created"`
	Deleted     strfmt.Date `db:"deleted"`
	Metadata    interface{} `db:"metadata"`
	Formats     interface{} `db:"formats"`
	Provider    *string     `db:"provider"`
	Thumb       string      `db:"preview_url"`
	Hash        string      `db:"hash"`
	Ext         string      `db:"ext"`
	Caption     string      `db:"caption"`
	URL         string      `db:"url"`
	Alt         string      `db:"alt"`
	Name        string      `db:"name"`
	Mime        string      `db:"mime"`
	Height      int64       `db:"height"`
	CreatedByID int64       `db:"created_by_id"`
	UpdatedByID int64       `db:"updated_by_id"`
	Size        float64     `db:"size"`
	ID          int64       `db:"id"`
	Width       int64       `db:"width"`
}

type FilesRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error

	Create(ctx context.Context, fileMap map[string]interface{}) error
	Update(ctx context.Context, fileMap map[string]interface{}) ([]*models.File, error)
	GetFiles(ctx context.Context) ([]*models.File, error)
	GetFileByName(ctx context.Context, name string) (*models.File, error)
	GetFileByID(ctx context.Context, id int64) (*models.File, error)
	GetPublicFilesByProvider(ctx context.Context, provider string) ([]*models.PublicFile, error)
	GetPublicFileByName(ctx context.Context, name string) (*models.PublicFile, error)
	GetPublicFileByID(ctx context.Context, id int64) (*models.PublicFile, error)
}

// FilesRepository - repository to store files.
type FilesRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewFilesRepository(db *sqlx.DB, tbl string) *FilesRepository {
	if db == nil {
		return nil
	}

	return &FilesRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *FilesRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *FilesRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *FilesRepository) Create(ctx context.Context, fileMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO files (name, alt, caption, hash, mime, ext, size, width, height, provider, url, created_by_id, updated_by_id)
VALUES (:name, :alt, :caption, :hash, :mime, :ext, :size, :width, :height, :provider, :url, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING`, fileMap)

	return err
}

func (r *FilesRepository) Update(ctx context.Context, fileMap map[string]interface{}) ([]*models.File, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	fileMap["url"] = fmt.Sprintf("%s%s%s", PublicDir, fileMap["name"], fileMap["ext"])

	_, err := r.DB.NamedExecContext(ctx, `
UPDATE files SET
	name=:name,
	alt=:alt,
	hash=:hash,
	caption=:caption,
	ext=:ext,
	mime=:mime,
	size=:size,
	width=:width,
	height=:height,
	url=:url,
	provider=:provider WHERE id =:id`, fileMap)
	if err != nil {
		return nil, err
	}

	return r.GetFiles(ctx)
}

func (r *FilesRepository) GetFiles(ctx context.Context) ([]*models.File, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		files    []*models.File
	)

	rows, err := r.DB.QueryxContext(ctx, `
SELECT id, name, alt, caption, hash, mime, ext, size, width, height, provider, url, created_by_id, updated_by_id from files;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file File

		err = rows.Scan(&file.ID, &file.Name, &file.Alt, &file.Caption, &file.Hash, &file.Mime, &file.Ext, &file.Size,
			&file.Width, &file.Height, &file.Provider, &file.URL, &file.CreatedByID, &file.UpdatedByID)
		if err != nil {
			return nil, err
		}

		provider := ""
		if file.Provider != nil {
			provider = *file.Provider
		}

		files = append(files, &models.File{
			ID:          file.ID,
			Name:        file.Name + file.Ext,
			Thumb:       file.URL,
			Alt:         file.Alt,
			Caption:     file.Caption,
			Hash:        file.Hash,
			Type:        file.Mime,
			Ext:         file.Ext,
			Size:        file.Size,
			Width:       file.Width,
			Height:      file.Height,
			Provider:    provider,
			CreatedByID: file.CreatedByID,
			UpdatedByID: file.UpdatedByID,
		})
	}

	defer rows.Close()

	return files, multierr.ErrorOrNil()
}

func (r *FilesRepository) GetFileByName(ctx context.Context, name string) (*models.File, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var file File

	sqlStatement := `SELECT id, name, alt, caption, hash, mime, ext, size, width, height, provider, url, created_by_id, updated_by_id from files WHERE name=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, name)

	err := row.Scan(&file.ID, &file.Name, &file.Alt, &file.Caption, &file.Hash, &file.Mime, &file.Ext, &file.Size,
		&file.Width, &file.Height, &file.Provider, &file.URL, &file.CreatedByID, &file.UpdatedByID)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("file not found")
	case nil:
		provider := ""
		if file.Provider != nil {
			provider = *file.Provider
		}

		return &models.File{
			ID:          file.ID,
			Name:        file.Name + file.Ext,
			Thumb:       file.URL,
			Alt:         file.Alt,
			Caption:     file.Caption,
			Hash:        file.Hash,
			Type:        file.Mime,
			Ext:         file.Ext,
			Size:        file.Size,
			Width:       file.Width,
			Height:      file.Height,
			Provider:    provider,
			CreatedByID: file.CreatedByID,
			UpdatedByID: file.UpdatedByID,
		}, err
	default:
		return nil, err
	}
}

func (r *FilesRepository) GetFileByID(ctx context.Context, id int64) (*models.File, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var file File

	sqlStatement := `SELECT id, name, alt, caption, hash, mime, ext, size, width, height, provider, url, created_by_id, updated_by_id from files WHERE id=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&file.ID, &file.Name, &file.Alt, &file.Caption, &file.Hash, &file.Mime, &file.Ext, &file.Size,
		&file.Width, &file.Height, &file.Provider, &file.URL, &file.CreatedByID, &file.UpdatedByID)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("file not found")
	case nil:
		provider := ""
		if file.Provider != nil {
			provider = *file.Provider
		}

		return &models.File{
			ID:          file.ID,
			Name:        file.Name,
			Thumb:       file.URL,
			Alt:         file.Alt,
			Caption:     file.Caption,
			Hash:        file.Hash,
			Type:        file.Mime,
			Ext:         file.Ext,
			Size:        file.Size,
			Width:       file.Width,
			Height:      file.Height,
			Provider:    provider,
			CreatedByID: file.CreatedByID,
			UpdatedByID: file.UpdatedByID,
		}, err
	default:
		return nil, err
	}
}

func (r *FilesRepository) GetPublicFilesByProvider(ctx context.Context, provider string) ([]*models.PublicFile, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		files    []*models.PublicFile
	)

	rows, err := r.DB.QueryxContext(ctx, `
SELECT id, name, alt, caption, hash, mime, ext, size, width, height, provider, url from files where provider=$1;`, provider)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file File

		err = rows.Scan(&file.ID, &file.Name, &file.Alt, &file.Caption, &file.Hash, &file.Mime, &file.Ext, &file.Size,
			&file.Width, &file.Height, &file.Provider, &file.URL)
		if err != nil {
			return nil, err
		}

		provider := ""
		if file.Provider != nil {
			provider = *file.Provider
		}

		files = append(files, &models.PublicFile{
			ID: file.ID,
			Attributes: &models.PublicFileAttributes{
				Alt:      file.Alt,
				Caption:  file.Caption,
				Ext:      file.Ext,
				Hash:     file.Hash,
				Height:   file.Height,
				Width:    file.Width,
				Mime:     file.Mime,
				Name:     file.Name,
				Provider: provider,
				Size:     file.Size,
				URL:      file.URL,
			},
		})
	}
	defer rows.Close()

	return files, multierr.ErrorOrNil()
}

func (r *FilesRepository) GetPublicFileByName(ctx context.Context, name string) (*models.PublicFile, error) {
	file, err := r.GetFileByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &models.PublicFile{
		ID: file.ID,
		Attributes: &models.PublicFileAttributes{
			Alt:      file.Alt,
			Caption:  file.Caption,
			Ext:      file.Ext,
			Hash:     file.Hash,
			Height:   file.Height,
			Mime:     file.Type,
			Name:     file.Name,
			Provider: file.Provider,
			Size:     file.Size,
			URL:      file.Thumb,
			Width:    file.Width,
		},
	}, nil
}

func (r *FilesRepository) GetPublicFileByID(ctx context.Context, id int64) (*models.PublicFile, error) {
	file, err := r.GetFileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.PublicFile{
		ID: file.ID,
		Attributes: &models.PublicFileAttributes{
			Alt:      file.Alt,
			Caption:  file.Caption,
			Ext:      file.Ext,
			Hash:     file.Hash,
			Height:   file.Height,
			Mime:     file.Type,
			Name:     file.Name,
			Provider: file.Provider,
			Size:     file.Size,
			URL:      file.Thumb,
			Width:    file.Width,
		},
	}, nil
}
