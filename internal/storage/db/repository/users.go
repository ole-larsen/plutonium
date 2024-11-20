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
	_ "github.com/lib/pq" // add lib
	"github.com/ole-larsen/plutonium/internal/hash"
)

// ErrDBNotInitialized - default database connection error.
var ErrDBNotInitialized = fmt.Errorf("db not initialised")

// UsersRepositoryInterface - interface to store users.
type UsersRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	MigrateContext(ctx context.Context) error
	Create(ctx context.Context, userMap map[string]interface{}) error
	GetOne(ctx context.Context, email string) (*User, error)
}

type User struct {
	Created              strfmt.Date `db:"created"`
	Updated              strfmt.Date `db:"updated"`
	Deleted              strfmt.Date `db:"deleted"`
	PasswordResetToken   *string     `db:"password_reset_token"`
	PasswordResetExpires *int64      `db:"password_reset_expires"`
	RSASecret            string      `db:"rsa_secret"`
	Email                string      `db:"email"`
	Password             string      `db:"password"`
	Secret               string      `db:"secret"`
	ID                   int64       `db:"id"`
	Enabled              bool        `db:"enabled"`
}

// UsersRepository - repository to store users.
type UsersRepository struct {
	DB  sqlx.DB
	TBL string
}

func SetPassword(password interface{}) (string, error) {
	var pwd string

	if plainPwd, ok := password.(string); ok {
		if plainPwd == "" {
			return pwd, errors.New("empty password not allowed")
		}

		return hash.Password([]byte(plainPwd))
	}

	return pwd, errors.New("password must be a string")
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

func (r *UsersRepository) MigrateContext(ctx context.Context) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id                      SERIAL PRIMARY KEY,
	    email                   VARCHAR(255),
    	password                VARCHAR(255) NOT NULL,
		password_reset_token    VARCHAR(255),
		password_reset_expires  BIGINT,
		enabled                 bool NOT NULL DEFAULT TRUE,
		secret                  VARCHAR(255),
    	rsa_secret              VARCHAR(255),
    	created                 TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		updated                 TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		deleted                 TIMESTAMP WITH TIME ZONE DEFAULT NULL
	);`, r.TBL))

	return err
}

func (r *UsersRepository) Create(ctx context.Context, userMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	var err error
	if userMap["password"], err = SetPassword(userMap["password"]); err != nil {
		return NewError(err)
	}
	// JNUGNHA27JMIHA5I
	// generate secret per user
	length := 16
	userMap["rsa_secret"] = hash.RandStringBytes(length)

	_, dbErr := r.DB.NamedExecContext(ctx, fmt.Sprintf(`
		INSERT INTO %s (email, password, secret, rsa_secret)
		VALUES (:email, :password, :secret, :rsa_secret)
		ON CONFLICT DO NOTHING`, r.TBL), userMap)

	return dbErr
}

func (r *UsersRepository) GetOne(ctx context.Context, email string) (*User, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var user User

	sqlStatement := fmt.Sprintf(`SELECT 
		id,
	    email,
    	password,
		password_reset_token,
		password_reset_expires,
		enabled,
		secret,
		rsa_secret,
		created,
		updated,
		deleted
	from %s WHERE email=$1;`, r.TBL)
	row := r.DB.QueryRowxContext(ctx, sqlStatement, email)

	err := row.StructScan(&user)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewError(fmt.Errorf("user not found"))
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
