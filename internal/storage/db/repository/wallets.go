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

type Wallet struct {
	Created     strfmt.Date `db:"created"`
	Updated     strfmt.Date `db:"updated"`
	Deleted     strfmt.Date `db:"deleted"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	ID          int64       `db:"id"`
	ImageID     int64       `db:"image_id"`
	OrderBy     int64       `db:"order_by"`
	CreatedBy   int64       `db:"created_by"`
	UpdatedBy   int64       `db:"updated_by"`
	Enabled     bool        `db:"enabled"`
}

type WalletsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, walletMap map[string]interface{}) error
	Update(ctx context.Context, walletMap map[string]interface{}) ([]*models.WalletConnect, error)
	GetWallets(ctx context.Context) ([]*models.WalletConnect, error)
	GetWalletByID(ctx context.Context, id int64) (*models.WalletConnect, error)
	GetPublicWalletConnect(ctx context.Context) ([]*models.PublicWalletConnectItem, error)
}

// WalletsRepository - repository to store users.
type WalletsRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewWalletsRepository(db *sqlx.DB, tbl string) *WalletsRepository {
	if db == nil {
		return nil
	}

	return &WalletsRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *WalletsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *WalletsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *WalletsRepository) Create(ctx context.Context, walletMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO wallets (title, description, image_id, enabled, order_by, created_by_id, updated_by_id)
VALUES (:title, :description, :image_id, :enabled, :order_by, :created_by_id, :updated_by_id)
ON CONFLICT DO NOTHING`, walletMap)

	return err
}

func (r *WalletsRepository) Update(ctx context.Context, walletMap map[string]interface{}) ([]*models.WalletConnect, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
UPDATE wallets SET
	title=:title,
	description=:description,
	image_id=:image_id,
	enabled=:enabled,
	order_by=:order_by,
	updated_by_id=:updated_by_id WHERE id =:id`, walletMap)
	if err != nil {
		return nil, err
	}

	return r.GetWallets(ctx)
}

func (r *WalletsRepository) GetWallets(ctx context.Context) ([]*models.WalletConnect, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		wallets  []*models.WalletConnect
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id from wallets;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var wallet Wallet

		err = rows.Scan(&wallet.ID, &wallet.Title, &wallet.Description,
			&wallet.ImageID, &wallet.Enabled, &wallet.OrderBy, &wallet.CreatedBy, &wallet.UpdatedBy)
		if err != nil {
			return nil, err
		}

		wallets = append(wallets, &models.WalletConnect{
			ID:          wallet.ID,
			Title:       wallet.Title,
			Description: wallet.Description,
			ImageID:     wallet.ImageID,
			Enabled:     wallet.Enabled,
			OrderBy:     wallet.OrderBy,
			CreatedByID: wallet.CreatedBy,
			UpdatedByID: wallet.UpdatedBy,
		})
	}

	defer rows.Close()

	return wallets, multierr.ErrorOrNil()
}

func (r *WalletsRepository) GetWalletByID(ctx context.Context, id int64) (*models.WalletConnect, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var wallet Wallet

	sqlStatement := `SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id from wallets where id=$1;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, id)
	if err := row.Scan(&wallet.ID, &wallet.Title, &wallet.Description,
		&wallet.ImageID, &wallet.Enabled, &wallet.OrderBy, &wallet.CreatedBy, &wallet.UpdatedBy); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("wallet not found")
		}

		return nil, NewError(err)
	}

	return &models.WalletConnect{
		ID:          wallet.ID,
		Title:       wallet.Title,
		Description: wallet.Description,
		ImageID:     wallet.ImageID,
		Enabled:     wallet.Enabled,
		OrderBy:     wallet.OrderBy,
		CreatedByID: wallet.CreatedBy,
		UpdatedByID: wallet.UpdatedBy,
	}, nil
}

func (r *WalletsRepository) GetPublicWalletConnect(ctx context.Context) ([]*models.PublicWalletConnectItem, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr multierror.Error
		wallets  []*models.PublicWalletConnectItem
	)

	rows, err := r.DB.QueryxContext(ctx,
		`SELECT 
			w.title, 
			w.description, 
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
								) FROM files f WHERE f.id = w.image_id)
			) FROM files f WHERE f.id = w.image_id) as image
		FROM wallets w
		WHERE
			w.enabled = true AND w.deleted isNULL  GROUP BY w.id ORDER BY w.order_by ASC;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var wallet Wallet

		var image AggregatedImageJSON

		err = rows.Scan(&wallet.Title, &wallet.Description, &image)
		if err != nil {
			return nil, err
		}

		wallets = append(wallets, &models.PublicWalletConnectItem{
			ID: wallet.ID,
			Attributes: &models.PublicWalletConnectItemAttributes{
				Title:       wallet.Title,
				Description: wallet.Description,
				Image: &models.PublicFile{
					Attributes: image.Attributes,
					ID:         image.ID,
				},
			},
		})
	}

	defer rows.Close()

	return wallets, multierr.ErrorOrNil()
}
