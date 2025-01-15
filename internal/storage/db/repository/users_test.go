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

// TestUsersRepository_GetUserByEmail tests the GetUserByEmail method of UsersRepository.
func TestUsersRepository_GetUserByEmail(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	email := "test@example.com"

	// Test successful GetUserByEmail
	rows := sqlmock.NewRows([]string{"id", "email", "password", "password_reset_token", "password_reset_expires", "enabled", "secret", "rsa_secret", "created", "updated", "deleted"}).
		AddRow(1, email, "hashedpassword", "resetToken", 1234567890, true, "mysecret", "auth-token", "2024-01-01", "2024-01-01", "2024-01-01")

	// Updated query pattern to be more flexible with whitespace and new lines
	queryPattern := `(?i)SELECT\s+id,\s+email,\s+password,\s+password_reset_token,\s+password_reset_expires,\s+enabled,\s+secret,\s+rsa_secret,\s+created,\s+updated,\s+deleted\s+FROM\s+users\s+WHERE\s+email=\$1;`

	mock.ExpectQuery(queryPattern).
		WithArgs(email).
		WillReturnRows(rows)

	user, err := repo.GetUserByEmail(ctx, email)
	assert.NoError(t, err, "GetUserByEmail() should not return an error")
	assert.NotNil(t, user, "GetUserByEmail() should return a user")
	assert.Equal(t, email, user.Email, "GetUserByEmail() should return the correct user")

	// Test GetUserByEmail when user not found
	mock.ExpectQuery(queryPattern).
		WithArgs(email).
		WillReturnError(sql.ErrNoRows) // Simulate user not found error

	user, err = repo.GetUserByEmail(ctx, email)
	assert.Error(t, err, "GetUserByEmail() should return an error when user is not found")

	// Check if the error is wrapped by NewError
	var customErr *repository.Error

	assert.True(t, errors.As(err, &customErr), "GetUserByEmail() should return a NewError when user is not found")
	assert.Nil(t, user, "GetUserByEmail() should return nil user when not found")

	// Test GetUserByEmail with database error
	mock.ExpectQuery(queryPattern).
		WithArgs(email).
		WillReturnError(repository.NewError(errors.New("user not found")))

	user, err = repo.GetUserByEmail(ctx, email)
	assert.Error(t, err, "GetUserByEmail() should return an error when there's a database error")
	assert.Nil(t, user, "GetUserByEmail() should return nil user when there's a database error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	user, err = nilRepo.GetUserByEmail(ctx, email)
	assert.Error(t, err, "GetUserByEmail() on nil repository should return an error")
	assert.Nil(t, user, "GetUserByEmail() on nil repository should return nil user")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetUserByEmail() should return ErrDBNotInitialized")
}

func TestUsersRepository_GetPublicUserByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	email := "test@example.com"
	userID := int64(1)
	uuid := "21e49d82-5240-423f-8d92-1ffd4a1cf600"
	// Test successful GetUserByEmail
	rows := sqlmock.NewRows([]string{"id", "uuid", "username", "email", "address"}).
		AddRow(userID, uuid, "testuser", email, "0x1234567890abcdef")

	// Updated query pattern to be more flexible with whitespace and new lines
	queryPattern := `(?i)SELECT\s+u\.id,\s+a\.uuid,\s+u\.username,\s+u\.email,\s+a\.address\s+FROM\s+users\s+u\s+JOIN\s+users_addresses\s+a\s+ON\s+a\.user_id\s+=\s+u\.id\s+LEFT\s+JOIN\s+some_table\s+f\s+ON\s+f\.user_id\s+=\s+u\.id\s+WHERE\s+u\.deleted\s+IS\s+NULL\s+AND\s+u\.id\s+=\s+\$1\s+AND\s+u\.deleted\s+IS\s+NULL;`

	mock.ExpectQuery(queryPattern).
		WithArgs(userID).
		WillReturnRows(rows)

	user, err := repo.GetPublicUserByID(ctx, userID)
	assert.NoError(t, err, "GetPublicUserByID() should not return an error")
	assert.NotNil(t, user, "GetPublicUserByID() should return a user")
	assert.Equal(t, userID, user.ID, "GetPublicUserByID() should return the correct user ID")
	assert.Equal(t, uuid, user.UUID, "GetPublicUserByID() should return the correct UUID")
	assert.Equal(t, "testuser", user.Username, "GetPublicUserByID() should return the correct username")
	assert.Equal(t, email, user.Email, "GetPublicUserByID() should return the correct email")
	assert.Equal(t, "0x1234567890abcdef", user.Address, "GetPublicUserByID() should return the correct address")

	// Case: User not found
	mock.ExpectQuery(queryPattern).
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	user, err = repo.GetPublicUserByID(ctx, userID)
	assert.Error(t, err, "GetPublicUserByID() should return an error when user is not found")
	assert.Nil(t, user, "GetPublicUserByID() should return nil when user is not found")

	// Case: Database error
	mock.ExpectQuery(queryPattern).
		WithArgs(userID).
		WillReturnError(errors.New("database error"))

	user, err = repo.GetPublicUserByID(ctx, userID)
	assert.Error(t, err, "GetPublicUserByID() should return an error on database error")
	assert.Nil(t, user, "GetPublicUserByID() should return nil on database error")

	// Case: Nil repository
	var nilRepo *repository.UsersRepository
	user, err = nilRepo.GetPublicUserByID(ctx, userID)
	assert.Error(t, err, "GetPublicUserByID() on nil repository should return an error")
	assert.Nil(t, user, "GetPublicUserByID() on nil repository should return nil")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetPublicUserByID() should return ErrDBNotInitialized")
}

