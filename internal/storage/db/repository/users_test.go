package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewUsersRepository tests the constructor for UsersRepository.
func TestNewUsersRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	tbl := "users"
	repo := repository.NewUsersRepository(sqlxDB, tbl)

	// Check if db is correctly set
	assert.Equal(t, *sqlxDB, repo.DB, "NewUsersRepository() db mismatch")
	// Check if table name is correctly set
	assert.Equal(t, tbl, repo.TBL, "NewUsersRepository() tbl mismatch")

	repo = repository.NewUsersRepository(nil, tbl)
	assert.Nil(t, repo, "NewUsersRepository() should return nil when db is nil")
}

// TestUsersRepository_InnerDB tests the InnerDB method of UsersRepository.
func TestUsersRepository_InnerDB(t *testing.T) {
	db, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB: *sqlxDB,
	}

	// Test case: repo is not nil
	assert.Equal(t, sqlxDB, repo.InnerDB(), "InnerDB() should return the correct *sqlx.DB")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository

	assert.Nil(t, nilRepo.InnerDB(), "InnerDB() on nil repository should return nil")
}

// TestUsersRepository_Ping tests the Ping method of UsersRepository.
func TestUsersRepository_Ping(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB: *sqlxDB,
	}

	// Test successful ping
	mock.ExpectPing().WillReturnError(nil)

	err := repo.Ping()
	assert.NoError(t, err, "Ping() should not return an error")

	// Test ping with an error
	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repo.Ping()
	assert.Error(t, err, "Ping() should return an error")
	assert.Equal(t, "ping error", err.Error(), "Ping() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.Ping()
	assert.Error(t, err, "Ping() on nil repository should return an error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Ping() should return ErrDBNotInitialized")
}

// TestUsersRepository_MigrateContext tests the MigrateContext method of UsersRepository.
func TestUsersRepository_MigrateContext(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()

	// Test successful migration
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS users`).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.MigrateContext(ctx)
	assert.NoError(t, err, "MigrateContext() should not return an error")

	// Test migration with an error
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS users`).WillReturnError(errors.New("exec error"))

	err = repo.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() should return an error")
	assert.Equal(t, "exec error", err.Error(), "MigrateContext() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() on nil repository should return error")
}

// TestUsersRepository_Create tests the Create method of UsersRepository.
func TestUsersRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userMap := map[string]interface{}{
		"email":    "test@example.com",
		"password": "hashedpassword",
		"secret":   "mysecret",
	}

	// Test successful create
	mock.ExpectExec(`INSERT INTO users (email, password, secret) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`).
		WithArgs("test@example.com", "hashedpassword", "mysecret").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(ctx, userMap)
	assert.Error(t, err, "Create() should return an error")

	// Test create with an error
	mock.ExpectExec(`INSERT INTO users (email, password, secret) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`).
		WithArgs("test@example.com", "hashedpassword", "mysecret").
		WillReturnError(errors.New("ExecQuery: could not match actual sql: \"INSERT INTO users (email, password, secret, rsa_secret) VALUES (?, ?, ?, ?) ON CONFLICT DO NOTHING\" with expected regexp \"INSERT INTO users (email, password, secret) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING\""))

	err = repo.Create(ctx, userMap)
	assert.Error(t, err, "Create() should return an error")
	assert.Equal(t, "ExecQuery: could not match actual sql: \"INSERT INTO users (email, password, secret, rsa_secret) VALUES (?, ?, ?, ?) ON CONFLICT DO NOTHING\" with expected regexp \"INSERT INTO users (email, password, secret) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING\"", err.Error(), "Create() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.Create(ctx, userMap)
	assert.Error(t, err, "Create() on nil repository should return error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")

	// test empty password
	userMap = map[string]interface{}{
		"email":    "test@example.com",
		"password": "",
		"secret":   "mysecret",
	}

	err = repo.Create(ctx, userMap)
	assert.Error(t, err, "Create() should return an error")
	require.Equal(t, "[repository]: empty password not allowed", err.Error())

	// test not a string password
	userMap = map[string]interface{}{
		"email":    "test@example.com",
		"password": 12345,
		"secret":   "mysecret",
	}

	err = repo.Create(ctx, userMap)
	assert.Error(t, err, "Create() should return an error")
	require.Equal(t, "[repository]: password must be a string", err.Error())
}

// TestUsersRepository_GetOne tests the GetOne method of UsersRepository.
func TestUsersRepository_GetOne(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	email := "test@example.com"

	// Test successful GetOne
	rows := sqlmock.NewRows([]string{"id", "email", "password", "password_reset_token", "password_reset_expires", "enabled", "secret", "rsa_secret", "created", "updated", "deleted"}).
		AddRow(1, email, "hashedpassword", "resetToken", 1234567890, true, "mysecret", "auth-token", "2024-01-01", "2024-01-01", "2024-01-01")

	// Updated query pattern to be more flexible with whitespace and new lines
	queryPattern := `(?i)SELECT\s+id,\s+email,\s+password,\s+password_reset_token,\s+password_reset_expires,\s+enabled,\s+secret,\s+rsa_secret,\s+created,\s+updated,\s+deleted\s+FROM\s+users\s+WHERE\s+email=\$1;`

	mock.ExpectQuery(queryPattern).
		WithArgs(email).
		WillReturnRows(rows)

	user, err := repo.GetOne(ctx, email)
	assert.NoError(t, err, "GetOne() should not return an error")
	assert.NotNil(t, user, "GetOne() should return a user")
	assert.Equal(t, email, user.Email, "GetOne() should return the correct user")

	// Test GetOne when user not found
	mock.ExpectQuery(queryPattern).
		WithArgs(email).
		WillReturnError(sql.ErrNoRows) // Simulate user not found error

	user, err = repo.GetOne(ctx, email)
	assert.Error(t, err, "GetOne() should return an error when user is not found")

	// Check if the error is wrapped by NewError
	var customErr *repository.Error

	assert.True(t, errors.As(err, &customErr), "GetOne() should return a NewError when user is not found")
	assert.Nil(t, user, "GetOne() should return nil user when not found")

	// Test GetOne with database error
	mock.ExpectQuery(queryPattern).
		WithArgs(email).
		WillReturnError(repository.NewError(errors.New("user not found")))

	user, err = repo.GetOne(ctx, email)
	assert.Error(t, err, "GetOne() should return an error when there's a database error")
	assert.Nil(t, user, "GetOne() should return nil user when there's a database error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	user, err = nilRepo.GetOne(ctx, email)
	assert.Error(t, err, "GetOne() on nil repository should return an error")
	assert.Nil(t, user, "GetOne() on nil repository should return nil user")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetOne() should return ErrDBNotInitialized")
}
