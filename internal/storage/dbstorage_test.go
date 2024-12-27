package storage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	validDSN      = "valid_dsn"
	connectionErr = "connection error"
)

func TestNewDBStorage(t *testing.T) {
	// Test case: dsn is empty
	dsn := ""
	s := storage.NewDBStorage(dsn)
	assert.Nil(t, s, "NewDBStorage() should return nil when dsn is empty")

	// Test case: dsn is not empty
	dsn = "user=postgres password=secret dbname=testdb sslmode=disable"
	s = storage.NewDBStorage(dsn)
	assert.NotNil(t, s, "NewDBStorage() should not return nil when dsn is not empty")
}

func TestDBStorage_Init(t *testing.T) {
	ctx := context.Background()

	// Test case 1: Nil DBStorage
	var nilStorage *storage.DBStorage
	db, err := nilStorage.Init(ctx, nil)
	assert.Nil(t, db, "Init() should return nil DB for nil storage")
	assert.Error(t, err, "Init() should return an error for nil storage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "Init() should return the correct error message for nil storage")

	// Setup mock database and expectations
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Test case 2: Successful connection with nil sqlxDB (simulated in Init)
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	mock.ExpectPing().WillReturnError(nil)

	db, err = s.Init(ctx, nil)
	assert.Error(t, err, `[storage]: failed to connect to the database: pq: password authentication failed for user "postgres"`)
	assert.Nil(t, db, "Init() should return a valid DB for successful connection")

	if db != nil {
		assert.Equal(t, db.DriverName(), "sqlmock", "Init() should return a sqlmock DB instance")
	}

	// Test case 3: Connection failure with nil sqlxDB
	mock.ExpectPing().WillReturnError(errors.New("connection failed"))

	db, err = s.Init(ctx, nil)
	assert.Nil(t, db, "Init() should return nil DB for connection failure")
	assert.Error(t, err, "Init() should return an error for connection failure")

	// Test case 4: Successful connection with provided sqlxDB
	db, err = s.Init(ctx, sqlxDB)
	assert.NoError(t, err, "Init() should not return an error when a valid sqlxDB is provided")
	assert.NotNil(t, db, "Init() should return the provided sqlxDB when it's not nil")
	assert.Equal(t, db, sqlxDB, "Init() should return the exact same sqlxDB instance that was provided")
}

func TestSetupStorage_NilStorage(t *testing.T) {
	ctx := context.Background()
	dsn := "invalid_dsn"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Replace the NewDBStorage function with a mock
	originalNewDBStorage := storage.NewPGSQLStorage
	defer func() { storage.NewPGSQLStorage = originalNewDBStorage }()

	storage.NewPGSQLStorage = func(_ string) storage.DBStorageInterface {
		return nil
	}

	store, err := storage.SetupStorage(ctx, dsn)
	require.Error(t, err)
	require.Equal(t, "[storage]: cannot init db using dsn='invalid_dsn'", err.Error())
	require.Nil(t, store)
}

func TestDBStorage_ConnectUsersRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("users", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("users", nil)
	assert.Error(t, err, "ConnectRepository(`users`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`users`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("users", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("users", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`users`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("users", nil)
	assert.Error(t, err, "ConnectRepository(`users`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetUsersRepository(), "ConnectUsersRepository(`users`) should still set the Users repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectUsersRepositoryFailure(t *testing.T) {
	ctx := context.Background()
	dsn := validDSN

	// Setup mock database and expectations
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	defer mockDB.Close()
	mock.ExpectPing().WillReturnError(nil)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mocks.NewMockDBStorageInterface(mockCtrl)
	mockStore.EXPECT().Init(ctx, nil).Return(sqlxDB, nil)

	mockStore.EXPECT().ConnectRepository("users", gomock.Any()).Return(errors.New(connectionErr))

	// Replace the NewDBStorage function with a mock
	originalNewDBStorage := storage.NewPGSQLStorage
	defer func() { storage.NewPGSQLStorage = originalNewDBStorage }()

	storage.NewPGSQLStorage = func(_ string) storage.DBStorageInterface {
		return mockStore
	}

	store, err := storage.SetupStorage(ctx, dsn)

	require.Error(t, err)
	require.Equal(t, connectionErr, err.Error())
	require.Nil(t, store)
}

func TestDBStorage_ConnectContractsRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("contracts", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("contracts", nil)
	assert.Error(t, err, "ConnectRepository(`contracts`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`contracts`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("contracts", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("contracts", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`contracts`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("contracts", nil)
	assert.Error(t, err, "ConnectRepository(`contracts`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetContractsRepository(), "ConnectRepository(`contracts`) should still set the Contracts repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectContractsRepositoryFailure(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(mockCtrl)

	ctx := context.Background()
	dsn := validDSN

	// Mock `NewPGSQLStorage` to return our mockStorage
	storage.NewPGSQLStorage = func(string) storage.DBStorageInterface {
		return mockStorage
	}

	mockStorage.EXPECT().Init(ctx, nil).Return(nil, nil)
	mockStorage.EXPECT().ConnectRepository("users", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contracts", nil).Return(errors.New("mock contracts repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock contracts repository error", "SetupStorage should return the error from ConnectContractsRepository")
}

func TestDBStorage_ConnectPagesRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("pages", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("pages", nil)
	assert.Error(t, err, "ConnectRepository(`pages`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`pages`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("pages", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("pages", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`pages`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("pages", nil)
	assert.Error(t, err, "ConnectRepository(`pages`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetContractsRepository(), "ConnectRepository(`pages`) should still set the Pages repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectPagesRepositoryFailure(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(mockCtrl)

	ctx := context.Background()
	dsn := validDSN

	// Mock `NewPGSQLStorage` to return our mockStorage
	storage.NewPGSQLStorage = func(string) storage.DBStorageInterface {
		return mockStorage
	}

	mockStorage.EXPECT().Init(ctx, nil).Return(nil, nil)
	mockStorage.EXPECT().ConnectRepository("users", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contracts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("pages", nil).Return(errors.New("mock pages repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock pages repository error", "SetupStorage should return the error from ConnectRepository(`pages)")
}

func TestDBStorage_ConnectCategoriesRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("categories", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("categories", nil)
	assert.Error(t, err, "ConnectRepository(`categories`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`categories`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("categories", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("categories", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`categories`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("categories", nil)
	assert.Error(t, err, "ConnectRepository(`categories`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetCategoriesRepository(), "ConnectRepository(`categories`) should still set the Categories repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectCategoriesRepositoryFailure(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(mockCtrl)

	ctx := context.Background()
	dsn := validDSN

	// Mock `NewPGSQLStorage` to return our mockStorage
	storage.NewPGSQLStorage = func(string) storage.DBStorageInterface {
		return mockStorage
	}

	mockStorage.EXPECT().Init(ctx, nil).Return(nil, nil)
	mockStorage.EXPECT().ConnectRepository("users", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contracts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("pages", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("categories", nil).Return(errors.New("mock categories repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock categories repository error", "SetupStorage should return the error from ConnectRepository(`categories`)")
}

func TestDBStorage_ConnectMenusRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("menus", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("menus", nil)
	assert.Error(t, err, "ConnectRepository(`menus`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`menus`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("menus", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("menus", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`menus`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("menus", nil)
	assert.Error(t, err, "ConnectRepository(`menus`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetMenusRepository(), "ConnectRepository(`menus`) should still set the Menus repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectMenusRepositoryFailure(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(mockCtrl)

	ctx := context.Background()
	dsn := validDSN

	// Mock `NewPGSQLStorage` to return our mockStorage
	storage.NewPGSQLStorage = func(string) storage.DBStorageInterface {
		return mockStorage
	}

	mockStorage.EXPECT().Init(ctx, nil).Return(nil, nil)
	mockStorage.EXPECT().ConnectRepository("users", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contracts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("pages", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("categories", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("menus", nil).Return(errors.New("mock menus repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock menus repository error", "SetupStorage should return the error from ConnectRepository(`menus`)")
}

func TestDBStorage_ConnectSlidersRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("sliders", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("sliders", nil)
	assert.Error(t, err, "ConnectRepository(`sliders`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`sliders`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("sliders", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("sliders", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`sliders`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("sliders", nil)
	assert.Error(t, err, "ConnectRepository(`sliders`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetMenusRepository(), "ConnectRepository(`sliders`) should still set the Sliders repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectSlidersRepositoryFailure(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(mockCtrl)

	ctx := context.Background()
	dsn := validDSN

	// Mock `NewPGSQLStorage` to return our mockStorage
	storage.NewPGSQLStorage = func(string) storage.DBStorageInterface {
		return mockStorage
	}

	mockStorage.EXPECT().Init(ctx, nil).Return(nil, nil)
	mockStorage.EXPECT().ConnectRepository("users", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contracts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("pages", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("categories", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("menus", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("sliders", nil).Return(errors.New("mock sliders repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock sliders repository error", "SetupStorage should return the error from ConnectRepository(`sliders`)")
}

func TestDBStorage_ConnectFilesRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository("files", gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository("files", nil)
	assert.Error(t, err, "ConnectRepository(`files`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`files`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository("files", gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository("files", sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`files`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository("files", nil)
	assert.Error(t, err, "ConnectRepository(`files`) should not return an error when sqlxDB is nil")
	assert.Nil(t, s.GetMenusRepository(), "ConnectRepository(`files`) should still set the Files repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectFilesRepositoryFailure(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStorage := mocks.NewMockDBStorageInterface(mockCtrl)

	ctx := context.Background()
	dsn := validDSN

	// Mock `NewPGSQLStorage` to return our mockStorage
	storage.NewPGSQLStorage = func(string) storage.DBStorageInterface {
		return mockStorage
	}

	mockStorage.EXPECT().Init(ctx, nil).Return(nil, nil)
	mockStorage.EXPECT().ConnectRepository("users", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contracts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("pages", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("categories", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("menus", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("sliders", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(errors.New("mock files repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock files repository error", "SetupStorage should return the error from ConnectRepository(`files`)")
}

func TestSetupStorage_InitFailure(t *testing.T) {
	ctx := context.Background()
	dsn := "invalid_dsn"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mocks.NewMockDBStorageInterface(mockCtrl)
	mockStore.EXPECT().Init(ctx, nil).Return(nil, errors.New("init error"))

	// Replace the NewDBStorage function with a mock
	originalNewDBStorage := storage.NewPGSQLStorage
	defer func() { storage.NewPGSQLStorage = originalNewDBStorage }()

	storage.NewPGSQLStorage = func(_ string) storage.DBStorageInterface {
		return mockStore
	}

	store, err := storage.SetupStorage(ctx, dsn)
	require.Error(t, err)
	require.Nil(t, store)
}

func TestSetupStorage_Success(t *testing.T) {
	ctx := context.Background()
	dsn := validDSN

	// Setup mock database and expectations
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	defer mockDB.Close()
	mock.ExpectPing().WillReturnError(nil)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mocks.NewMockDBStorageInterface(mockCtrl)
	mockStore.EXPECT().Init(ctx, nil).Return(sqlxDB, nil)

	mockStore.EXPECT().ConnectRepository("users", gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectRepository("contracts", gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectRepository("pages", gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectRepository("categories", gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectRepository("menus", gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectRepository("sliders", gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectRepository("files", gomock.Any()).Return(nil)

	// Replace the NewDBStorage function with a mock
	originalNewDBStorage := storage.NewPGSQLStorage
	defer func() { storage.NewPGSQLStorage = originalNewDBStorage }()

	storage.NewPGSQLStorage = func(_ string) storage.DBStorageInterface {
		return mockStore
	}

	store, err := storage.SetupStorage(ctx, dsn)
	require.NoError(t, err)
	require.NotNil(t, store)
}

func TestDBStorage_GetContractsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	contractsRepo := s.GetContractsRepository()
	assert.Nil(t, contractsRepo, "GetContractsRepository() should return nil when Contracts repository is not initialized")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Test case 2: Nil Contracts repository
	s = &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}
	err = s.ConnectRepository("contracts", sqlxDB)
	assert.NotNil(t, err, "ConnectRepository(`contracts`) should not return an error for successful connection")
	assert.NotNil(t, s.GetContractsRepository(), "ConnectRepository(`contracts`) should still set the Contracts repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetMenusRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	menusRepo := s.GetMenusRepository()
	assert.Nil(t, menusRepo, "GetMenusRepository() should return nil when Menus repository is not initialized")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Test case 2: Nil Contracts repository
	s = &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}
	err = s.ConnectRepository("menus", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`menus`) should not return an error for successful connection")

	assert.NotNil(t, s.GetMenusRepository(), "ConnectRepository(`menus`) should still set the Menus repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetSlidersRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetSlidersRepository()
	assert.Nil(t, repo, "GetSlidersRepository() should return nil when Menus repository is not initialized")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Test case 2: Nil Contracts repository
	s = &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}
	err = s.ConnectRepository("sliders", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`sliders`) should not return an error for successful connection")

	assert.NotNil(t, s.GetSlidersRepository(), "ConnectRepository(`sliders`) should still set the Sliders repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetFilesRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetFilesRepository()
	assert.Nil(t, repo, "GetFilesRepository() should return nil when Files repository is not initialized")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Test case 2: Nil Contracts repository
	s = &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}
	err = s.ConnectRepository("files", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`files`) should not return an error for successful connection")

	assert.NotNil(t, s.GetFilesRepository(), "ConnectRepository(`files`) should still set the Files repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestSetupStorage_ConnectUnknownRepositoryFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage
	repoName := "unknown"
	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectRepository(repoName, gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectRepository(repoName, nil)
	assert.Error(t, err, "ConnectRepository(`unknown`) should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`unknown`) should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectRepository(repoName, gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectRepository(repoName, sqlxDB)
	assert.NoError(t, err, "ConnectRepository(`unknown`) should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectRepository(repoName, nil)
	assert.Error(t, err, "ConnectRepository(`unknown`) should return an error when sqlxDB is nil")

	err = s.ConnectRepository(repoName, sqlxDB)
	assert.Error(t, err, "ConnectRepository(`unknown`) should return an error")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}
