package mocks_test

import (
	"context"
	"testing"

	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockTagsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTagsRepositoryInterface(ctrl)
	ctx := context.Background()

	// Example data
	tagID := int64(1)
	tag := &models.Tag{ID: tagID, Title: "Test Tag"}
	tags := []*models.Tag{tag}
	tagMap := map[string]interface{}{"id": tagID, "name": "Updated Tag"}

	// Test Create
	mockRepo.EXPECT().Create(ctx, tagMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, tagMap)
	assert.NoError(t, err, "Create method should not return an error")

	// Test GetTagByID
	mockRepo.EXPECT().GetTagByID(ctx, tagID).Return(tag, nil).Times(1)
	returnedTag, err := mockRepo.GetTagByID(ctx, tagID)
	assert.NoError(t, err, "GetTagByID method should not return an error")
	assert.Equal(t, tag, returnedTag, "GetTagByID should return the expected tag")

	// Test GetTags
	mockRepo.EXPECT().GetTags(ctx).Return(tags, nil).Times(1)
	returnedTags, err := mockRepo.GetTags(ctx)
	assert.NoError(t, err, "GetTags method should not return an error")
	assert.Equal(t, tags, returnedTags, "GetTags should return the expected tags")

	// Test InnerDB
	mockRepo.EXPECT().InnerDB().Return(nil).Times(1)
	db := mockRepo.InnerDB()
	assert.Nil(t, db, "InnerDB should return nil for this mock test")

	// Test Ping
	mockRepo.EXPECT().Ping().Return(nil).Times(1)
	err = mockRepo.Ping()
	assert.NoError(t, err, "Ping method should not return an error")

	// Test Update
	mockRepo.EXPECT().Update(ctx, tagMap).Return(tags, nil).Times(1)
	updatedTags, err := mockRepo.Update(ctx, tagMap)
	assert.NoError(t, err, "Update method should not return an error")
	assert.Equal(t, tags, updatedTags, "Update should return the updated tags")
}
