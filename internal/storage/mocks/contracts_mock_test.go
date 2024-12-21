package mocks_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	sqlx "github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMockContractsRepositoryInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockContractsRepositoryInterface(ctrl)
	ctx := context.Background()

	// Test: Create
	t.Run("Create", func(t *testing.T) {
		contractMap := map[string]interface{}{"name": "test_contract"}
		mockRepo.EXPECT().Create(ctx, contractMap).Return(nil)

		err := mockRepo.Create(ctx, contractMap)
		assert.NoError(t, err)
	})

	// Test: GetAuctions
	t.Run("GetAuctions", func(t *testing.T) {
		expectedContracts := []*models.Contract{
			{Name: "Auction1"},
			{Name: "Auction2"},
		}
		mockRepo.EXPECT().GetAuctions(ctx).Return(expectedContracts, nil)

		contracts, err := mockRepo.GetAuctions(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedContracts, contracts)
	})

	// Test: GetByAddress
	t.Run("GetByAddress", func(t *testing.T) {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		expectedContract := &models.Contract{Name: "Contract1"}
		mockRepo.EXPECT().GetByAddress(ctx, address).Return(expectedContract, nil)

		contract, err := mockRepo.GetByAddress(ctx, address)
		assert.NoError(t, err)
		assert.Equal(t, expectedContract, contract)
	})

	// Test: GetCollectionsContracts
	t.Run("GetCollectionsContracts", func(t *testing.T) {
		expectedContracts := []*models.Contract{
			{Name: "Collection1"},
			{Name: "Collection2"},
		}
		mockRepo.EXPECT().GetCollectionsContracts(ctx).Return(expectedContracts, nil)

		contracts, err := mockRepo.GetCollectionsContracts(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedContracts, contracts)
	})

	// Test: GetOne
	t.Run("GetOne", func(t *testing.T) {
		name := "ContractName"
		expectedContract := &models.Contract{Name: name}
		mockRepo.EXPECT().GetOne(ctx, name).Return(expectedContract, nil)

		contract, err := mockRepo.GetOne(ctx, name)
		assert.NoError(t, err)
		assert.Equal(t, expectedContract, contract)
	})

	// Test: InnerDB
	t.Run("InnerDB", func(t *testing.T) {
		expectedDB := &sqlx.DB{}
		mockRepo.EXPECT().InnerDB().Return(expectedDB)

		db := mockRepo.InnerDB()
		assert.Equal(t, expectedDB, db)
	})

	// Test: MigrateContext
	t.Run("MigrateContext", func(t *testing.T) {
		mockRepo.EXPECT().MigrateContext(ctx).Return(nil)

		err := mockRepo.MigrateContext(ctx)
		assert.NoError(t, err)
	})

	// Test: Ping
	t.Run("Ping", func(t *testing.T) {
		mockRepo.EXPECT().Ping().Return(nil)

		err := mockRepo.Ping()
		assert.NoError(t, err)
	})
}
