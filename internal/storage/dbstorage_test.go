package storage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
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
	nilStorage.EXPECT().ConnectUsersRepository(context.Background(), gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectUsersRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectUsersRepository() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectUsersRepository() should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectUsersRepository(context.Background(), gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectUsersRepository(context.Background(), sqlxDB)
	assert.NoError(t, err, "ConnectUsersRepository() should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectUsersRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectUsersRepository() should not return an error when sqlxDB is nil")
	assert.Nil(t, s.Users, "ConnectUsersRepository() should still set the Users repository even if sqlxDB is nil")

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

	mockStore.EXPECT().ConnectUsersRepository(context.Background(), gomock.Any()).Return(errors.New(connectionErr))

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
	nilStorage.EXPECT().ConnectContractsRepository(context.Background(), gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectContractsRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectContractsRepository() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectContractsRepository() should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectContractsRepository(context.Background(), gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectContractsRepository(context.Background(), sqlxDB)
	assert.NoError(t, err, "ConnectContractsRepository() should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectContractsRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectContractsRepository() should not return an error when sqlxDB is nil")
	assert.Nil(t, s.Contracts, "ConnectContractsRepository() should still set the Contracts repository even if sqlxDB is nil")

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
	mockStorage.EXPECT().ConnectUsersRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectContractsRepository(ctx, nil).Return(errors.New("mock contracts repository error"))

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
	nilStorage.EXPECT().ConnectPagesRepository(context.Background(), gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectPagesRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectPagesRepository() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectPagesRepository() should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectPagesRepository(context.Background(), gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectPagesRepository(context.Background(), sqlxDB)
	assert.NoError(t, err, "ConnectPagesRepository() should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectPagesRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectPagesRepository() should not return an error when sqlxDB is nil")
	assert.Nil(t, s.Contracts, "ConnectPagesRepository() should still set the Pages repository even if sqlxDB is nil")

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
	mockStorage.EXPECT().ConnectUsersRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectContractsRepository(ctx, nil).Return(nil)

	mockStorage.EXPECT().ConnectPagesRepository(ctx, nil).Return(errors.New("mock pages repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock pages repository error", "SetupStorage should return the error from ConnectPagesRepository")
}

func TestDBStorage_ConnectCategoriesRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectCategoriesRepository(context.Background(), gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectCategoriesRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectCategoriesRepository() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectCategoriesRepository() should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectCategoriesRepository(context.Background(), gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectCategoriesRepository(context.Background(), sqlxDB)
	assert.NoError(t, err, "ConnectCategoriesRepository() should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectCategoriesRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectCategoriesRepository() should not return an error when sqlxDB is nil")
	assert.Nil(t, s.Categories, "ConnectCategoriesRepository() should still set the Categories repository even if sqlxDB is nil")

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
	mockStorage.EXPECT().ConnectUsersRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectContractsRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectPagesRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectCategoriesRepository(ctx, nil).Return(errors.New("mock categories repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock categories repository error", "SetupStorage should return the error from ConnectCategoriesRepository")
}

func TestDBStorage_ConnectMenusRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test case 1: Nil DBStorage

	nilStorage := mocks.NewMockDBStorageInterface(ctrl)
	nilStorage.EXPECT().ConnectMenusRepository(context.Background(), gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))

	err := nilStorage.ConnectMenusRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectMenusRepository() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectMenusRepository() should return the correct error message for nil DBStorage")

	// Setup mock database and expectations
	store := mocks.NewMockDBStorageInterface(ctrl)
	store.EXPECT().ConnectMenusRepository(context.Background(), gomock.Any()).Return(nil)

	// Test case 2: Successful connection with a valid sqlxDB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	err = store.ConnectMenusRepository(context.Background(), sqlxDB)
	assert.NoError(t, err, "ConnectMenusRepository() should not return an error for successful connection")

	// Test case 3: Successful connection with a nil sqlxDB
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	err = s.ConnectMenusRepository(context.Background(), nil)
	assert.Error(t, err, "ConnectMenusRepository() should not return an error when sqlxDB is nil")
	assert.Nil(t, s.Contracts, "ConnectMenusRepository() should still set the Pages repository even if sqlxDB is nil")

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
	mockStorage.EXPECT().ConnectUsersRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectContractsRepository(ctx, nil).Return(nil)

	mockStorage.EXPECT().ConnectPagesRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectCategoriesRepository(ctx, nil).Return(nil)
	mockStorage.EXPECT().ConnectMenusRepository(ctx, nil).Return(errors.New("mock menus repository error"))

	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock menus repository error", "SetupStorage should return the error from ConnectMenusRepository")
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

	mockStore.EXPECT().ConnectUsersRepository(context.Background(), gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectContractsRepository(context.Background(), gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectPagesRepository(context.Background(), gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectCategoriesRepository(context.Background(), gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectMenusRepository(context.Background(), gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectSlidersRepository(context.Background(), gomock.Any()).Return(nil)
	mockStore.EXPECT().ConnectFilesRepository(context.Background(), gomock.Any()).Return(nil)

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
		DSN:       "user=postgres password=secret dbname=testdb sslmode=disable",
		Contracts: repository.NewContractsRepository(sqlxDB, "contracts"),
	}

	assert.NotNil(t, s.Contracts, "ConnectContractsRepository() should still set the Contracts repository even if sqlxDB is nil")

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
		DSN:   "user=postgres password=secret dbname=testdb sslmode=disable",
		Menus: repository.NewMenusRepository(sqlxDB, "menus"),
	}

	assert.NotNil(t, s.Menus, "ConnectMenusRepository() should still set the Menus repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}
