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

func TestMockBlogsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogsRepositoryInterface(ctrl)

	// Test Create method
	blogMap := map[string]any{"title": "Test Blog", "content": "This is a test blog."}
	mockRepo.EXPECT().Create(context.Background(), blogMap).Return(nil).Times(1)
	err := mockRepo.Create(context.Background(), blogMap)
	assert.NoError(t, err)

	// Test GetBlogByID method
	mockBlog := &models.Blog{ID: 1, Title: "Test Blog"}
	mockRepo.EXPECT().GetBlogByID(context.Background(), int64(1)).Return(mockBlog, nil).Times(1)
	blog, err := mockRepo.GetBlogByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, mockBlog, blog)

	// Test GetBlogs method
	mockBlogs := []*models.Blog{{ID: 1, Title: "Test Blog"}}
	mockRepo.EXPECT().GetBlogs(context.Background()).Return(mockBlogs, nil).Times(1)
	blogs, err := mockRepo.GetBlogs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, mockBlogs, blogs)

	// Test GetPublicBlogItem method
	mockPublicBlogItem := &models.PublicBlogItem{ID: 1, Title: "Test Blog"}
	mockRepo.EXPECT().GetPublicBlogItem(context.Background(), "test-blog").Return(mockPublicBlogItem, nil).Times(1)
	publicBlogItem, err := mockRepo.GetPublicBlogItem(context.Background(), "test-blog")
	assert.NoError(t, err)
	assert.Equal(t, mockPublicBlogItem, publicBlogItem)

	// Test GetPublicBlogs method
	mockPublicBlogs := []*models.PublicBlogItem{{ID: 1, Title: "Test Blog"}}
	mockRepo.EXPECT().GetPublicBlogs(context.Background()).Return(mockPublicBlogs, nil).Times(1)
	publicBlogs, err := mockRepo.GetPublicBlogs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, mockPublicBlogs, publicBlogs)

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
	updatedBlogs := []*models.Blog{{ID: 1, Title: "Updated Blog"}}
	mockRepo.EXPECT().Update(context.Background(), blogMap).Return(updatedBlogs, nil).Times(1)
	blogs, err = mockRepo.Update(context.Background(), blogMap)
	assert.NoError(t, err)
	assert.Equal(t, updatedBlogs, blogs)
}
