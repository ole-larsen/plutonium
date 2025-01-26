package mocks_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockCreateAndSellRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCreateAndSellRepositoryInterface(ctrl)

	// Test Create method
	ctx := context.Background()
	createMap := map[string]any{"key": "value"}
	mockRepo.EXPECT().Create(ctx, createMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, createMap)
	assert.NoError(t, err)

	// Test GetCreateAndSell method
	mockCreateAndSell := []*models.CreateAndSell{{}}
	mockRepo.EXPECT().GetCreateAndSell(ctx).Return(mockCreateAndSell, nil).Times(1)
	items, err := mockRepo.GetCreateAndSell(ctx)
	assert.NoError(t, err)
	assert.Equal(t, mockCreateAndSell, items)

	// Test GetCreateAndSellByID method
	id := int64(123)
	mockCreateAndSellItem := &models.CreateAndSell{}
	mockRepo.EXPECT().GetCreateAndSellByID(ctx, id).Return(mockCreateAndSellItem, nil).Times(1)
	item, err := mockRepo.GetCreateAndSellByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, mockCreateAndSellItem, item)

	// Test GetPublicCreateAndSell method
	mockPublicItems := []*models.PublicCreateAndSellItem{{}}
	mockRepo.EXPECT().GetPublicCreateAndSell(ctx).Return(mockPublicItems, nil).Times(1)
	publicItems, err := mockRepo.GetPublicCreateAndSell(ctx)
	assert.NoError(t, err)
	assert.Equal(t, mockPublicItems, publicItems)

	// Test InnerDB method
	mockDB := &sqlx.DB{}
	mockRepo.EXPECT().InnerDB().Return(mockDB).Times(1)
	db := mockRepo.InnerDB()
	assert.Equal(t, mockDB, db)

	// Test Ping method
	mockRepo.EXPECT().Ping().Return(nil).Times(1)
	err = mockRepo.Ping()
	assert.NoError(t, err)

	// Test Update method
	updateMap := map[string]any{"update_key": "update_value"}
	mockRepo.EXPECT().Update(ctx, updateMap).Return(mockCreateAndSell, nil).Times(1)
	updatedItems, err := mockRepo.Update(ctx, updateMap)
	assert.NoError(t, err)
	assert.Equal(t, mockCreateAndSell, updatedItems)
}
