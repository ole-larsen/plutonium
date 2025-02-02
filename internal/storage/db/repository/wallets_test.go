package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletsRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	tbl := "users"
	repo := repository.NewWalletsRepository(sqlxDB, tbl)

	// Check if db is correctly set
	assert.Equal(t, *sqlxDB, repo.DB, "NewWalletsRepository() db mismatch")
	// Check if table name is correctly set
	assert.Equal(t, tbl, repo.TBL, "NewWalletsRepository() tbl mismatch")

	repo = repository.NewWalletsRepository(nil, tbl)
	assert.Nil(t, repo, "NewWalletsRepository() should return nil when db is nil")
}

func TestPing(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.WalletsRepository{
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
	var nilRepo *repository.WalletsRepository
	err = nilRepo.Ping()
	assert.Error(t, err, "Ping() on nil repository should return an error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Ping() should return ErrDBNotInitialized")
}

func TestWalletsRepository_InnerDB(t *testing.T) {
	db, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.WalletsRepository{
		DB: *sqlxDB,
	}

	// Test case: repo is not nil
	assert.Equal(t, sqlxDB, repo.InnerDB(), "InnerDB() should return the correct *sqlx.DB")

	// Test case: repo is nil
	var nilRepo *repository.WalletsRepository

	assert.Nil(t, nilRepo.InnerDB(), "InnerDB() on nil repository should return nil")
}

func TestCreateWallet(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database execution
	mock.ExpectExec("INSERT INTO wallets").
		WithArgs("Test Wallet", "Description", 1, true, 0, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test Create method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	err = repo.Create(context.Background(), map[string]interface{}{
		"title":         "Test Wallet",
		"description":   "Description",
		"image_id":      1,
		"enabled":       true,
		"order_by":      0,
		"created_by_id": 1,
		"updated_by_id": 1,
	})
	assert.NoError(t, err)

	// Test case: repo is nil
	ctx := context.Background()

	var nilRepo *repository.WalletsRepository
	err = nilRepo.Create(ctx, map[string]interface{}{
		"title":         "Test Wallet",
		"description":   "Description",
		"image_id":      1,
		"enabled":       true,
		"order_by":      0,
		"created_by_id": 1,
		"updated_by_id": 1,
	})
	assert.Error(t, err, "Create() on nil repository should return error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestCreate_SQLQueryError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock an error during the insert operation
	mock.ExpectExec("INSERT INTO wallets").
		WillReturnError(fmt.Errorf("SQL insert error"))

	// Test Create method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	err = repo.Create(context.Background(), map[string]interface{}{
		"title":         "New Wallet",
		"description":   "Description of the new wallet",
		"image_id":      1,
		"enabled":       true,
		"order_by":      1,
		"created_by_id": 1,
		"updated_by_id": 1,
	})

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, "SQL insert error")
}

func TestUpdate_SQLExecError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock an error during the execution of the update query
	mock.ExpectExec("UPDATE wallets").
		WillReturnError(fmt.Errorf("SQL update execution error"))

	// Test Update method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	_, err = repo.Update(context.Background(), map[string]interface{}{
		"id":            1,
		"title":         "Updated Title",
		"description":   "Updated Description",
		"image_id":      2,
		"enabled":       true,
		"order_by":      1,
		"updated_by_id": 1,
	})

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, "SQL update execution error")
}

func TestUpdateWallet(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database execution for the UPDATE statement
	mock.ExpectExec("UPDATE wallets").
		WithArgs("Updated Title", "Updated Description", 2, true, 1, 1, int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock the GetWallets method to return a sample wallet list
	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "image_id", "enabled", "order_by", "created_by_id", "updated_by_id"}).
			AddRow(1, "Updated Title", "Updated Description", 2, true, 1, 1, 1))

	// Test the Update method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallets, err := repo.Update(context.Background(), map[string]interface{}{
		"id":            int64(1),
		"title":         "Updated Title",
		"description":   "Updated Description",
		"image_id":      2,
		"enabled":       true,
		"order_by":      1,
		"updated_by_id": 1,
	})
	assert.NoError(t, err)
	assert.Len(t, wallets, 1)
	assert.Equal(t, "Updated Title", wallets[0].Title)

	// Test case: repo is nil
	ctx := context.Background()

	var nilRepo *repository.WalletsRepository
	res, err := nilRepo.Update(ctx, map[string]interface{}{
		"id":            int64(1),
		"title":         "Updated Title",
		"description":   "Updated Description",
		"image_id":      2,
		"enabled":       true,
		"order_by":      1,
		"updated_by_id": 1,
	})
	assert.Error(t, err, "Update() on nil repository should return error")
	assert.Nil(t, res)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Create() should return ErrDBNotInitialized")
}

