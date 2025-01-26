// Package repository contains all database logic for storage
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-openapi/strfmt"
	_ "github.com/go-sql-driver/mysql" // add driver
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ole-larsen/plutonium/models"
)

// ErrDBNotInitialized - default database connection error.
var ErrDBNotInitialized = fmt.Errorf("db not initialised")

// UsersRepositoryInterface - interface to store users.
type UsersRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	Create(ctx context.Context, userMap map[string]interface{}) error
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetPublicUserByID(ctx context.Context, id int64) (*models.PublicUser, error)
	UpdateNonce(ctx context.Context, userMap map[string]interface{}) error
	UpdateGravatar(ctx context.Context, userMap map[string]interface{}) error
	UpdateWallpaper(ctx context.Context, userMap map[string]interface{}) error
	UpdateSecret(ctx context.Context, userMap map[string]interface{}) error
}

type User struct {
	Deleted              strfmt.Date    `db:"deleted"`
	Created              strfmt.Date    `db:"created"`
	Updated              strfmt.Date    `db:"updated"`
	Gravatar             string         `db:"gravatar"`
	Secret               string         `db:"secret"`
	Username             string         `db:"username"`
	RSASecret            string         `db:"rsa_secret"`
	Email                string         `db:"email"`
	Password             string         `db:"password"`
	PasswordResetToken   sql.NullString `db:"password_reset_token"`
	Address              pq.StringArray `db:"address"`
	Nonce                sql.NullString `db:"nonce"`
	UUID                 sql.NullString `db:"uuid"`
	Wallpaper            sql.NullString `db:"wallpaper"`
	PasswordResetExpires sql.NullInt64  `db:"password_reset_expires"`
	WallpaperID          sql.NullInt64  `db:"wallpaper_id"`
	ID                   int64          `db:"id"`
	Enabled              bool           `db:"enabled"`
}

// UsersRepository - repository to store users.
type UsersRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewUsersRepository(db *sqlx.DB, tbl string) *UsersRepository {
	if db == nil {
		return nil
	}

	return &UsersRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *UsersRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *UsersRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *UsersRepository) Create(ctx context.Context, userMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, dbErr := r.DB.NamedExecContext(ctx, fmt.Sprintf(`
INSERT INTO %s (email, username, password, secret, rsa_secret, address, nonce, gravatar)
VALUES (:email, :username, :password, :secret, :rsa_secret, :address, :nonce, :gravatar)
ON CONFLICT DO NOTHING`, r.TBL), userMap)

	return dbErr
}

func (r *UsersRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var user User

	sqlStatement := fmt.Sprintf(`
SELECT 
	id,
	email,
	password,
	password_reset_token,
	password_reset_expires,
	enabled,
	secret,
	rsa_secret,
	gravatar,
	f.url as wallpaper,
	created,
	updated,
	deleted
FROM %s 
LEFT JOIN files f ON wallpaper_id = f.id
WHERE id=$1;`, r.TBL)

	row := r.DB.QueryRowxContext(ctx, sqlStatement, id)

	err := row.StructScan(&user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("user not found"))
		}

		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var user User

	sqlStatement := fmt.Sprintf(`
SELECT 
	id,
	email,
	password,
	password_reset_token,
	password_reset_expires,
	enabled,
	secret,
	rsa_secret,
	gravatar,
    f.url as wallpaper,
	created,
	updated,
	deleted
FROM %s 
LEFT JOIN files f ON wallpaper_id = f.id
WHERE email=$1;`, r.TBL)

	row := r.DB.QueryRowxContext(ctx, sqlStatement, email)

	err := row.StructScan(&user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("user not found"))
		}

		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) GetPublicUserByID(ctx context.Context, id int64) (*models.PublicUser, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	row := r.DB.QueryRowContext(ctx, `
SELECT 
    u.id, 
    u.uuid, 
    u.username, 
    u.email
	u.gravatar,
	f.url as wallpaper
FROM users u
LEFT JOIN files f ON u.wallpaper_id = f.id
WHERE u.deleted IS NULL
  AND u.id = $1
  AND u.deleted IS NULL;`, id)

	var user User

	if err := row.Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Gravatar, &user.Wallpaper); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("user not found"))
		}
		return nil, err
	}

	wallpaper := ""
	if user.Wallpaper.Valid {
		wallpaper = user.Wallpaper.String
	}
	return &models.PublicUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Gravatar:  user.Gravatar,
		Wallpaper: wallpaper,
	}, nil
}

func (r *UsersRepository) GetUserByAddress(ctx context.Context, address string) (*User, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var user User

	sqlStatement := `
SELECT 
	u.id, 
	u.uuid,
	u.username, 
	u.email, 
	u.password,
	u.address,
	u.nonce,
	u.gravatar,
    f.url as wallpaper
FROM users u
LEFT JOIN files f ON u.wallpaper_id = f.id
WHERE $1 = ANY(u.address) AND u.deleted IS NULL;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, address)

	if err := row.Scan(&user.ID, &user.UUID, &user.Username, &user.Email, &user.Password, &user.Address, &user.Nonce, &user.Gravatar, &user.Wallpaper); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewError(fmt.Errorf("user not found"))
		}

		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) UpdateNonce(ctx context.Context, userMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE users SET nonce=:nonce WHERE id=:id`, userMap)

	return err
}

func (r *UsersRepository) UpdateGravatar(ctx context.Context, userMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE users SET gravatar=:gravatar WHERE id =:id`, userMap)

	return err
}

func (r *UsersRepository) UpdateWallpaper(ctx context.Context, userMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE users SET wallpaper_id=:wallpaper_id WHERE id =:id`, userMap)

	return err
}
func (r *UsersRepository) UpdateSecret(ctx context.Context, userMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `UPDATE users SET secret=:secret WHERE id =:id`, userMap)

	return err
}
