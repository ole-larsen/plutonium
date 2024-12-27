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

func setupMockController(t *testing.T) (*gomock.Controller, *mocks.MockDBStorageInterface) {
	t.Helper()
	ctrl := gomock.NewController(t)

	return ctrl, mocks.NewMockDBStorageInterface(ctrl)
}

func TestMockDBStorageInterface_ConnectRepository(t *testing.T) {
	ctrl, mockDBStorage := setupMockController(t)
	defer ctrl.Finish()

	mockSQLxDB := &sqlx.DB{}

	t.Run("Successful Connection", func(t *testing.T) {
		mockDBStorage.EXPECT().ConnectRepository("users", mockSQLxDB).Return(nil).Times(1)

		err := mockDBStorage.ConnectRepository("users", mockSQLxDB)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Error on Connection", func(t *testing.T) {
		mockError := fmt.Errorf("connection error")
		mockDBStorage.EXPECT().ConnectRepository("users", mockSQLxDB).Return(mockError).Times(1)

		err := mockDBStorage.ConnectRepository("users", mockSQLxDB)
		if err == nil || err.Error() != mockError.Error() {
			t.Errorf("Expected error %v, got %v", mockError, err)
		}
	})
}

func TestMockDBStorageInterface_Init(t *testing.T) {
	ctrl, mockDBStorage := setupMockController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockSQLxDB := &sqlx.DB{}

	t.Run("Successful Initialization", func(t *testing.T) {
		mockDBStorage.EXPECT().Init(ctx, mockSQLxDB).Return(mockSQLxDB, nil).Times(1)

		db, err := mockDBStorage.Init(ctx, mockSQLxDB)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if db != mockSQLxDB {
			t.Errorf("Expected %v, got %v", mockSQLxDB, db)
		}
	})

	t.Run("Error on Initialization", func(t *testing.T) {
		mockError := fmt.Errorf("init error")
		mockDBStorage.EXPECT().Init(ctx, mockSQLxDB).Return(nil, mockError).Times(1)

		db, err := mockDBStorage.Init(ctx, mockSQLxDB)
		if err == nil || err.Error() != mockError.Error() {
			t.Errorf("Expected error %v, got %v", mockError, err)
		}

		if db != nil {
			t.Errorf("Expected nil, got %v", db)
		}
	})
}

func TestMockDBStorageInterface_GetRepositories(t *testing.T) {
	ctrl, mockDBStorage := setupMockController(t)
	defer ctrl.Finish()

	t.Run("GetCategoriesRepository", func(t *testing.T) {
		expectedRepo := &repository.CategoriesRepository{}
		mockDBStorage.EXPECT().GetCategoriesRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetCategoriesRepository()
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("GetContractsRepository", func(t *testing.T) {
		expectedRepo := &repository.ContractsRepository{}
		mockDBStorage.EXPECT().GetContractsRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetContractsRepository()
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("GetFilesRepository", func(t *testing.T) {
		expectedRepo := &repository.FilesRepository{}
		mockDBStorage.EXPECT().GetFilesRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetFilesRepository()
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("GetMenusRepository", func(t *testing.T) {
		expectedRepo := &repository.MenusRepository{}
		mockDBStorage.EXPECT().GetMenusRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetMenusRepository()
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("GetSlidersRepository", func(t *testing.T) {
		expectedRepo := &repository.SlidersRepository{}
		mockDBStorage.EXPECT().GetSlidersRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetSlidersRepository()
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})

	t.Run("GetUsersRepository", func(t *testing.T) {
		expectedRepo := &repository.UsersRepository{}
		mockDBStorage.EXPECT().GetUsersRepository().Return(expectedRepo).Times(1)

		repo := mockDBStorage.GetUsersRepository()
		if repo != expectedRepo {
			t.Errorf("Expected %v, got %v", expectedRepo, repo)
		}
	})
}

func TestMockDBStorageInterface_ErrorScenarios(t *testing.T) {
	ctrl, mockDBStorage := setupMockController(t)
	defer ctrl.Finish()

	t.Run("Nil Repositories", func(t *testing.T) {
		mockDBStorage.EXPECT().GetCategoriesRepository().Return(nil).Times(1)
		mockDBStorage.EXPECT().GetContractsRepository().Return(nil).Times(1)

		if repo := mockDBStorage.GetCategoriesRepository(); repo != nil {
			t.Errorf("Expected nil, got %v", repo)
		}

		if repo := mockDBStorage.GetContractsRepository(); repo != nil {
			t.Errorf("Expected nil, got %v", repo)
		}
	})
}
