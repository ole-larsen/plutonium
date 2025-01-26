package mocks_test

import (
	"context"
	"testing"

	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockPagesRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPagesRepositoryInterface(ctrl)
	ctx := context.Background()

	// Example data
	pageID := int64(1)
	page := &models.Page{ID: pageID, Title: "Test Page"}
	pages := []*models.Page{page}
	pageMap := map[string]interface{}{"id": pageID, "title": "Updated Page"}
	slug := "test-slug"
	publicPage := &models.PublicPage{ID: pageID, Attributes: &models.PublicPageAttributes{
		Title: slug,
	}}

	// Test Create
	mockRepo.EXPECT().Create(ctx, pageMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, pageMap)
	assert.NoError(t, err, "Create method should not return an error")

	// Test GetPageByID
	mockRepo.EXPECT().GetPageByID(ctx, pageID).Return(page, nil).Times(1)
	returnedPage, err := mockRepo.GetPageByID(ctx, pageID)
	assert.NoError(t, err, "GetPageByID method should not return an error")
	assert.Equal(t, page, returnedPage, "GetPageByID should return the expected page")

	// Test GetPageBySlug
	mockRepo.EXPECT().GetPageBySlug(ctx, slug).Return(publicPage, nil).Times(1)
	returnedPublicPage, err := mockRepo.GetPageBySlug(ctx, slug)
	assert.NoError(t, err, "GetPageBySlug method should not return an error")
	assert.Equal(t, publicPage, returnedPublicPage, "GetPageBySlug should return the expected public page")

	// Test GetPages
	mockRepo.EXPECT().GetPages(ctx).Return(pages, nil).Times(1)
	returnedPages, err := mockRepo.GetPages(ctx)
	assert.NoError(t, err, "GetPages method should not return an error")
	assert.Equal(t, pages, returnedPages, "GetPages should return the expected pages")

	// Test InnerDB
	mockRepo.EXPECT().InnerDB().Return(nil).Times(1)
	db := mockRepo.InnerDB()
	assert.Nil(t, db, "InnerDB should return nil for this mock test")

	// Test Ping
	mockRepo.EXPECT().Ping().Return(nil).Times(1)
	err = mockRepo.Ping()
	assert.NoError(t, err, "Ping method should not return an error")

	// Test Update
	mockRepo.EXPECT().Update(ctx, pageMap).Return(pages, nil).Times(1)
	updatedPages, err := mockRepo.Update(ctx, pageMap)
	assert.NoError(t, err, "Update method should not return an error")
	assert.Equal(t, pages, updatedPages, "Update should return the updated pages")
}