// TestUsersRepository_GetUserByAddress tests the GetUserByAddress method of UsersRepository.
func TestUsersRepository_GetUserByAddress(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	address := "0x1234567890abcdef"

	// Test successful GetUserByAddress
	rows := sqlmock.NewRows([]string{"id", "uuid", "username", "email", "password", "address", "nonce"}).
		AddRow(1, "uuid_value", "testuser", "testuser@example.com", "hashedpassword", address, "nonce_value")

	queryPattern := `(?i)SELECT\s+u\.id,\s+a\.uuid,\s+u\.username,\s+u\.email,\s+u\.password,\s+a\.address,\s+a\.nonce\s+FROM\s+users\s+u\s+LEFT\s+JOIN\s+users_addresses\s+a\s+ON\s+a\.user_id\s+=\s+u\.id\s+WHERE\s+a\.address=\$1\s+AND\s+u\.deleted\s+IS\s+NULL;`

	mock.ExpectQuery(queryPattern).
		WithArgs(address).
		WillReturnRows(rows)

	user, err := repo.GetUserByAddress(ctx, address)
	assert.NoError(t, err, "GetUserByAddress() should not return an error")
	assert.NotNil(t, user, "GetUserByAddress() should return a user")
	assert.Equal(t, address, user.Address, "GetUserByAddress() should return the correct address")
	assert.Equal(t, "testuser", user.Username, "GetUserByAddress() should return the correct username")

	// Test GetUserByAddress when user is not found
	mock.ExpectQuery(queryPattern).
		WithArgs(address).
		WillReturnError(sql.ErrNoRows)

	user, err = repo.GetUserByAddress(ctx, address)
	assert.Error(t, err, "GetUserByAddress() should return an error when user is not found")
	assert.Nil(t, user, "GetUserByAddress() should return nil when user is not found")

	// Check if the error is the expected "user not found"
	assert.EqualError(t, err, "user not found")

	// Test GetUserByAddress with database error
	mock.ExpectQuery(queryPattern).
		WithArgs(address).
		WillReturnError(errors.New("database error"))

	user, err = repo.GetUserByAddress(ctx, address)
	assert.Error(t, err, "GetUserByAddress() should return an error when there's a database error")
	assert.Nil(t, user, "GetUserByAddress() should return nil when there's a database error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	user, err = nilRepo.GetUserByAddress(ctx, address)
	assert.Error(t, err, "GetUserByAddress() on nil repository should return an error")
	assert.Nil(t, user, "GetUserByAddress() on nil repository should return nil")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetUserByAddress() should return ErrDBNotInitialized")
}
