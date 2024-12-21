package mocks_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	repository "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const testEmail = "test@example.com"

func setupMockController(t *testing.T) (*gomock.Controller, *mocks.MockDBStorageInterface) {
	t.Helper()
	ctrl := gomock.NewController(t)

	return ctrl, mocks.NewMockDBStorageInterface(ctrl)
}

func TestMockDBStorageInterface_ConnectUsersRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	mockSQLxDB := &sqlx.DB{}

	// Test successful connection
	mockDBStorage.EXPECT().ConnectUsersRepository(ctx, mockSQLxDB).Return(nil).Times(1)

	if err := mockDBStorage.ConnectUsersRepository(ctx, mockSQLxDB); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test connection with error
	mockError := fmt.Errorf("connection error")
	mockDBStorage.EXPECT().ConnectUsersRepository(ctx, mockSQLxDB).Return(mockError).Times(1)

	if err := mockDBStorage.ConnectUsersRepository(ctx, mockSQLxDB); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_ConnectContractsRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	mockSQLxDB := &sqlx.DB{}

	// Test successful connection
	mockDBStorage.EXPECT().ConnectContractsRepository(ctx, mockSQLxDB).Return(nil).Times(1)

	if err := mockDBStorage.ConnectContractsRepository(ctx, mockSQLxDB); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test connection with error
	mockError := fmt.Errorf("connection error")
	mockDBStorage.EXPECT().ConnectContractsRepository(ctx, mockSQLxDB).Return(mockError).Times(1)

	if err := mockDBStorage.ConnectContractsRepository(ctx, mockSQLxDB); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)
	ctx := context.Background()
	sqlxDB := &sqlx.DB{} // Mock or create a dummy *sqlx.DB instance

	// Test successful initialization
	mockDBStorage.EXPECT().Init(ctx, sqlxDB).Return(sqlxDB, nil).Times(1)

	db, err := mockDBStorage.Init(ctx, sqlxDB)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if db != sqlxDB {
		t.Errorf("Expected %v, got %v", sqlxDB, db)
	}

	// Test initialization with error
	mockError := fmt.Errorf("init error")
	mockDBStorage.EXPECT().Init(ctx, sqlxDB).Return(nil, mockError).Times(1)

	db, err = mockDBStorage.Init(ctx, sqlxDB)
	if err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}

	if db != nil {
		t.Errorf("Expected nil, got %v", db)
	}
}

func TestMockDBStorageInterface_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	// Test successful ping
	mockDBStorage.EXPECT().Ping().Return(nil).Times(1)

	if err := mockDBStorage.Ping(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test ping with error
	mockError := fmt.Errorf("ping error")
	mockDBStorage.EXPECT().Ping().Return(mockError).Times(1)

	if err := mockDBStorage.Ping(); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)
	ctx := context.Background()
	userMap := map[string]interface{}{"key": "value"}

	// Test successful user creation
	mockDBStorage.EXPECT().CreateUser(ctx, userMap).Return(nil, nil).Times(1)

	if _, err := mockDBStorage.CreateUser(ctx, userMap); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test user creation with error
	mockError := fmt.Errorf("create user error")
	mockDBStorage.EXPECT().CreateUser(ctx, userMap).Return(nil, mockError).Times(1)

	if _, err := mockDBStorage.CreateUser(ctx, userMap); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)
	ctx := context.Background()
	email := testEmail
	expectedUser := &repository.User{
		ID:    1,
		Email: email,
	}

	// Test successful user retrieval
	mockDBStorage.EXPECT().GetUser(ctx, email).Return(expectedUser, nil).Times(1)

	user, err := mockDBStorage.GetUser(ctx, email)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user != expectedUser {
		t.Errorf("Expected %v, got %v", expectedUser, user)
	}

	// Test user retrieval with error
	mockError := fmt.Errorf("get user error")
	mockDBStorage.EXPECT().GetUser(ctx, email).Return(nil, mockError).Times(1)

	user, err = mockDBStorage.GetUser(ctx, email)
	if user != nil {
		t.Errorf("Expected nil user, got %v", user)
	}

	if err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_GetUser2(t *testing.T) {
	ctrl, mockDBStorage := setupMockController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	email := testEmail
	expectedUser := &repository.User{
		ID:    1,
		Email: email,
	}

	t.Run("Success", func(t *testing.T) {
		mockDBStorage.EXPECT().GetUser(ctx, email).Return(expectedUser, nil)

		user, err := mockDBStorage.GetUser(ctx, email)

		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, expectedUser, user, "Expected user to match")
	})

	t.Run("Error", func(t *testing.T) {
		mockError := fmt.Errorf("get user error")
		mockDBStorage.EXPECT().GetUser(ctx, email).Return(nil, mockError)

		user, err := mockDBStorage.GetUser(ctx, email)

		assert.Nil(t, user, "Expected user to be nil")
		assert.ErrorIs(t, err, mockError, "Expected specific error")
	})
}

func TestMockDBStorageInterface_GetContractsRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)
	expectedRepo := &repository.ContractsRepository{}

	t.Run("Success", func(t *testing.T) {
		// Set expectation for GetContractsRepository call to return the mock repository.
		mockDBStorage.EXPECT().GetContractsRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetContractsRepository()

		// Assertions
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("NilRepository", func(t *testing.T) {
		// Set expectation for GetContractsRepository call to return nil.
		mockDBStorage.EXPECT().GetContractsRepository().Return(nil).Times(1)

		repo := mockDBStorage.GetContractsRepository()

		// Assertions
		if repo != nil {
			t.Errorf("Expected nil, got %v", repo)
		}
	})
}
