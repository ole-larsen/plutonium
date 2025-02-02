package mocks_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	repository "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockUsersRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUsersRepositoryInterface(ctrl)
	ctx := context.Background()

	// Sample data
	sampleUser := &models.PublicUser{
		ID: 1,
		Attributes: &models.PublicUserAttributes{
			Username: "testuser",
			Email:    "test@example.com",
		},
	}

	sampleRepoUser := &repository.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	userMap := map[string]any{
		"username": "updatedUser",
		"email":    "updated@example.com",
	}

	// Test: Create
	t.Run("Create", func(t *testing.T) {
		mockRepo.EXPECT().Create(ctx, userMap).Return(nil)

		err := mockRepo.Create(ctx, userMap)
		assert.NoError(t, err)
	})

	t.Run("Create with error", func(t *testing.T) {
		mockRepo.EXPECT().Create(ctx, userMap).Return(errors.New("create error"))

		err := mockRepo.Create(ctx, userMap)
		assert.Error(t, err)
		assert.EqualError(t, err, "create error")
	})

	// Test: GetPublicUserByID
	t.Run("GetPublicUserByID", func(t *testing.T) {
		mockRepo.EXPECT().GetPublicUserByID(ctx, int64(1)).Return(sampleUser, nil)

		user, err := mockRepo.GetPublicUserByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, sampleUser, user)
	})

	t.Run("GetPublicUserByID with error", func(t *testing.T) {
		mockRepo.EXPECT().GetPublicUserByID(ctx, int64(1)).Return(nil, errors.New("not found"))

		user, err := mockRepo.GetPublicUserByID(ctx, 1)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "not found")
	})

	// Test: GetUserByAddress
	t.Run("GetUserByAddress", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByAddress(ctx, "address123").Return(sampleRepoUser, nil)

		user, err := mockRepo.GetUserByAddress(ctx, "address123")
		assert.NoError(t, err)
		assert.Equal(t, sampleRepoUser, user)
	})

	t.Run("GetUserByAddress with error", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByAddress(ctx, "address123").Return(nil, errors.New("address not found"))

		user, err := mockRepo.GetUserByAddress(ctx, "address123")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "address not found")
	})

	// Test: GetUserByEmail
	t.Run("GetUserByEmail", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByEmail(ctx, "test@example.com").Return(sampleRepoUser, nil)

		user, err := mockRepo.GetUserByEmail(ctx, "test@example.com")
		assert.NoError(t, err)
		assert.Equal(t, sampleRepoUser, user)
	})

	t.Run("GetUserByEmail with error", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByEmail(ctx, "test@example.com").Return(nil, errors.New("email not found"))

		user, err := mockRepo.GetUserByEmail(ctx, "test@example.com")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "email not found")
	})

	// Test: GetUserByID
	t.Run("GetUserByID", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(ctx, int64(1)).Return(sampleRepoUser, nil)

		user, err := mockRepo.GetUserByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, sampleRepoUser, user)
	})

	t.Run("GetUserByID with error", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(ctx, int64(1)).Return(nil, errors.New("user not found"))

		user, err := mockRepo.GetUserByID(ctx, 1)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})

	// Test: InnerDB
	t.Run("InnerDB", func(t *testing.T) {
		mockDB := &sqlx.DB{}
		mockRepo.EXPECT().InnerDB().Return(mockDB)

		db := mockRepo.InnerDB()
		assert.Equal(t, mockDB, db)
	})

	// Test: Ping
	t.Run("Ping", func(t *testing.T) {
		mockRepo.EXPECT().Ping().Return(nil)

		err := mockRepo.Ping()
		assert.NoError(t, err)
	})

	t.Run("Ping with error", func(t *testing.T) {
		mockRepo.EXPECT().Ping().Return(errors.New("ping error"))

		err := mockRepo.Ping()
		assert.Error(t, err)
		assert.EqualError(t, err, "ping error")
	})

	// Test: UpdateGravatar
	t.Run("UpdateGravatar", func(t *testing.T) {
		mockRepo.EXPECT().UpdateGravatar(ctx, userMap).Return(nil)

		err := mockRepo.UpdateGravatar(ctx, userMap)
		assert.NoError(t, err)
	})

	// Test: UpdateNonce
	t.Run("UpdateNonce", func(t *testing.T) {
		mockRepo.EXPECT().UpdateNonce(ctx, userMap).Return(nil)

		err := mockRepo.UpdateNonce(ctx, userMap)
		assert.NoError(t, err)
	})

	// Test: UpdateSecret
	t.Run("UpdateSecret", func(t *testing.T) {
		mockRepo.EXPECT().UpdateSecret(ctx, userMap).Return(nil)

		err := mockRepo.UpdateSecret(ctx, userMap)
		assert.NoError(t, err)
	})

	// Test: UpdateWallpaper
	t.Run("UpdateWallpaper", func(t *testing.T) {
		mockRepo.EXPECT().UpdateWallpaper(ctx, userMap).Return(nil)

		err := mockRepo.UpdateWallpaper(ctx, userMap)
		assert.NoError(t, err)
	})
}
