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

func TestSetupStorage_ConnectAuthorsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(errors.New("mock authors repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock authors repository error", "SetupStorage should return the error from ConnectRepository(`authors`)")
}

func TestSetupStorage_ConnectContactsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(errors.New("mock contacts repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock contacts repository error", "SetupStorage should return the error from ConnectRepository(`contacts`)")
}

func TestSetupStorage_ConnectContactFormsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(errors.New("mock contactForms repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock contactForms repository error", "SetupStorage should return the error from ConnectRepository(`contactForms`)")
}

func TestSetupStorage_ConnectFaqsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("faqs", nil).Return(errors.New("mock faqs repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock faqs repository error", "SetupStorage should return the error from ConnectRepository(`faqs`)")
}

func TestSetupStorage_ConnectHelpCenterRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("faqs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("helpCenter", nil).Return(errors.New("mock helpCenter repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock helpCenter repository error", "SetupStorage should return the error from ConnectRepository(`helpCenter`)")
}

func TestSetupStorage_ConnectBlogsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("faqs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("helpCenter", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("blogs", nil).Return(errors.New("mock blogs repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock blogs repository error", "SetupStorage should return the error from ConnectRepository(`blogs`)")
}

func TestSetupStorage_ConnectTagsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("faqs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("helpCenter", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("blogs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("tags", nil).Return(errors.New("mock tags repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock tags repository error", "SetupStorage should return the error from ConnectRepository(`tags`)")
}

func TestSetupStorage_ConnectWalletsRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("faqs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("helpCenter", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("blogs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("tags", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("wallets", nil).Return(errors.New("mock wallets repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock wallets repository error", "SetupStorage should return the error from ConnectRepository(`wallets`)")
}

func TestSetupStorage_ConnectCreateAndSellRepositoryFailure(t *testing.T) {
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
	mockStorage.EXPECT().ConnectRepository("files", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("authors", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contacts", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("contactForms", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("faqs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("helpCenter", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("blogs", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("tags", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("wallets", nil).Return(nil)
	mockStorage.EXPECT().ConnectRepository("create-and-sell", nil).Return(errors.New("mock create-and-sell repository error"))
	// Act
	store, err := storage.SetupStorage(ctx, dsn)

	// Assert
	assert.Nil(t, store, "SetupStorage should return nil storage on error")
	assert.EqualError(t, err, "mock create-and-sell repository error", "SetupStorage should return the error from ConnectRepository(`create-and-sell`)")
}

func TestDBStorage_ConnectRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	for _, name := range []string{
		"users",
		"contracts",
		"pages",
		"categories",
		"menus",
		"sliders",
		"files",
		"authors",
		"contacts",
		"contactForms",
		"faqs",
		"helpCenter",
		"blogs",
		"tags",
		"wallets",
		"create-and-sell",
	} {
		// Test case 1: Nil DBStorage

		nilStorage := mocks.NewMockDBStorageInterface(ctrl)
		nilStorage.EXPECT().ConnectRepository(name, gomock.Any()).Return(errors.New("[storage]: DBStorage is nil"))
		err := nilStorage.ConnectRepository(name, nil)
		assert.Error(t, err, "ConnectRepository(`"+name+"`) should return an error for nil DBStorage")
		assert.Equal(t, "[storage]: DBStorage is nil", err.Error(), "ConnectRepository(`"+name+"`) should return the correct error message for nil DBStorage")

		// Setup mock database and expectations
		store := mocks.NewMockDBStorageInterface(ctrl)
		store.EXPECT().ConnectRepository(name, gomock.Any()).Return(nil)

		// Test case 2: Successful connection with a valid sqlxDB
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

		err = store.ConnectRepository(name, sqlxDB)
		assert.NoError(t, err, "ConnectRepository(`"+name+"`) should not return an error for successful connection")

		// Test case 3: Successful connection with a nil sqlxDB

		err = s.ConnectRepository(name, nil)
		assert.Error(t, err, "ConnectRepository(`"+name+"`) should not return an error when sqlxDB is nil")
		assert.Nil(t, s.GetMenusRepository(), "ConnectRepository(`"+name+"`) should still set the Files repository even if sqlxDB is nil")

		// Ensure all expectations were met
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err, "All expectations should be met")
	}

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

	for _, name := range []string{
		"users",
		"contracts",
		"pages",
		"categories",
		"menus",
		"sliders",
		"files",
		"authors",
		"contacts",
		"contactForms",
		"faqs",
		"helpCenter",
		"blogs",
		"tags",
		"wallets",
		"create-and-sell",
	} {
		mockStore.EXPECT().ConnectRepository(name, gomock.Any()).Return(nil)
	}

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

func TestDBStorage_GetUsersRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetUsersRepository()
	assert.Nil(t, repo, "GetUsersRepository() should return nil when users repository is not initialized")

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
	err = s.ConnectRepository("users", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`users`) should not return an error for successful connection")
	assert.NotNil(t, s.GetUsersRepository(), "ConnectRepository(`contracts`) should still set the users repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
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
	assert.Nil(t, err, "ConnectRepository(`contracts`) should not return an error for successful connection")
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

func TestDBStorage_GetCategoriesRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetCategoriesRepository()
	assert.Nil(t, repo, "GetCategoriesRepository() should return nil when categories repository is not initialized")

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
	err = s.ConnectRepository("categories", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`categories`) should not return an error for successful connection")

	assert.NotNil(t, s.GetCategoriesRepository(), "ConnectRepository(`categories`) should still set the categories repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetAuthorsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetAuthorsRepository()
	assert.Nil(t, repo, "GetAuthorsRepository() should return nil when authors repository is not initialized")

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
	err = s.ConnectRepository("authors", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`authors`) should not return an error for successful connection")

	assert.NotNil(t, s.GetAuthorsRepository(), "ConnectRepository(`authors`) should still set the authors repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetContactFormsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetContactFormsRepository()
	assert.Nil(t, repo, "GetContactFormsRepository() should return nil when contactForms repository is not initialized")

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
	err = s.ConnectRepository("contactForms", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`contactForms`) should not return an error for successful connection")

	assert.NotNil(t, s.GetContactFormsRepository(), "ConnectRepository(`contactForms`) should still set the contactForms repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetFaqsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetFaqsRepository()
	assert.Nil(t, repo, "GetFaqsRepository() should return nil when faqs repository is not initialized")

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
	err = s.ConnectRepository("faqs", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`faqs`) should not return an error for successful connection")

	assert.NotNil(t, s.GetFaqsRepository(), "ConnectRepository(`faqs`) should still set the faqs repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetHelpCenterRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetHelpCenterRepository()
	assert.Nil(t, repo, "GetHelpCenterRepository() should return nil when helpCenter repository is not initialized")

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
	err = s.ConnectRepository("helpCenter", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`helpCenter`) should not return an error for successful connection")

	assert.NotNil(t, s.GetHelpCenterRepository(), "ConnectRepository(`helpCenter`) should still set the helpCenter repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetBlogsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetBlogsRepository()
	assert.Nil(t, repo, "GetBlogsRepository() should return nil when blogs repository is not initialized")

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
	err = s.ConnectRepository("blogs", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`blogs`) should not return an error for successful connection")

	assert.NotNil(t, s.GetBlogsRepository(), "ConnectRepository(`blogs`) should still set the blogs repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetTagsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetTagsRepository()
	assert.Nil(t, repo, "GetTagsRepository() should return nil when tags repository is not initialized")

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
	err = s.ConnectRepository("tags", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`tags`) should not return an error for successful connection")

	assert.NotNil(t, s.GetTagsRepository(), "ConnectRepository(`tags`) should still set the tags repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetContactsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetContactsRepository()
	assert.Nil(t, repo, "GetContactsRepository() should return nil when contacts repository is not initialized")

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
	err = s.ConnectRepository("contacts", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`contacts`) should not return an error for successful connection")

	assert.NotNil(t, s.GetContactsRepository(), "ConnectRepository(`contacts`) should still set the contacts repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetWalletsRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetWalletsRepository()
	assert.Nil(t, repo, "GetWalletsRepository() should return nil when wallets repository is not initialized")

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
	err = s.ConnectRepository("wallets", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`wallets`) should not return an error for successful connection")

	assert.NotNil(t, s.GetWalletsRepository(), "ConnectRepository(`wallets`) should still set the wallets repository even if sqlxDB is nil")

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestDBStorage_GetCreateAndSellRepository(t *testing.T) {
	// Test case 1: Nil DBStorage
	s := &storage.DBStorage{
		DSN: "user=postgres password=secret dbname=testdb sslmode=disable",
	}

	repo := s.GetCreateAndSellRepository()
	assert.Nil(t, repo, "GetCreateAndSellRepository() should return nil when create-and-sell repository is not initialized")

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
	err = s.ConnectRepository("create-and-sell", sqlxDB)
	assert.Nil(t, err, "ConnectRepository(`create-and-sell`) should not return an error for successful connection")

	assert.NotNil(t, s.GetCreateAndSellRepository(), "ConnectRepository(`create-and-sell`) should still set the create-and-sell repository even if sqlxDB is nil")

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
