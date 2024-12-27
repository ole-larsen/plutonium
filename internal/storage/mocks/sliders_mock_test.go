package mocks_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"go.uber.org/mock/gomock"
)

func TestMockSlidersRepositoryInterface_InnerDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	mockDB := &sqlx.DB{}
	mockRepo.EXPECT().InnerDB().Return(mockDB).Times(1)

	db := mockRepo.InnerDB()
	if db != mockDB {
		t.Errorf("Expected InnerDB to return %v, got %v", mockDB, db)
	}
}

func TestMockSlidersRepositoryInterface_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)

	// Successful Ping
	mockRepo.EXPECT().Ping().Return(nil).Times(1)

	if err := mockRepo.Ping(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Ping with error
	mockError := fmt.Errorf("ping error")
	mockRepo.EXPECT().Ping().Return(mockError).Times(1)

	if err := mockRepo.Ping(); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockSlidersRepositoryInterface_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	ctx := context.Background()
	sliderMap := map[string]interface{}{"title": "test slider"}

	// Successful creation
	mockRepo.EXPECT().Create(ctx, sliderMap).Return(nil).Times(1)

	if err := mockRepo.Create(ctx, sliderMap); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Creation with error
	mockError := fmt.Errorf("creation error")
	mockRepo.EXPECT().Create(ctx, sliderMap).Return(mockError).Times(1)

	if err := mockRepo.Create(ctx, sliderMap); err == nil || err.Error() != mockError.Error() {
		t.Errorf("Expected error %v, got %v", mockError, err)
	}
}

func TestMockSlidersRepositoryInterface_GetSliderByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	ctx := context.Background()
	id := int64(1)
	expectedSlider := &models.Slider{ID: id, Title: "test slider"}

	// Successful retrieval
	mockRepo.EXPECT().GetSliderByID(ctx, id).Return(expectedSlider, nil).Times(1)

	slider, err := mockRepo.GetSliderByID(ctx, id)
	if err != nil || slider != expectedSlider {
		t.Errorf("Expected %v, got %v (error: %v)", expectedSlider, slider, err)
	}

	// Retrieval with error
	mockError := fmt.Errorf("slider not found")
	mockRepo.EXPECT().GetSliderByID(ctx, id).Return(nil, mockError).Times(1)

	slider, err = mockRepo.GetSliderByID(ctx, id)
	if err == nil || err.Error() != mockError.Error() || slider != nil {
		t.Errorf("Expected error %v and nil slider, got %v and %v", mockError, slider, err)
	}
}

func TestMockSlidersRepositoryInterface_GetSliderByProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	ctx := context.Background()
	provider := "test-provider"
	expectedPublicSlider := &models.PublicSlider{
		ID: 1,
		Attributes: &models.PublicSliderAttributes{
			SlidesItem: []*models.PublicSliderItem{},
		},
	}

	// Successful retrieval
	mockRepo.EXPECT().GetSliderByProvider(ctx, provider).Return(expectedPublicSlider, nil).Times(1)

	publicSlider, err := mockRepo.GetSliderByProvider(ctx, provider)
	if err != nil || publicSlider != expectedPublicSlider {
		t.Errorf("Expected %v, got %v (error: %v)", expectedPublicSlider, publicSlider, err)
	}

	// Retrieval with error
	mockError := fmt.Errorf("slider not found")
	mockRepo.EXPECT().GetSliderByProvider(ctx, provider).Return(nil, mockError).Times(1)

	publicSlider, err = mockRepo.GetSliderByProvider(ctx, provider)
	if err == nil || err.Error() != mockError.Error() || publicSlider != nil {
		t.Errorf("Expected error %v and nil slider, got %v and %v", mockError, publicSlider, err)
	}
}

func TestMockSlidersRepositoryInterface_GetSliderByTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	ctx := context.Background()
	title := "test-title"
	expectedSlider := &models.Slider{Title: title}

	// Successful retrieval
	mockRepo.EXPECT().GetSliderByTitle(ctx, title).Return(expectedSlider, nil).Times(1)

	slider, err := mockRepo.GetSliderByTitle(ctx, title)
	if err != nil || slider != expectedSlider {
		t.Errorf("Expected %v, got %v (error: %v)", expectedSlider, slider, err)
	}

	// Retrieval with error
	mockError := fmt.Errorf("slider not found")
	mockRepo.EXPECT().GetSliderByTitle(ctx, title).Return(nil, mockError).Times(1)

	slider, err = mockRepo.GetSliderByTitle(ctx, title)
	if err == nil || err.Error() != mockError.Error() || slider != nil {
		t.Errorf("Expected error %v and nil slider, got %v and %v", mockError, slider, err)
	}
}

func TestMockSlidersRepositoryInterface_GetSliders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	ctx := context.Background()
	expectedSliders := []*models.Slider{
		{ID: 1, Title: "slider1"},
		{ID: 2, Title: "slider2"},
	}

	// Successful retrieval
	mockRepo.EXPECT().GetSliders(ctx).Return(expectedSliders, nil).Times(1)

	sliders, err := mockRepo.GetSliders(ctx)
	if err != nil || len(sliders) != len(expectedSliders) {
		t.Errorf("Expected %v, got %v (error: %v)", expectedSliders, sliders, err)
	}

	// Retrieval with error
	mockError := fmt.Errorf("error retrieving sliders")
	mockRepo.EXPECT().GetSliders(ctx).Return(nil, mockError).Times(1)

	sliders, err = mockRepo.GetSliders(ctx)
	if err == nil || err.Error() != mockError.Error() || sliders != nil {
		t.Errorf("Expected error %v and nil sliders, got %v and %v", mockError, sliders, err)
	}
}

func TestMockSlidersRepositoryInterface_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSlidersRepositoryInterface(ctrl)
	ctx := context.Background()
	sliderMap := map[string]interface{}{"title": "updated slider"}
	expectedUpdatedSliders := []*models.Slider{
		{ID: 1, Title: "updated slider"},
	}

	// Successful update
	mockRepo.EXPECT().Update(ctx, sliderMap).Return(expectedUpdatedSliders, nil).Times(1)

	updatedSliders, err := mockRepo.Update(ctx, sliderMap)
	if err != nil || len(updatedSliders) != len(expectedUpdatedSliders) {
		t.Errorf("Expected %v, got %v (error: %v)", expectedUpdatedSliders, updatedSliders, err)
	}

	// Update with error
	mockError := fmt.Errorf("update error")
	mockRepo.EXPECT().Update(ctx, sliderMap).Return(nil, mockError).Times(1)

	updatedSliders, err = mockRepo.Update(ctx, sliderMap)
	if err == nil || err.Error() != mockError.Error() || updatedSliders != nil {
		t.Errorf("Expected error %v and nil sliders, got %v and %v", mockError, updatedSliders, err)
	}
}
