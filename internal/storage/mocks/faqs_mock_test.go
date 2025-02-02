package mocks_test

import (
	"context"
	"testing"

	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMockFaqsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFaqsRepositoryInterface(ctrl)
	ctx := context.Background()

	// Example data
	faqID := int64(1)
	faq := &models.Faq{ID: faqID, Question: "What is Plutonium?", Answer: "It's a digital marketplace."}
	faqs := []*models.Faq{faq}
	faqMap := map[string]interface{}{"id": faqID, "question": "Updated Question", "answer": "Updated Answer"}
	publicFaqItem := &models.PublicFaqItem{ID: faqID, Attributes: &models.PublicFaqItemAttributes{Question: "What is Plutonium?"}}
	publicFaqs := []*models.PublicFaqItem{publicFaqItem}

	// Test Create
	mockRepo.EXPECT().Create(ctx, faqMap).Return(nil).Times(1)
	err := mockRepo.Create(ctx, faqMap)
	assert.NoError(t, err, "Create method should not return an error")

	// Test GetFaqByID
	mockRepo.EXPECT().GetFaqByID(ctx, faqID).Return(faq, nil).Times(1)
	returnedFaq, err := mockRepo.GetFaqByID(ctx, faqID)
	assert.NoError(t, err, "GetFaqByID method should not return an error")
	assert.Equal(t, faq, returnedFaq, "GetFaqByID should return the expected FAQ")

	// Test GetFaqs
	mockRepo.EXPECT().GetFaqs(ctx).Return(faqs, nil).Times(1)
	returnedFaqs, err := mockRepo.GetFaqs(ctx)
	assert.NoError(t, err, "GetFaqs method should not return an error")
	assert.Equal(t, faqs, returnedFaqs, "GetFaqs should return the expected FAQs")

	// Test GetPublicFaqs
	mockRepo.EXPECT().GetPublicFaqs(ctx).Return(publicFaqs, nil).Times(1)
	returnedPublicFaqs, err := mockRepo.GetPublicFaqs(ctx)
	assert.NoError(t, err, "GetPublicFaqs method should not return an error")
	assert.Equal(t, publicFaqs, returnedPublicFaqs, "GetPublicFaqs should return the expected public FAQs")

	// Test InnerDB
	mockRepo.EXPECT().InnerDB().Return(nil).Times(1)
	db := mockRepo.InnerDB()
	assert.Nil(t, db, "InnerDB should return nil for this mock test")

	// Test Ping
	mockRepo.EXPECT().Ping().Return(nil).Times(1)
	err = mockRepo.Ping()
	assert.NoError(t, err, "Ping method should not return an error")

	// Test Update
	mockRepo.EXPECT().Update(ctx, faqMap).Return(faqs, nil).Times(1)
	updatedFaqs, err := mockRepo.Update(ctx, faqMap)
	assert.NoError(t, err, "Update method should not return an error")
	assert.Equal(t, faqs, updatedFaqs, "Update should return the updated FAQs")
}
