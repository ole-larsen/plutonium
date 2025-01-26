package mocks_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockContactFormRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockContactFormRepositoryInterface(ctrl)

	// Test Create method
	ctx := context.Background()
	contactFormMap := map[string]any{"name": "John Doe", "email": "john.doe@example.com"}
	mockRepo.EXPECT().Create(ctx, contactFormMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, contactFormMap)
	assert.NoError(t, err)

	// Test InnerDB method
	mockDB := &sqlx.DB{}
	mockRepo.EXPECT().InnerDB().Return(mockDB).Times(1)
	db := mockRepo.InnerDB()
	assert.Equal(t, mockDB, db)

	// Test Ping method
	mockRepo.EXPECT().Ping().Return(nil).Times(1)
	err = mockRepo.Ping()
	assert.NoError(t, err)
}
