package mocks_test

import (
	context "context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMockMenusRepositoryInterface_GetMenuByProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMenusRepositoryInterface(ctrl)

	provider := "test-provider"
	expectedMenu := &models.PublicMenu{
		ID: 1,
		Attributes: &models.PublicMenuAttributes{
			Name:  "Test Menu",
			Link:  "/test",
			Items: nil,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Success case
	mockRepo.EXPECT().GetMenuByProvider(ctx, provider).Return(expectedMenu, nil).Times(1)

	menu, err := mockRepo.GetMenuByProvider(ctx, provider)
	assert.NoError(t, err, "GetMenuByProvider should not return an error on success")
	assert.Equal(t, expectedMenu, menu, "GetMenuByProvider should return the expected menu")

	// Error case
	mockRepo.EXPECT().GetMenuByProvider(ctx, provider).Return(nil, errors.New("menu not found")).Times(1)

	menu, err = mockRepo.GetMenuByProvider(ctx, provider)
	assert.Error(t, err, "GetMenuByProvider should return an error if the menu is not found")
	assert.Nil(t, menu, "GetMenuByProvider should return nil when an error occurs")
}

func TestMockMenusRepositoryInterface_InnerDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMenusRepositoryInterface(ctrl)
	db := &sqlx.DB{}

	// Expectation
	mockRepo.EXPECT().InnerDB().Return(db).Times(1)

	// Call the method
	result := mockRepo.InnerDB()

	// Assert
	assert.Equal(t, db, result, "InnerDB should return the expected database instance")
}

func TestMockMenusRepositoryInterface_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMenusRepositoryInterface(ctrl)

	// Success case
	mockRepo.EXPECT().Ping().Return(nil).Times(1)

	err := mockRepo.Ping()
	assert.NoError(t, err, "Ping should not return an error on success")

	// Error case
	mockRepo.EXPECT().Ping().Return(errors.New("ping failed")).Times(1)

	err = mockRepo.Ping()
	assert.Error(t, err, "Ping should return an error when ping fails")
	assert.Equal(t, "ping failed", err.Error())
}
