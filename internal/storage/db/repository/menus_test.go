package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	repo "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMenusRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "menus"
	repository := repo.NewMenusRepository(sqlxDB, tbl)

	assert.NotNil(t, repository)
	assert.Equal(t, tbl, repository.TBL)

	// Nil database case
	repositoryNil := repo.NewMenusRepository(nil, "menus")
	assert.Nil(t, repositoryNil)
}

func TestMenusRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewMenusRepository(sqlxDB, "menus")

	assert.Equal(t, sqlxDB, repository.InnerDB())

	// Nil receiver case
	var nilRepository *repo.MenusRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestMenusRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewMenusRepository(sqlxDB, "menus")

	mock.ExpectPing().WillReturnError(nil)

	err = repository.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repository.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repo.MenusRepository
	err = nilRepository.Ping()
	assert.Equal(t, repo.ErrDBNotInitialized, err)
}

func TestMenusRepository_PingEdgeCase(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewMenusRepository(sqlxDB, "menus")

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err = repository.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}

// TestMenusRepository_MigrateContext tests the MigrateContext method of MenusRepository.
func TestMenusRepository_MigrateContext(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repository := &repo.MenusRepository{
		DB:  *sqlxDB,
		TBL: "menus",
	}

	ctx := context.Background()

	// Test successful migration
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS menus`).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repository.MigrateContext(ctx)
	assert.NoError(t, err, "MigrateContext() should not return an error")

	// Test migration with an error
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS menus`).WillReturnError(errors.New("exec error"))

	err = repository.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() should return an error")
	assert.Equal(t, "exec error", err.Error(), "MigrateContext() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repo.MenusRepository
	err = nilRepo.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() on nil repository should return error")
}

func TestMenusRepository_GetMenuByProvider(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewMenusRepository(sqlxDB, "menus")

	provider := "test-provider"
	expectedMenu := &models.PublicMenu{
		ID: 1,
		Attributes: &models.PublicMenuAttributes{
			Name: "Test Menu",
			Link: "/test",
			Items: []*models.PublicMenu{
				{
					ID: 2,
					Attributes: &models.PublicMenuAttributes{
						Name: "Test Item",
						Link: "/test-1",
					},
				},
			},
		},
	}

	// Success case
	mock.ExpectQuery(`SELECT m1\.id, JSON_BUILD_OBJECT`).
		WithArgs(provider).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "attributes"}).
				AddRow(
					expectedMenu.ID,
					`{"name":"Test Menu","link":"/test","items":[{"name":"Test Item","link":"/test-1"}]}`,
				),
		)

	menu, err := repository.GetMenuByProvider(provider)
	assert.NoError(t, err)
	assert.NotNil(t, menu)
	assert.Equal(t, expectedMenu.ID, menu.ID)
	assert.Equal(t, expectedMenu.Attributes.Name, menu.Attributes.Name)

	// No rows case
	mock.ExpectQuery(`SELECT m1\.id, JSON_BUILD_OBJECT`).
		WithArgs(provider).
		WillReturnError(sql.ErrNoRows)

	menu, err = repository.GetMenuByProvider(provider)
	assert.Error(t, err, "[repository]: menu not found")
	assert.Nil(t, menu)

	// Query error case
	mock.ExpectQuery(`SELECT m1\.id, JSON_BUILD_OBJECT`).
		WithArgs(provider).
		WillReturnError(errors.New("query error"))

	menu, err = repository.GetMenuByProvider(provider)
	assert.Error(t, err)
	assert.Equal(t, "query error", err.Error())
	require.Nil(t, menu)

	// Nil receiver case
	var nilRepo *repo.MenusRepository
	menu, err = nilRepo.GetMenuByProvider(provider)
	assert.Nil(t, menu)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
}
