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

func TestMockContactsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockContactsRepositoryInterface(ctrl)

	// Test Create method
	ctx := context.Background()
	contactMap := map[string]any{"key": "value"}
	mockRepo.EXPECT().Create(ctx, contactMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, contactMap)
	assert.NoError(t, err)

	// Test GetContactByID method
	id := int64(123)
	mockContact := &models.Contact{}
	mockRepo.EXPECT().GetContactByID(ctx, id).Return(mockContact, nil).Times(1)
	contact, err := mockRepo.GetContactByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, mockContact, contact)

	// Test GetContactByPageID method
	pageID := int64(456)
	mockPublicContact := &models.PublicContact{}
	mockRepo.EXPECT().GetContactByPageID(ctx, pageID).Return(mockPublicContact, nil).Times(1)
	publicContact, err := mockRepo.GetContactByPageID(ctx, pageID)
	assert.NoError(t, err)
	assert.Equal(t, mockPublicContact, publicContact)

	// Test GetContacts method
	mockContacts := []*models.Contact{{}, {}}
	mockRepo.EXPECT().GetContacts(ctx).Return(mockContacts, nil).Times(1)
	contacts, err := mockRepo.GetContacts(ctx)
	assert.NoError(t, err)
	assert.Equal(t, mockContacts, contacts)

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
	mockRepo.EXPECT().Update(ctx, updateMap).Return(mockContacts, nil).Times(1)
	updatedContacts, err := mockRepo.Update(ctx, updateMap)
	assert.NoError(t, err)
	assert.Equal(t, mockContacts, updatedContacts)
}
