package mocks_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	repository "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
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

func TestMockDBStorageInterface_ConnectPagesRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	mockSQLxDB := &sqlx.DB{}

	// Test successful connection
	mockDBStorage.EXPECT().ConnectPagesRepository(ctx, mockSQLxDB).Return(nil).Times(1)

	if err := mockDBStorage.ConnectPagesRepository(ctx, mockSQLxDB); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test connection with error
	mockError := fmt.Errorf("connection error")
	mockDBStorage.EXPECT().ConnectPagesRepository(ctx, mockSQLxDB).Return(mockError).Times(1)

	if err := mockDBStorage.ConnectPagesRepository(ctx, mockSQLxDB); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_ConnectCategoriesRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	mockSQLxDB := &sqlx.DB{}

	// Test successful connection
	mockDBStorage.EXPECT().ConnectCategoriesRepository(ctx, mockSQLxDB).Return(nil).Times(1)

	if err := mockDBStorage.ConnectCategoriesRepository(ctx, mockSQLxDB); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test connection with error
	mockError := fmt.Errorf("connection error")
	mockDBStorage.EXPECT().ConnectCategoriesRepository(ctx, mockSQLxDB).Return(mockError).Times(1)

	if err := mockDBStorage.ConnectCategoriesRepository(ctx, mockSQLxDB); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockDBStorageInterface_ConnectMenusRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	mockSQLxDB := &sqlx.DB{}

	// Test successful connection
	mockDBStorage.EXPECT().ConnectMenusRepository(ctx, mockSQLxDB).Return(nil).Times(1)

	if err := mockDBStorage.ConnectMenusRepository(ctx, mockSQLxDB); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test connection with error
	mockError := fmt.Errorf("connection error")
	mockDBStorage.EXPECT().ConnectMenusRepository(ctx, mockSQLxDB).Return(mockError).Times(1)

	if err := mockDBStorage.ConnectMenusRepository(ctx, mockSQLxDB); err == nil || err.Error() != mockError.Error() {
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

func TestMockDBStorageInterface_GetMenusRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)
	expectedRepo := &repository.MenusRepository{}

	t.Run("Success", func(t *testing.T) {
		// Set expectation for GetMenusRepository call to return the mock repository.
		mockDBStorage.EXPECT().GetMenusRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetMenusRepository()

		// Assertions
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("NilRepository", func(t *testing.T) {
		// Set expectation for GetMenusRepository call to return nil.
		mockDBStorage.EXPECT().GetMenusRepository().Return(nil).Times(1)

		repo := mockDBStorage.GetMenusRepository()

		// Assertions
		if repo != nil {
			t.Errorf("Expected nil, got %v", repo)
		}
	})
}
