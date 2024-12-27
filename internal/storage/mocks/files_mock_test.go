package mocks_test

import (
	"context"
	"testing"

	sqlx "github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockFilesRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFilesRepositoryInterface(ctrl)
	ctx := context.TODO()

	t.Run("Test Create", func(t *testing.T) {
		mockRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)
		err := mockRepo.Create(ctx, map[string]interface{}{"name": "test"})
		assert.NoError(t, err)
	})

	t.Run("Test GetFileByID", func(t *testing.T) {
		expectedFile := &models.File{ID: 1, Name: "test"}
		mockRepo.EXPECT().GetFileByID(ctx, int64(1)).Return(expectedFile, nil)
		file, err := mockRepo.GetFileByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("Test GetFileByName", func(t *testing.T) {
		expectedFile := &models.File{Name: "test"}
		mockRepo.EXPECT().GetFileByName(ctx, "test").Return(expectedFile, nil)
		file, err := mockRepo.GetFileByName(ctx, "test")
		assert.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("Test GetFiles", func(t *testing.T) {
		expectedFiles := []*models.File{
			{
				ID:   1,
				Name: "file1",
			},
			{
				ID:   2,
				Name: "file2",
			},
		}
		mockRepo.EXPECT().GetFiles(ctx).Return(expectedFiles, nil)
		files, err := mockRepo.GetFiles(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedFiles, files)
	})

	t.Run("Test GetPublicFileByID", func(t *testing.T) {
		expectedFile := &models.PublicFile{
			ID: 1,
			Attributes: &models.PublicFileAttributes{
				Name: "public-file",
			},
		}
		mockRepo.EXPECT().GetPublicFileByID(ctx, int64(1)).Return(expectedFile, nil)
		file, err := mockRepo.GetPublicFileByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("Test GetPublicFileByName", func(t *testing.T) {
		expectedFile := &models.PublicFile{
			ID: 1,
			Attributes: &models.PublicFileAttributes{
				Name: "public-file",
			},
		}
		mockRepo.EXPECT().GetPublicFileByName(ctx, "public-file").Return(expectedFile, nil)
		file, err := mockRepo.GetPublicFileByName(ctx, "public-file")
		assert.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("Test GetPublicFilesByProvider", func(t *testing.T) {
		expectedFiles := []*models.PublicFile{
			{
				ID: 1,
				Attributes: &models.PublicFileAttributes{
					Name: "public-file1",
				},
			},
			{
				ID: 2,
				Attributes: &models.PublicFileAttributes{
					Name: "public-file2",
				},
			},
		}
		mockRepo.EXPECT().GetPublicFilesByProvider(ctx, "provider").Return(expectedFiles, nil)
		files, err := mockRepo.GetPublicFilesByProvider(ctx, "provider")
		assert.NoError(t, err)
		assert.Equal(t, expectedFiles, files)
	})

	t.Run("Test InnerDB", func(t *testing.T) {
		mockDB := &sqlx.DB{}
		mockRepo.EXPECT().InnerDB().Return(mockDB)
		db := mockRepo.InnerDB()
		assert.Equal(t, mockDB, db)
	})

	t.Run("Test Ping", func(t *testing.T) {
		mockRepo.EXPECT().Ping().Return(nil)
		err := mockRepo.Ping()
		assert.NoError(t, err)
	})

	t.Run("Test Update", func(t *testing.T) {
		expectedFiles := []*models.File{
			{ID: 1, Name: "updated-file1"},
			{ID: 2, Name: "updated-file2"},
		}
		mockRepo.EXPECT().Update(ctx, gomock.Any()).Return(expectedFiles, nil)
		files, err := mockRepo.Update(ctx, map[string]interface{}{"name": "updated"})
		assert.NoError(t, err)
		assert.Equal(t, expectedFiles, files)
	})
}
