package mocks_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockDBStorageInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStorage := mocks.NewMockDBStorageInterface(ctrl)

	// Test Init
	ctx := context.Background()
	mockSQLxDB := &sqlx.DB{}
	mockDBStorage.EXPECT().Init(ctx, mockSQLxDB).Return(mockSQLxDB, nil).Times(1)
	db, err := mockDBStorage.Init(ctx, mockSQLxDB)
	assert.NoError(t, err)
	assert.Equal(t, mockSQLxDB, db)

	// Test ConnectRepository
	repoName := "test_repo"
	mockDBStorage.EXPECT().ConnectRepository(repoName, mockSQLxDB).Return(nil).Times(1)
	err = mockDBStorage.ConnectRepository(repoName, mockSQLxDB)
	assert.NoError(t, err)

	// Test GetAuthorsRepository
	mockAuthorsRepo := &repository.AuthorsRepository{}
	mockDBStorage.EXPECT().GetAuthorsRepository().Return(mockAuthorsRepo).Times(1)
	authorsRepo := mockDBStorage.GetAuthorsRepository()
	assert.Equal(t, mockAuthorsRepo, authorsRepo)

	// Test GetBlogsRepository
	mockBlogsRepo := &repository.BlogsRepository{}
	mockDBStorage.EXPECT().GetBlogsRepository().Return(mockBlogsRepo).Times(1)
	blogsRepo := mockDBStorage.GetBlogsRepository()
	assert.Equal(t, mockBlogsRepo, blogsRepo)

	// Test GetCategoriesRepository
	mockCategoriesRepo := &repository.CategoriesRepository{}
	mockDBStorage.EXPECT().GetCategoriesRepository().Return(mockCategoriesRepo).Times(1)
	categoriesRepo := mockDBStorage.GetCategoriesRepository()
	assert.Equal(t, mockCategoriesRepo, categoriesRepo)

	// Test GetContactFormsRepository
	mockContactFormsRepo := &repository.ContactFormRepository{}
	mockDBStorage.EXPECT().GetContactFormsRepository().Return(mockContactFormsRepo).Times(1)
	contactFormsRepo := mockDBStorage.GetContactFormsRepository()
	assert.Equal(t, mockContactFormsRepo, contactFormsRepo)

	// Test GetContactsRepository
	mockContactsRepo := &repository.ContactsRepository{}
	mockDBStorage.EXPECT().GetContactsRepository().Return(mockContactsRepo).Times(1)
	contactsRepo := mockDBStorage.GetContactsRepository()
	assert.Equal(t, mockContactsRepo, contactsRepo)

	// Test GetContractsRepository
	mockContractsRepo := &repository.ContractsRepository{}
	mockDBStorage.EXPECT().GetContractsRepository().Return(mockContractsRepo).Times(1)
	contractsRepo := mockDBStorage.GetContractsRepository()
	assert.Equal(t, mockContractsRepo, contractsRepo)

	// Test GetCreateAndSellRepository
	mockCreateAndSellRepo := &repository.CreateAndSellRepository{}
	mockDBStorage.EXPECT().GetCreateAndSellRepository().Return(mockCreateAndSellRepo).Times(1)
	createAndSellRepo := mockDBStorage.GetCreateAndSellRepository()
	assert.Equal(t, mockCreateAndSellRepo, createAndSellRepo)

	// Test GetFaqsRepository
	mockFaqsRepo := &repository.FaqsRepository{}
	mockDBStorage.EXPECT().GetFaqsRepository().Return(mockFaqsRepo).Times(1)
	faqsRepo := mockDBStorage.GetFaqsRepository()
	assert.Equal(t, mockFaqsRepo, faqsRepo)

	// Test GetFilesRepository
	mockFilesRepo := &repository.FilesRepository{}
	mockDBStorage.EXPECT().GetFilesRepository().Return(mockFilesRepo).Times(1)
	filesRepo := mockDBStorage.GetFilesRepository()
	assert.Equal(t, mockFilesRepo, filesRepo)

	// Test GetHelpCenterRepository
	mockHelpCenterRepo := &repository.HelpCenterRepository{}
	mockDBStorage.EXPECT().GetHelpCenterRepository().Return(mockHelpCenterRepo).Times(1)
	helpCenterRepo := mockDBStorage.GetHelpCenterRepository()
	assert.Equal(t, mockHelpCenterRepo, helpCenterRepo)

	// Test GetMenusRepository
	mockMenusRepo := &repository.MenusRepository{}
	mockDBStorage.EXPECT().GetMenusRepository().Return(mockMenusRepo).Times(1)
	menusRepo := mockDBStorage.GetMenusRepository()
	assert.Equal(t, mockMenusRepo, menusRepo)

	// Test GetPagesRepository
	mockPagesRepo := &repository.PagesRepository{}
	mockDBStorage.EXPECT().GetPagesRepository().Return(mockPagesRepo).Times(1)
	pagesRepo := mockDBStorage.GetPagesRepository()
	assert.Equal(t, mockPagesRepo, pagesRepo)

	// Test GetSlidersRepository
	mockSlidersRepo := &repository.SlidersRepository{}
	mockDBStorage.EXPECT().GetSlidersRepository().Return(mockSlidersRepo).Times(1)
	slidersRepo := mockDBStorage.GetSlidersRepository()
	assert.Equal(t, mockSlidersRepo, slidersRepo)

	// Test GetTagsRepository
	mockTagsRepo := &repository.TagsRepository{}
	mockDBStorage.EXPECT().GetTagsRepository().Return(mockTagsRepo).Times(1)
	tagsRepo := mockDBStorage.GetTagsRepository()
	assert.Equal(t, mockTagsRepo, tagsRepo)

	// Test GetUsersRepository
	mockUsersRepo := &repository.UsersRepository{}
	mockDBStorage.EXPECT().GetUsersRepository().Return(mockUsersRepo).Times(1)
	usersRepo := mockDBStorage.GetUsersRepository()
	assert.Equal(t, mockUsersRepo, usersRepo)

	// Test GetWalletsRepository
	mockWalletsRepo := &repository.WalletsRepository{}
	mockDBStorage.EXPECT().GetWalletsRepository().Return(mockWalletsRepo).Times(1)
	walletsRepo := mockDBStorage.GetWalletsRepository()
	assert.Equal(t, mockWalletsRepo, walletsRepo)
}