func TestGetWallets(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database rows
	rows := sqlmock.NewRows([]string{"id", "title", "description", "image_id", "enabled", "order_by", "created_by_id", "updated_by_id"}).
		AddRow(1, "Wallet 1", "Description 1", 1, true, 1, 1, 1).
		AddRow(2, "Wallet 2", "Description 2", 2, false, 2, 2, 2)

	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnRows(rows)

	// Test GetWallets method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallets, err := repo.GetWallets(context.Background())
	assert.NoError(t, err)
	assert.Len(t, wallets, 2)
	assert.Equal(t, "Wallet 1", wallets[0].Title)

	// Test case: repo is nil
	ctx := context.Background()

	var nilRepo *repository.WalletsRepository
	res, err := nilRepo.GetWallets(ctx)
	assert.Error(t, err, "GetWallets() on nil repository should return error")
	assert.Nil(t, res)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetWallets() should return ErrDBNotInitialized")
}

func TestGetWallets_SQLQueryError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock a query error (e.g., database connection issues or SQL syntax errors)
	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnError(fmt.Errorf("database query error"))

	// Test GetWallets method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	_, err = repo.GetWallets(context.Background())

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, "database query error")
}

func TestGetWallets_ScanError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the query execution with invalid row data for scanning (e.g., mismatched types)
	rows := sqlmock.NewRows([]string{"id", "title", "description", "image_id", "enabled", "order_by", "created_by_id", "updated_by_id"}).
		AddRow("invalid_id", "Wallet 1", "Description 1", 1, true, 1, 1, 1) // Invalid data type for 'id'

	// Mock the database query execution
	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnRows(rows)

	// Test GetWallets method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	_, err = repo.GetWallets(context.Background())

	// Assert that scanning error occurred
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sql: Scan error on column index 0")
}

func TestGetWalletByID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database row
	row := sqlmock.NewRows([]string{"id", "title", "description", "image_id", "enabled", "order_by", "created_by_id", "updated_by_id"}).
		AddRow(1, "Wallet 1", "Description 1", 1, true, 1, 1, 1)

	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WithArgs(1).
		WillReturnRows(row)

	// Test GetWalletByID method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallet, err := repo.GetWalletByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Wallet 1", wallet.Title)

	// Test case: repo is nil
	ctx := context.Background()

	var nilRepo *repository.WalletsRepository
	res, err := nilRepo.GetWalletByID(ctx, 1)
	assert.Error(t, err, "GetWalletByID() on nil repository should return error")
	assert.Nil(t, res)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetWalletByID() should return ErrDBNotInitialized")
}

func TestGetWalletByID_DBNotInitialized(t *testing.T) {
	// Test for uninitialized repository (nil repository)
	var repo *repository.WalletsRepository
	_, err := repo.GetWalletByID(context.Background(), 1)

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, repository.ErrDBNotInitialized.Error())
}

func TestGetWalletByID_SQLQueryError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock a query error (e.g., database connection issues or SQL syntax errors)
	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnError(fmt.Errorf("database query error"))

	// Test GetWalletByID method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	_, err = repo.GetWalletByID(context.Background(), 1)

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, "[repository]: database query error")
}

func TestGetWalletByID_NoRowsFound(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the case where no rows are returned (sql.ErrNoRows)
	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnError(sql.ErrNoRows)

	// Test GetWalletByID method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	_, err = repo.GetWalletByID(context.Background(), 1)

	// Assert that the wallet not found error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, "wallet not found")
}

func TestGetWalletByID_ScanError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the query execution with invalid row data for scanning (e.g., mismatched types)
	rows := sqlmock.NewRows([]string{"id", "title", "description", "image_id", "enabled", "order_by", "created_by_id", "updated_by_id"}).
		AddRow("invalid_id", "Wallet Title", "Wallet Description", 1, true, 1, 1, 1) // Invalid data type for 'id'

	// Mock the database query execution
	mock.ExpectQuery("SELECT id, title, description, image_id, enabled, order_by, created_by_id, updated_by_id").
		WillReturnRows(rows)

	// Test GetWalletByID method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	_, err = repo.GetWalletByID(context.Background(), 1)

	// Assert that scanning error occurred
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sql: Scan error on column index 0")
}

