package mocks_test

import (
	"context"
	"testing"

	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockHelpCenterRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockHelpCenterRepositoryInterface(ctrl)
	ctx := context.Background()

	// Example data
	helpCenterID := int64(1)
	helpCenter := &models.HelpCenter{ID: helpCenterID, Title: "Test Help Center"}
	helpCenters := []*models.HelpCenter{helpCenter}
	helpCenterMap := map[string]interface{}{"id": helpCenterID, "title": "Updated Help Center"}
	publicHelpCenterItem := &models.PublicHelpCenterItem{ID: helpCenterID, Title: "Public Help Center"}
	publicHelpCenters := []*models.PublicHelpCenterItem{publicHelpCenterItem}

	// Test Create
	mockRepo.EXPECT().Create(ctx, helpCenterMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, helpCenterMap)
	assert.NoError(t, err, "Create method should not return an error")

	// Test GetHelpCenter
	mockRepo.EXPECT().GetHelpCenter(ctx).Return(helpCenters, nil).Times(1)
	returnedHelpCenters, err := mockRepo.GetHelpCenter(ctx)
	assert.NoError(t, err, "GetHelpCenter method should not return an error")
	assert.Equal(t, helpCenters, returnedHelpCenters, "GetHelpCenter should return the expected help centers")

	// Test GetHelpCenterByID
	mockRepo.EXPECT().GetHelpCenterByID(ctx, helpCenterID).Return(helpCenter, nil).Times(1)
	returnedHelpCenter, err := mockRepo.GetHelpCenterByID(ctx, helpCenterID)
	assert.NoError(t, err, "GetHelpCenterByID method should not return an error")
	assert.Equal(t, helpCenter, returnedHelpCenter, "GetHelpCenterByID should return the expected help center")

	// Test GetPublicHelpCenter
	mockRepo.EXPECT().GetPublicHelpCenter(ctx).Return(publicHelpCenters, nil).Times(1)
	returnedPublicHelpCenters, err := mockRepo.GetPublicHelpCenter(ctx)
	assert.NoError(t, err, "GetPublicHelpCenter method should not return an error")
	assert.Equal(t, publicHelpCenters, returnedPublicHelpCenters, "GetPublicHelpCenter should return the expected public help center items")

	// Test InnerDB
	mockRepo.EXPECT().InnerDB().Return(nil).Times(1)
	db := mockRepo.InnerDB()
	assert.Nil(t, db, "InnerDB should return nil for this mock test")

	// Test Ping
	mockRepo.EXPECT().Ping().Return(nil).Times(1)
	err = mockRepo.Ping()
	assert.NoError(t, err, "Ping method should not return an error")

	// Test Update
	mockRepo.EXPECT().Update(ctx, helpCenterMap).Return(helpCenters, nil).Times(1)
	updatedHelpCenters, err := mockRepo.Update(ctx, helpCenterMap)
	assert.NoError(t, err, "Update method should not return an error")
	assert.Equal(t, helpCenters, updatedHelpCenters, "Update should return the updated help centers")
}
