package mocks_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMockPagesRepositoryInterface_InnerDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPagesRepositoryInterface(ctrl)
	db := &sqlx.DB{}

	// Expectation
	mockRepo.EXPECT().InnerDB().Return(db).Times(1)

	// Call the method
	result := mockRepo.InnerDB()

	// Assert
	assert.Equal(t, db, result, "InnerDB should return the expected database instance")
}

func TestMockPagesRepositoryInterface_MigrateContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPagesRepositoryInterface(ctrl)
	ctx := context.Background()

	// Test success case
	mockRepo.EXPECT().MigrateContext(ctx).Return(nil).Times(1)

	err := mockRepo.MigrateContext(ctx)
	assert.NoError(t, err, "MigrateContext should not return an error in success case")

	// Test failure case
	mockRepo.EXPECT().MigrateContext(ctx).Return(errors.New("migration failed")).Times(1)

	err = mockRepo.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext should return an error in failure case")
	assert.Equal(t, "migration failed", err.Error())
}

func TestMockPagesRepositoryInterface_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPagesRepositoryInterface(ctrl)

	// Test success case
	mockRepo.EXPECT().Ping().Return(nil).Times(1)

	err := mockRepo.Ping()
	assert.NoError(t, err, "Ping should not return an error in success case")

	// Test failure case
	mockRepo.EXPECT().Ping().Return(errors.New("ping failed")).Times(1)

	err = mockRepo.Ping()
	assert.Error(t, err, "Ping should return an error in failure case")
	assert.Equal(t, "ping failed", err.Error())
}
