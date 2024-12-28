package mocks_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
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

func TestMockUsersRepositoryInterface_GetPublicUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	ctx := context.Background()
	userID := int64(123)

	// Test successful retrieval of public user
	expectedPublicUser := &models.PublicUser{
		ID:       userID,
		Username: "testuser",
	}
	mockRepo.EXPECT().GetPublicUserByID(ctx, userID).Return(expectedPublicUser, nil).Times(1)

	publicUser, err := mockRepo.GetPublicUserByID(ctx, userID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if publicUser != expectedPublicUser {
		t.Errorf("Expected public user %v, got %v", expectedPublicUser, publicUser)
	}

	// Test GetPublicUserByID with user not found error
	mockError := fmt.Errorf("public user not found")
	mockRepo.EXPECT().GetPublicUserByID(ctx, userID).Return(nil, mockError).Times(1)

	publicUser, err = mockRepo.GetPublicUserByID(ctx, userID)
	if err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}

	if publicUser != nil {
		t.Errorf("Expected public user to be nil, got %v", publicUser)
	}

	// Test GetPublicUserByID with database error
	dbError := fmt.Errorf("db error")
	mockRepo.EXPECT().GetPublicUserByID(ctx, userID).Return(nil, dbError).Times(1)

	publicUser, err = mockRepo.GetPublicUserByID(ctx, userID)
	if err == nil || err.Error() != dbError.Error() {
		t.Errorf("Expected error %v, got %v", dbError, err)
	}

	if publicUser != nil {
		t.Errorf("Expected public user to be nil, got %v", publicUser)
	}
}

func TestMockUsersRepositoryInterface_GetUserByAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	ctx := context.Background()
	address := "0x1234567890abcdef"

	// Test successful retrieval
	expectedUser := &repository.User{
		ID:      1,
		Email:   "test@example.com",
		Address: address,
	}
	mockRepo.EXPECT().GetUserByAddress(ctx, address).Return(expectedUser, nil).Times(1)

	user, err := mockRepo.GetUserByAddress(ctx, address)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user != expectedUser {
		t.Errorf("Expected user %v, got %v", expectedUser, user)
	}

	// Test GetUserByAddress with user not found error
	notFoundError := fmt.Errorf("user not found")
	mockRepo.EXPECT().GetUserByAddress(ctx, address).Return(nil, notFoundError).Times(1)

	user, err = mockRepo.GetUserByAddress(ctx, address)
	if err == nil || err.Error() != notFoundError.Error() {
		t.Errorf("Expected error %v, got %v", notFoundError, err)
	}

	if user != nil {
		t.Errorf("Expected user to be nil, got %v", user)
	}

	// Test GetUserByAddress with database error
	dbError := fmt.Errorf("database error")
	mockRepo.EXPECT().GetUserByAddress(ctx, address).Return(nil, dbError).Times(1)

	user, err = mockRepo.GetUserByAddress(ctx, address)
	if err == nil || err.Error() != dbError.Error() {
		t.Errorf("Expected error %v, got %v", dbError, err)
	}

	if user != nil {
		t.Errorf("Expected user to be nil, got %v", user)
	}
}
