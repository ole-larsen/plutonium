package storage_test

import (
	"context"
	"errors"
	"fmt"
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
	defaultKVType = "kv"
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

func TestDBStorage_Ping(t *testing.T) {
	// Test case 1: Nil DBStorage
	var nilStorage *storage.DBStorage
	err := nilStorage.Ping()
	assert.Error(t, err, "Ping() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "Ping() should return the correct error message for nil DBStorage")

	// Test case 2: Nil Users repository
	s := &storage.DBStorage{}
	err = s.Ping()
	assert.Error(t, err, "Ping() should return an error when Users repository is nil")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "Ping() should return the correct error message for nil Users repository")

	// Test case 3: Nil InnerDB
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsers := mocks.NewMockUsersRepositoryInterface(mockCtrl)

	s = &storage.DBStorage{
		Users: mockUsers,
	}

	mockUsers.EXPECT().InnerDB().Return(nil).Times(1)

	err = s.Ping()
	assert.Error(t, err, "Ping() should return an error when InnerDB is nil")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "Ping() should return the correct error message for nil InnerDB")

	// Test case 4: Successful Ping
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	mockUsers.EXPECT().InnerDB().Return(sqlxDB).Times(1)
	mockUsers.EXPECT().Ping().Return(nil).Times(1)

	s = &storage.DBStorage{
		Users: mockUsers,
	}

	err = s.Ping()
	assert.NoError(t, err, "Ping() should not return an error for successful ping")

	// Test case 5: Failed Ping
	mockDB, _, err = sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB = sqlx.NewDb(mockDB, "sqlmock")
	mockUsers.EXPECT().InnerDB().Return(sqlxDB).Times(1)
	mockUsers.EXPECT().Ping().Return(fmt.Errorf("ping error")).Times(1)

	s = &storage.DBStorage{
		Users: mockUsers,
	}

	err = s.Ping()
	assert.Error(t, err, "Ping() should return an error for failed ping")
	assert.Equal(t, "ping error", err.Error(), "Ping() should return the correct error message for failed ping")
}

func TestDBStorage_CreateUser(t *testing.T) {
	// Test case 1: Nil DBStorage
	var nilStorage *storage.DBStorage
	otcp, err := nilStorage.CreateUser(context.Background(), nil)
	assert.Error(t, err, "CreateUser() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "CreateUser() should return the correct error message for nil DBStorage")
	require.Nil(t, otcp)

	// Test case 2: Nil Users repository
	s := &storage.DBStorage{}
	otcp, err = s.CreateUser(context.Background(), nil)
	assert.Error(t, err, "CreateUser() should return an error when Users repository is nil")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "CreateUser() should return the correct error message for nil Users repository")
	require.Nil(t, otcp)

	// Test case 3: Nil InnerDB
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsers := mocks.NewMockUsersRepositoryInterface(mockCtrl)
	s = &storage.DBStorage{
		Users: mockUsers,
	}

	mockUsers.EXPECT().InnerDB().Return(nil).Times(1)

	otcp, err = s.CreateUser(context.Background(), nil)
	assert.Error(t, err, "CreateUser() should return an error when InnerDB is nil")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "CreateUser() should return the correct error message for nil InnerDB")
	require.Nil(t, otcp)
	// Test case 4: Successful CreateUser
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	mockUsers.EXPECT().InnerDB().Return(sqlxDB).Times(1)
	mockUsers.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	s = &storage.DBStorage{
		Users: mockUsers,
	}

	userMap := map[string]interface{}{
		"username": "testuser",
		"email":    "testuser@example.com",
	}

	otcp, err = s.CreateUser(context.Background(), userMap)
	assert.NoError(t, err, "CreateUser() should not return an error for successful creation")
	require.NotNil(t, otcp)
	// Test case 5: Failed CreateUser
	mockDB, _, err = sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB = sqlx.NewDb(mockDB, "sqlmock")
	mockUsers.EXPECT().InnerDB().Return(sqlxDB).Times(1)
	mockUsers.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	s = &storage.DBStorage{
		Users: mockUsers,
	}

	otpc, err := s.CreateUser(context.Background(), userMap)

	require.NoError(t, err)
	require.NotNil(t, otpc)
}

func TestDBStorage_GetUser(t *testing.T) {
	// Test case 1: Nil DBStorage
	var nilStorage *storage.DBStorage
	user, err := nilStorage.GetUser(context.Background(), "test@example.com")
	assert.Nil(t, user, "GetUser() should return nil user for nil DBStorage")
	assert.Error(t, err, "GetUser() should return an error for nil DBStorage")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "GetUser() should return the correct error message for nil DBStorage")

	// Test case 2: Nil Users repository
	s := &storage.DBStorage{}
	user, err = s.GetUser(context.Background(), "test@example.com")
	assert.Nil(t, user, "GetUser() should return nil user when Users repository is nil")
	assert.Error(t, err, "GetUser() should return an error when Users repository is nil")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "GetUser() should return the correct error message for nil Users repository")

	// Test case 3: Nil InnerDB
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsers := mocks.NewMockUsersRepositoryInterface(mockCtrl)
	s = &storage.DBStorage{
		Users: mockUsers,
	}

	mockUsers.EXPECT().InnerDB().Return(nil).Times(1)

	user, err = s.GetUser(context.Background(), "test@example.com")
	assert.Nil(t, user, "GetUser() should return nil user when InnerDB is nil")
	assert.Error(t, err, "GetUser() should return an error when InnerDB is nil")
	assert.Equal(t, "[storage]: DBStorage is nil or not initialized", err.Error(), "GetUser() should return the correct error message for nil InnerDB")

	// Test case 4: Successful GetUser
	mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsers = mocks.NewMockUsersRepositoryInterface(mockCtrl)
	s = &storage.DBStorage{
		Users: mockUsers,
	}

	mockUsers.EXPECT().InnerDB().Return(&sqlx.DB{}).Times(1)

	expectedUser := &repository.User{
		ID:    1,
		Email: "test@example.com",
	}
	mockUsers.EXPECT().GetOne(gomock.Any(), "test@example.com").Return(expectedUser, nil).Times(1)

	user, err = s.GetUser(context.Background(), "test@example.com")
	assert.NoError(t, err, "GetUser() should not return an error for successful retrieval")
	assert.Equal(t, expectedUser, user, "GetUser() should return the correct user")

	// Test case 5: Failed GetUser
	mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsers = mocks.NewMockUsersRepositoryInterface(mockCtrl)
	s = &storage.DBStorage{
		Users: mockUsers,
	}

	mockUsers.EXPECT().InnerDB().Return(&sqlx.DB{}).Times(1)
	mockUsers.EXPECT().GetOne(gomock.Any(), "test@example.com").Return(nil, fmt.Errorf("user not found")).Times(1)

	user, err = s.GetUser(context.Background(), "test@example.com")
	assert.Nil(t, user, "GetUser() should return nil user for failed retrieval")
	assert.Error(t, err, "GetUser() should return an error for failed retrieval")
	assert.Equal(t, "user not found", err.Error(), "GetUser() should return the correct error message for failed retrieval")
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