func TestGetPublicWalletConnect(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database rows with an additional 'image' field.
	// The image will be mocked as a JSON string, simulating the `JSON_BUILD_OBJECT` result.
	rows := sqlmock.NewRows([]string{"title", "description", "image"}).
		AddRow("Wallet 1", "Description 1", `{"id": 1, "attributes": {"name": "image1"}}`).
		AddRow("Wallet 2", "Description 2", `{"id": 2, "attributes": {"name": "image2"}}`)

	// Mock the database query execution
	mock.ExpectQuery("SELECT w.title, w.description,").
		WillReturnRows(rows)

	// Test GetPublicWalletConnect method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallets, err := repo.GetPublicWalletConnect(context.Background())
	assert.NoError(t, err)
	assert.Len(t, wallets, 2)

	// Check the results
	assert.Equal(t, "Wallet 1", wallets[0].Attributes.Title)
	assert.Equal(t, "Description 1", wallets[0].Attributes.Description)
	assert.NotNil(t, wallets[0].Attributes.Image)
	assert.Equal(t, "image1", wallets[0].Attributes.Image.Attributes.Name)

	assert.Equal(t, "Wallet 2", wallets[1].Attributes.Title)
	assert.Equal(t, "Description 2", wallets[1].Attributes.Description)
	assert.NotNil(t, wallets[1].Attributes.Image)
	assert.Equal(t, "image2", wallets[1].Attributes.Image.Attributes.Name)

	// Test case: repo is nil
	ctx := context.Background()

	var nilRepo *repository.WalletsRepository
	res, err := nilRepo.GetPublicWalletConnect(ctx)
	assert.Error(t, err, "GetPublicWalletConnect() on nil repository should return error")
	assert.Nil(t, res)
	assert.Equal(t, repository.ErrDBNotInitialized, err, "GetPublicWalletConnect() should return ErrDBNotInitialized")
}

func TestGetPublicWalletConnect_QueryError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the query execution to return an error
	mock.ExpectQuery("SELECT w.title, w.description").
		WillReturnError(fmt.Errorf("database query error"))

	// Test GetPublicWalletConnect method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallets, err := repo.GetPublicWalletConnect(context.Background())

	// Assert that error occurred and no wallets were returned
	assert.Error(t, err)
	assert.Nil(t, wallets)
	assert.EqualError(t, err, "database query error")
}

func TestGetPublicWalletConnect_ScanError(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database rows, but with incorrect number of columns for Scan
	rows := sqlmock.NewRows([]string{"title", "description"}).AddRow("Wallet 1", "Description 1")

	// Mock the database query execution
	mock.ExpectQuery("SELECT w.title, w.description").
		WillReturnRows(rows)

	// Test GetPublicWalletConnect method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallets, err := repo.GetPublicWalletConnect(context.Background())

	// Assert that error occurred due to Scan mismatch
	assert.Error(t, err)
	assert.Nil(t, wallets)
	assert.EqualError(t, err, "sql: expected 2 destination arguments in Scan, not 3")
}

func TestGetPublicWalletConnect_Success(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database rows with correct structure
	rows := sqlmock.NewRows([]string{"title", "description", "image"}).
		AddRow("Wallet 1", "Description 1", `{"id": 1, "attributes": {"name": "image1"}}`).
		AddRow("Wallet 2", "Description 2", `{"id": 2, "attributes": {"name": "image2"}}`)

	// Mock the database query execution
	mock.ExpectQuery("SELECT w.title, w.description,").
		WillReturnRows(rows)

	// Test GetPublicWalletConnect method
	repo := repository.NewWalletsRepository(&sqlx.DB{DB: db}, "wallets")
	wallets, err := repo.GetPublicWalletConnect(context.Background())

	// Assert that no error occurred and the wallets were returned correctly
	assert.NoError(t, err)
	assert.Len(t, wallets, 2)

	// Check values of the first wallet
	assert.Equal(t, "Wallet 1", wallets[0].Attributes.Title)
	assert.Equal(t, "Description 1", wallets[0].Attributes.Description)
	assert.NotNil(t, wallets[0].Attributes.Image)
	assert.Equal(t, "image1", wallets[0].Attributes.Image.Attributes.Name)

	// Check values of the second wallet
	assert.Equal(t, "Wallet 2", wallets[1].Attributes.Title)
	assert.Equal(t, "Description 2", wallets[1].Attributes.Description)
	assert.NotNil(t, wallets[1].Attributes.Image)
	assert.Equal(t, "image2", wallets[1].Attributes.Image.Attributes.Name)
}
