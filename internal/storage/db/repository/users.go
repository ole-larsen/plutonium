// Package repository contains all database logic for storage
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	Update(ctx context.Context, userMap map[string]interface{}) (*User, error)
}

type User struct {
	Deleted              strfmt.Date          `db:"deleted"`
	Created              strfmt.Date          `db:"created"`
	Updated              strfmt.Date          `db:"updated"`
	Password             string               `db:"password"`
	Wallpaper            *AggregatedImageJSON `db:"wallpaper"`
	Username             string               `db:"username"`
	RSASecret            string               `db:"rsa_secret"`
	Email                string               `db:"email"`
	Gravatar             string               `db:"gravatar"`
	Secret               string               `db:"secret"`
	Address              pq.StringArray       `db:"address"`
	Nonce                sql.NullString       `db:"nonce"`
	UUID                 sql.NullString       `db:"uuid"`
	PasswordResetToken   sql.NullString       `db:"password_reset_token"`
	PasswordResetExpires sql.NullInt64        `db:"password_reset_expires"`
	WallpaperID          sql.NullInt64        `db:"wallpaper_id"`
	ID                   int64                `db:"id"`
	Enabled              bool                 `db:"enabled"`
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
	u.id,
	u.email,
	u.password,
	u.password_reset_token,
	u.password_reset_expires,
	u.enabled,
	u.secret,
	u.rsa_secret,
	u.gravatar,
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
		) FROM files f WHERE f.id = u.wallpaper_id)
		) FROM files f WHERE f.id = u.wallpaper_id) as wallpaper,
	u.created,
	u.updated,
	u.deleted
FROM %s u
LEFT JOIN files f ON u.wallpaper_id = f.id
WHERE u.id=$1;`, r.TBL)

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
		) FROM files f WHERE f.id = u.wallpaper_id)
		) FROM files f WHERE f.id = u.wallpaper_id) as wallpaper,
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
		) FROM files f WHERE f.id = u.wallpaper_id)
		) FROM files f WHERE f.id = u.wallpaper_id) as wallpaper
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

	publicUser := &models.PublicUser{
		ID: user.ID,
		Attributes: &models.PublicUserAttributes{
			Username: user.Username,
			Email:    user.Email,
			Gravatar: user.Gravatar,
		},
	}
	if user.Wallpaper != nil {
		publicUser.Attributes.Wallpaper = &models.PublicFile{
			Attributes: user.Wallpaper.Attributes,
			ID:         user.Wallpaper.ID,
		}
	}

	return publicUser, nil
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
		) FROM files f WHERE f.id = u.wallpaper_id)
		) FROM files f WHERE f.id = u.wallpaper_id) as wallpaper
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

func (r *UsersRepository) Update(ctx context.Context, userMap map[string]interface{}) (*User, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	id, ok := userMap["id"]
	if !ok {
		return nil, fmt.Errorf("missing 'id' in userMap")
	}

	delete(userMap, "id") // ID should not be updated

	// Build the SET clause dynamically
	var setClauses []string
	var args []interface{}
	argIndex := 1
	for key, value := range userMap {
		setClauses = append(setClauses, fmt.Sprintf("%s=$%d", key, argIndex))
		args = append(args, value)
		argIndex++
	}

	args = append(args, id) // Append the user ID for WHERE clause

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d RETURNING *", strings.Join(setClauses, ", "), argIndex)
	fmt.Println(query, args)
	// Execute the query and scan the result into updatedUser
	var updatedUser User
	row := r.DB.QueryRowxContext(ctx, query, args...)
	if err := row.StructScan(&updatedUser); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &updatedUser, nil
}
