package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/stretchr/testify/assert"
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
		"email":      "test@example.com",
		"username":   "username",
		"password":   "hashedpassword",
		"secret":     "mysecret",
		"rsa_secret": "rsa_secret",
		"address":    "address",
		"nonce":      "nonce",
		"gravatar":   "gravatar",
	}

	// Test successful create
	mock.ExpectExec(`INSERT INTO users (email, username, password, secret, rsa_secret, address, nonce, gravatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING`).
		WithArgs("test@example.com", "username", "hashedpassword", "mysecret", "rsa_secret", "address", "nonce", "gravatar").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(ctx, userMap)
	assert.Error(t, err, "Create() should return an error")

	// Test create with an error
	mock.ExpectExec(`INSERT INTO users (email, username, password, secret, rsa_secret, address, nonce, geavatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING`).
		WithArgs("test@example.com", "username", "hashedpassword", "mysecret", "rsa_secret", "address", "nonce", "gravatar").
		WillReturnError(errors.New("ExecQuery: could not match actual sql: \"INSERT INTO users (email, username, password, secret, rsa_secret, address, nonce, gravatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING\" with expected regexp \"INSERT INTO users (email, username, password, secret, rsa_secret, address, nonce, gravatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING\""))

	err = repo.Create(ctx, userMap)
	assert.Error(t, err, "Create() should return an error")
	assert.Equal(t, "ExecQuery: could not match actual sql: \"INSERT INTO users (email, username, password, secret, rsa_secret, address, nonce, gravatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING\" with expected regexp \"INSERT INTO users (email, username, password, secret, rsa_secret, address, nonce, gravatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING\"", err.Error(), "Create() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.Create(ctx, userMap)
	assert.Error(t, err, "Create() on nil repository should return error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}
func TestUsersRepository_GetUserByID_Errors(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()

	// Test case: Repo is nil
	var nilRepo *repository.UsersRepository
	_, err := nilRepo.GetUserByID(ctx, 1)
	assert.Error(t, err, "GetUserByID() on nil repository should return an error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetUserByID() should return ErrDBNotInitialized")

	// Test case: User not found (sql.ErrNoRows)
	mock.ExpectQuery(`SELECT .* FROM users LEFT JOIN files f ON wallpaper_id = f.id WHERE id=\$1`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	user, err := repo.GetUserByID(ctx, 1)
	assert.Error(t, err, "GetUserByID() should return an error")
	assert.Nil(t, user, "GetUserByID() should return nil user when not found")
	assert.Equal(t, "[repository]: user not found", err.Error(), "GetUserByID() should return correct error message")

	// Test case: Unexpected database error (query failure)
	mock.ExpectQuery(`SELECT .* FROM users LEFT JOIN files f ON wallpaper_id = f.id WHERE id=\$1`).
		WithArgs(1).
		WillReturnError(errors.New("database query failed"))

	user, err = repo.GetUserByID(ctx, 1)
	assert.Error(t, err, "GetUserByID() should return an error")
	assert.Nil(t, user, "GetUserByID() should return nil user on unexpected error")
	assert.Equal(t, "database query failed", err.Error(), "GetUserByID() should return correct error message")

	// Test case: Successful retrieval of user
	mock.ExpectQuery(`SELECT .* FROM users LEFT JOIN files f ON wallpaper_id = f.id WHERE id=\$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "email", "password", "password_reset_token", "password_reset_expires",
			"enabled", "secret", "rsa_secret", "gravatar", "wallpaper", "created", "updated", "deleted"}).
			AddRow(1, "test@example.com", "hashedpassword", "reset_token", 1234567890, true, "mysecret", "rsa_secret", "gravatar", "wallpaper_url", "2025-01-01", "2025-01-02", nil))

	user, err = repo.GetUserByID(ctx, 1)
	assert.NoError(t, err, "GetUserByID() should not return an error")
	assert.NotNil(t, user, "GetUserByID() should return a user")
	assert.Equal(t, int64(1), user.ID, "GetUserByID() should return correct user ID")
	assert.Equal(t, "test@example.com", user.Email, "GetUserByID() should return correct email")
	assert.Equal(t, "hashedpassword", user.Password, "GetUserByID() should return correct password")
	assert.Equal(t, "gravatar", user.Gravatar, "GetUserByID() should return correct gravatar")

}

func TestUsersRepository_GetUserByID_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userID := int64(1)

	// Simulate error in query execution
	mock.ExpectQuery(`SELECT`).
		WithArgs(userID).
		WillReturnError(errors.New("query execution error"))

	user, err := repo.GetUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	user, err = nilRepo.GetUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestUsersRepository_GetUserByEmail_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	email := "test@example.com"

	// Simulate error in query execution
	mock.ExpectQuery(`SELECT`).
		WithArgs(email).
		WillReturnError(errors.New("query execution error"))

	user, err := repo.GetUserByEmail(ctx, email)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	user, err = nilRepo.GetUserByEmail(ctx, email)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")

	// Test case: User not found (sql.ErrNoRows)
	mock.ExpectQuery(`SELECT .* FROM users LEFT JOIN files f ON wallpaper_id = f.id WHERE email=\$1`).
		WithArgs("test@example.com").
		WillReturnError(sql.ErrNoRows)

	user, err = repo.GetUserByEmail(ctx, "test@example.com")
	assert.Error(t, err, "GetUserByEmail() should return an error when user not found")
	assert.Nil(t, user, "GetUserByEmail() should return nil user when not found")
	assert.Equal(t, "[repository]: user not found", err.Error(), "GetUserByEmail() should return correct error message")

	// Test case: Unexpected database error (query failure)
	mock.ExpectQuery(`SELECT .* FROM users LEFT JOIN files f ON wallpaper_id = f.id WHERE email=\$1`).
		WithArgs("test@example.com").
		WillReturnError(errors.New("database query failed"))

	user, err = repo.GetUserByEmail(ctx, "test@example.com")
	assert.Error(t, err, "GetUserByEmail() should return an error on unexpected database failure")
	assert.Nil(t, user, "GetUserByEmail() should return nil user on query failure")
	assert.Equal(t, "database query failed", err.Error(), "GetUserByEmail() should return correct error message")

	// Test case: Successful retrieval of user by email
	mock.ExpectQuery(`SELECT .* FROM users LEFT JOIN files f ON wallpaper_id = f.id WHERE email=\$1`).
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "email", "password", "password_reset_token", "password_reset_expires",
			"enabled", "secret", "rsa_secret", "gravatar", "wallpaper", "created", "updated", "deleted"}).
			AddRow(1, "test@example.com", "hashedpassword", "reset_token", 1234567890, true, "mysecret", "rsa_secret", "gravatar", "wallpaper_url", "2025-01-01", "2025-01-02", nil))

	user, err = repo.GetUserByEmail(ctx, "test@example.com")
	assert.NoError(t, err, "GetUserByEmail() should not return an error")
	assert.NotNil(t, user, "GetUserByEmail() should return a user")
	assert.Equal(t, int64(1), user.ID, "GetUserByEmail() should return correct user ID")
	assert.Equal(t, "test@example.com", user.Email, "GetUserByEmail() should return correct email")
	assert.Equal(t, "hashedpassword", user.Password, "GetUserByEmail() should return correct password")
	assert.Equal(t, "gravatar", user.Gravatar, "GetUserByEmail() should return correct gravatar")
}

func TestUsersRepository_GetPublicUserByID_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userID := int64(1)

	// Simulate error in query execution
	mock.ExpectQuery(`SELECT`).
		WithArgs(userID).
		WillReturnError(errors.New("query execution error"))

	publicUser, err := repo.GetPublicUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Nil(t, publicUser)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	user, err := nilRepo.GetPublicUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")

	// Test case: User not found (sql.ErrNoRows)
	mock.ExpectQuery(`SELECT .* FROM users u LEFT JOIN files f ON u.wallpaper_id = f.id WHERE u.deleted IS NULL AND u.id = \$1 AND u.deleted IS NULL`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	user, err = repo.GetPublicUserByID(ctx, 1)
	assert.Error(t, err, "GetPublicUserByID() should return an error when user not found")
	assert.Nil(t, user, "GetPublicUserByID() should return nil user when not found")
	assert.Equal(t, "[repository]: user not found", err.Error(), "GetPublicUserByID() should return correct error message")

	// Test case: Unexpected database error (query failure)
	mock.ExpectQuery(`SELECT .* FROM users u LEFT JOIN files f ON u.wallpaper_id = f.id WHERE u.deleted IS NULL AND u.id = \$1 AND u.deleted IS NULL`).
		WithArgs(1).
		WillReturnError(errors.New("database query failed"))

	user, err = repo.GetPublicUserByID(ctx, 1)
	assert.Error(t, err, "GetPublicUserByID() should return an error on unexpected database failure")
	assert.Nil(t, user, "GetPublicUserByID() should return nil user on query failure")
	assert.Equal(t, "database query failed", err.Error(), "GetPublicUserByID() should return correct error message")

	// Test case: Database query does not scan results properly
	mock.ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "uuid", "username", "email", "gravatar", "wallpaper"}).
			AddRow(1, "uuid-123", "user1", "user@example.com", "gravatar_url", "wallpaper_url").
			RowError(0, errors.New("scan failed")))

	user, err = repo.GetPublicUserByID(ctx, 1)
	assert.Error(t, err, "GetPublicUserByID() should return an error on scan failure")
	assert.Nil(t, user, "GetPublicUserByID() should return nil user on scan failure")
	assert.Equal(t, "scan failed", err.Error(), "GetPublicUserByID() should return correct error message")

	// Define the expected row that the mock database will return
	mock.ExpectQuery(`SELECT .* FROM users u LEFT JOIN files f ON u.wallpaper_id = f.id WHERE u.deleted IS NULL AND u.id = \$1 AND u.deleted IS NULL`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "uuid", "username", "email", "gravatar", "wallpaper"}).
			AddRow(1, "uuid-123", "user1", "user@example.com", "gravatar_url", "wallpaper_url"))

	// Call the function
	user, err = repo.GetPublicUserByID(ctx, 1)

	// Assert no error and valid user data
	assert.NoError(t, err, "GetPublicUserByID() should not return an error")
	assert.NotNil(t, user, "GetPublicUserByID() should return a user")
	assert.Equal(t, int64(1), user.ID, "GetPublicUserByID() should return correct user ID")
	assert.Equal(t, "user1", user.Username, "GetPublicUserByID() should return correct username")
	assert.Equal(t, "user@example.com", user.Email, "GetPublicUserByID() should return correct email")

	// Ensure the mock expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}

func TestUsersRepository_UpdateNonce_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userMap := map[string]interface{}{
		"id":    1,
		"nonce": "newnonce",
	}

	// Simulate error in query execution
	mock.ExpectExec(`UPDATE users`).
		WithArgs("newnonce", 1).
		WillReturnError(errors.New("query execution error"))

	err := repo.UpdateNonce(ctx, userMap)
	assert.Error(t, err)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.UpdateNonce(ctx, userMap)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestUsersRepository_UpdateGravatar_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userMap := map[string]interface{}{
		"id":       1,
		"gravatar": "newgravatar",
	}

	// Simulate error in query execution
	mock.ExpectExec(`UPDATE users`).
		WithArgs("newgravatar", 1).
		WillReturnError(errors.New("query execution error"))

	err := repo.UpdateGravatar(ctx, userMap)
	assert.Error(t, err)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.UpdateGravatar(ctx, userMap)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestUsersRepository_UpdateWallpaper_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userMap := map[string]interface{}{
		"id":           1,
		"wallpaper_id": 10,
	}

	// Simulate error in query execution
	mock.ExpectExec(`UPDATE users`).
		WithArgs(10, 1).
		WillReturnError(errors.New("query execution error"))

	err := repo.UpdateWallpaper(ctx, userMap)
	assert.Error(t, err)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.UpdateWallpaper(ctx, userMap)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestUsersRepository_UpdateSecret_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}

	ctx := context.Background()
	userMap := map[string]interface{}{
		"id":     1,
		"secret": "newsecret",
	}

	// Simulate error in query execution
	mock.ExpectExec(`UPDATE users`).
		WithArgs("newsecret", 1).
		WillReturnError(errors.New("query execution error"))

	err := repo.UpdateSecret(ctx, userMap)
	assert.Error(t, err)
	assert.EqualError(t, err, "query execution error")

	// Test case: repo is nil
	var nilRepo *repository.UsersRepository
	err = nilRepo.UpdateSecret(ctx, userMap)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestUsersRepository_GetUserByAddress(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.UsersRepository{
		DB:  *sqlxDB,
		TBL: "users",
	}
	ctx := context.Background()

	// Test Case 1: Nil repository
	t.Run("Nil repository", func(t *testing.T) {
		nilRepo := (*repository.UsersRepository)(nil)
		user, err := nilRepo.GetUserByAddress(ctx, "address1")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, repository.ErrDBNotInitialized, err)
	})

	// Test Case 2: User found by address
	t.Run("User found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT`).
			WithArgs("address1").
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "uuid", "username", "email", "password", "{address}", "nonce", "gravatar", "wallpaper"}).
				AddRow(1, "uuid-123", "user1", "user@example.com", "password", "{address1}", "nonce1", "gravatar_url", "wallpaper_url"))

		user, err := repo.GetUserByAddress(ctx, "address1")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, sql.NullString(sql.NullString{String: "uuid-123", Valid: true}), user.UUID)
		assert.Equal(t, "user1", user.Username)
		assert.Equal(t, "user@example.com", user.Email)
		assert.Equal(t, pq.StringArray(pq.StringArray{"address1"}), user.Address)
		assert.Equal(t, sql.NullString(sql.NullString{String: "nonce1", Valid: true}), user.Nonce)

		err = mock.ExpectationsWereMet()
		if err != nil {
			t.Errorf("there were unmet expectations: %s", err)
		}
	})

	// Test Case 3: User not found (sql.ErrNoRows)
	t.Run("User not found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT`).
			WithArgs("address1").
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserByAddress(ctx, "address1")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})

	// Test Case 4: SQL error (unexpected database error)
	t.Run("SQL error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT`).
			WithArgs("address1").
			WillReturnError(errors.New("some database error"))

		user, err := repo.GetUserByAddress(ctx, "address1")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "some database error", err.Error())
	})
}
