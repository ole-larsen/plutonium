package mocks_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMockCategoriesRepositoryInterface_InnerDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoriesRepositoryInterface(ctrl)
	db := &sqlx.DB{}

	// Expectation
	mockRepo.EXPECT().InnerDB().Return(db).Times(1)

	// Call the method
	result := mockRepo.InnerDB()

	// Assert
	assert.Equal(t, db, result, "InnerDB should return the expected database instance")
}

func TestMockCategoriesRepositoryInterface_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoriesRepositoryInterface(ctrl)

	// Test success case
	mockRepo.EXPECT().Ping().Return(nil).Times(1)

	err := mockRepo.Ping()
	assert.NoError(t, err, "Ping should not return an error in success case")

	// Test failure case
	mockRepo.EXPECT().Ping().Return(errors.New("ping failed")).Times(1)

	err = mockRepo.Ping()
	assert.Error(t, err, "Ping should return an error in failure case")
	assert.Equal(t, "ping failed", err.Error())
}

func TestMockCategoriesRepositoryInterface_GetPublicCollectibleCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoriesRepositoryInterface(ctrl)
	mockUsersRepo := &repository.UsersRepository{}
	ctx := context.TODO()

	// Test success case
	expectedCategories := []*models.PublicCategory{
		{
			ID: 1,
			Attributes: &models.PublicCategoryAttributes{
				Title: "Category 1",
			},
		},
		{
			ID: 2,
			Attributes: &models.PublicCategoryAttributes{
				Title: "Category 2",
			},
		},
	}
	mockRepo.EXPECT().GetPublicCollectibleCategories(ctx, mockUsersRepo).Return(expectedCategories, nil).Times(1)

	categories, err := mockRepo.GetPublicCollectibleCategories(ctx, mockUsersRepo)
	assert.NoError(t, err, "GetPublicCollectibleCategories should not return an error in success case")
	assert.Equal(t, expectedCategories, categories, "GetPublicCollectibleCategories should return the expected categories")

	// Test failure case
	mockRepo.EXPECT().GetPublicCollectibleCategories(ctx, mockUsersRepo).Return(nil, errors.New("failed to get categories")).Times(1)

	categories, err = mockRepo.GetPublicCollectibleCategories(ctx, mockUsersRepo)
	assert.Error(t, err, "GetPublicCollectibleCategories should return an error in failure case")
	assert.Nil(t, categories, "GetPublicCollectibleCategories should return nil categories in failure case")
	assert.Equal(t, "failed to get categories", err.Error())
}
