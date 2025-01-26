package mocks_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	sqlx "github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
)

func TestMockWalletsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockWalletsRepositoryInterface(ctrl)

	ctx := context.Background()

	// Sample WalletConnect objects for testing
	sampleWallet := &models.WalletConnect{
		Created:     strfmt.Date(time.Now()),
		CreatedByID: 1,
		Description: "Test Wallet",
		Enabled:     true,
		ID:          1,
		ImageID:     101,
		OrderBy:     1,
		Title:       "Test Title",
		Updated:     strfmt.Date(time.Now()),
		UpdatedByID: 2,
	}

	t.Run("Create", func(t *testing.T) {
		walletMap := map[string]any{
			"title":       "New Wallet",
			"description": "Wallet description",
			"enabled":     true,
		}
		mockRepo.EXPECT().Create(ctx, walletMap).Return(nil)

		err := mockRepo.Create(ctx, walletMap)
		assert.NoError(t, err)
	})

	t.Run("GetPublicWalletConnect", func(t *testing.T) {
		expected := []*models.PublicWalletConnectItem{
			{ID: 1, Attributes: &models.PublicWalletConnectItemAttributes{Title: "Wallet1"}},
			{ID: 2, Attributes: &models.PublicWalletConnectItemAttributes{Title: "Wallet2"}},
		}
		mockRepo.EXPECT().GetPublicWalletConnect(ctx).Return(expected, nil)

		result, err := mockRepo.GetPublicWalletConnect(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("GetWalletByID", func(t *testing.T) {
		mockRepo.EXPECT().GetWalletByID(ctx, int64(1)).Return(sampleWallet, nil)

		result, err := mockRepo.GetWalletByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, sampleWallet, result)
	})

	t.Run("GetWalletByID with error", func(t *testing.T) {
		mockRepo.EXPECT().GetWalletByID(ctx, int64(999)).Return(nil, errors.New("wallet not found"))

		result, err := mockRepo.GetWalletByID(ctx, 999)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "wallet not found")
	})

	t.Run("GetWallets", func(t *testing.T) {
		expected := []*models.WalletConnect{
			sampleWallet,
		}
		mockRepo.EXPECT().GetWallets(ctx).Return(expected, nil)

		result, err := mockRepo.GetWallets(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("InnerDB", func(t *testing.T) {
		mockDB := &sqlx.DB{}
		mockRepo.EXPECT().InnerDB().Return(mockDB)

		result := mockRepo.InnerDB()
		assert.Equal(t, mockDB, result)
	})

	t.Run("Ping", func(t *testing.T) {
		mockRepo.EXPECT().Ping().Return(nil)

		err := mockRepo.Ping()
		assert.NoError(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		walletMap := map[string]any{
			"title":       "Updated Wallet",
			"description": "Updated description",
		}
		expected := []*models.WalletConnect{
			{
				Created:     strfmt.Date(time.Now()),
				CreatedByID: 1,
				Description: "Updated description",
				Enabled:     true,
				ID:          1,
				ImageID:     101,
				OrderBy:     1,
				Title:       "Updated Wallet",
				Updated:     strfmt.Date(time.Now()),
				UpdatedByID: 2,
			},
		}
		mockRepo.EXPECT().Update(ctx, walletMap).Return(expected, nil)

		result, err := mockRepo.Update(ctx, walletMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Update with error", func(t *testing.T) {
		walletMap := map[string]any{
			"title":       "Invalid Wallet",
			"description": "Invalid description",
		}
		mockRepo.EXPECT().Update(ctx, walletMap).Return(nil, errors.New("update failed"))

		result, err := mockRepo.Update(ctx, walletMap)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "update failed")
	})
}
