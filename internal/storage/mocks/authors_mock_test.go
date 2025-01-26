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

func TestMockAuthorsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthorsRepositoryInterface(ctrl)

	// Test BindSocial method
	socialMap := map[string]any{"platform": "Twitter", "handle": "user_handle"}
	mockRepo.EXPECT().BindSocial(context.Background(), socialMap).Return(nil).Times(1)
	err := mockRepo.BindSocial(context.Background(), socialMap)
	assert.NoError(t, err)

	// Test BindWallet method
	walletMap := map[string]any{"address": "0x1234567890abcdef"}
	mockRepo.EXPECT().BindWallet(context.Background(), walletMap).Return(nil).Times(1)
	err = mockRepo.BindWallet(context.Background(), walletMap)
	assert.NoError(t, err)

	// Test Create method
	authorMap := map[string]any{"name": "John Doe"}
	socials := []*models.Social{{Name: "Twitter", Link: "john_doe"}}
	wallets := []*models.Wallet{{Address: "0x1234567890abcdef"}}
	mockRepo.EXPECT().Create(context.Background(), authorMap, socials, wallets).Return(nil).Times(1)
	err = mockRepo.Create(context.Background(), authorMap, socials, wallets)
	assert.NoError(t, err)

	// Test GetAuthorByID method
	mockAuthor := &models.Author{ID: 1, Name: "John Doe"}
	mockRepo.EXPECT().GetAuthorByID(context.Background(), int64(1)).Return(mockAuthor, nil).Times(1)
	author, err := mockRepo.GetAuthorByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, mockAuthor, author)

	// Test GetAuthors method
	mockAuthors := []*models.Author{{ID: 1, Name: "John Doe"}}
	mockRepo.EXPECT().GetAuthors(context.Background()).Return(mockAuthors, nil).Times(1)
	authors, err := mockRepo.GetAuthors(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, mockAuthors, authors)

	// Test GetPublicAuthor method
	mockPublicAuthor := &models.PublicAuthorItem{ID: 1, Name: "John Doe"}
	mockRepo.EXPECT().GetPublicAuthor(context.Background(), "john-doe").Return(mockPublicAuthor, nil).Times(1)
	publicAuthor, err := mockRepo.GetPublicAuthor(context.Background(), "john-doe")
	assert.NoError(t, err)
	assert.Equal(t, mockPublicAuthor, publicAuthor)

	// Test GetPublicAuthors method
	mockPublicAuthors := []*models.PublicAuthorItem{{ID: 1, Name: "John Doe"}}
	mockRepo.EXPECT().GetPublicAuthors(context.Background()).Return(mockPublicAuthors, nil).Times(1)
	publicAuthors, err := mockRepo.GetPublicAuthors(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, mockPublicAuthors, publicAuthors)

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
	updatedAuthors := []*models.Author{{ID: 1, Name: "Updated John Doe"}}
	mockRepo.EXPECT().Update(context.Background(), authorMap, socials, wallets).Return(updatedAuthors, nil).Times(1)
	authors, err = mockRepo.Update(context.Background(), authorMap, socials, wallets)
	assert.NoError(t, err)
	assert.Equal(t, updatedAuthors, authors)
}
