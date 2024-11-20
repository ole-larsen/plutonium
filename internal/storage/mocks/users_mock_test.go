package mocks_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"go.uber.org/mock/gomock"
)

func TestMockUsersRepositoryInterface_InnerDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	mockDB := &sqlx.DB{}
	mockRepo.EXPECT().InnerDB().Return(mockDB).Times(1)

	db := mockRepo.InnerDB()
	if db != mockDB {
		t.Errorf("Expected InnerDB to return %v, got %v", mockDB, db)
	}
}

func TestMockUsersRepositoryInterface_MigrateContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	ctx := context.Background()

	// Test successful migration
	mockRepo.EXPECT().MigrateContext(ctx).Return(nil).Times(1)

	if err := mockRepo.MigrateContext(ctx); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test migration with error
	mockError := &repository.Error{Err: fmt.Errorf("migration error")}
	mockRepo.EXPECT().MigrateContext(ctx).Return(mockError).Times(1)

	if err := mockRepo.MigrateContext(ctx); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockUsersRepositoryInterface_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)

	// Test successful ping
	mockRepo.EXPECT().Ping().Return(nil).Times(1)

	if err := mockRepo.Ping(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test ping with error
	mockError := fmt.Errorf("ping error")
	mockRepo.EXPECT().Ping().Return(mockError).Times(1)

	if err := mockRepo.Ping(); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockUsersRepositoryInterface_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	ctx := context.Background()
	userMap := map[string]interface{}{
		"email":    "test@example.com",
		"password": "hashedpassword",
		"secret":   "mysecret",
	}

	// Test successful creation
	mockRepo.EXPECT().Create(ctx, userMap).Return(nil).Times(1)

	if err := mockRepo.Create(ctx, userMap); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test creation with error
	mockError := fmt.Errorf("creation error")
	mockRepo.EXPECT().Create(ctx, userMap).Return(mockError).Times(1)

	if err := mockRepo.Create(ctx, userMap); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockUsersRepositoryInterface_GetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	ctx := context.Background()
	email := "test@example.com"

	// Test successful retrieval
	expectedUser := &repository.User{Email: email}
	mockRepo.EXPECT().GetOne(ctx, email).Return(expectedUser, nil).Times(1)

	user, err := mockRepo.GetOne(ctx, email)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user != expectedUser {
		t.Errorf("Expected user %v, got %v", expectedUser, user)
	}

	// Test GetOne with user not found error
	mockError := fmt.Errorf("user not found")
	mockRepo.EXPECT().GetOne(ctx, email).Return(nil, mockError).Times(1)

	user, err = mockRepo.GetOne(ctx, email)
	if err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}

	if user != nil {
		t.Errorf("Expected user to be nil, got %v", user)
	}

	// Test GetOne with database error
	dbError := fmt.Errorf("db error")
	mockRepo.EXPECT().GetOne(ctx, email).Return(nil, dbError).Times(1)

	user, err = mockRepo.GetOne(ctx, email)
	if err == nil || err.Error() != dbError.Error() {
		t.Errorf("Expected error %v, got %v", dbError, err)
	}

	if user != nil {
		t.Errorf("Expected user to be nil, got %v", user)
	}
}
