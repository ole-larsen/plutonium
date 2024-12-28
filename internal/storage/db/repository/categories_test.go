package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func setupTestCategoriesRepo(t *testing.T) (*repository.CategoriesRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewCategoriesRepository(sqlxDB, "categories")

	return repo, mock
}

func TestNewCategoriesRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "categories"
	repo := repository.NewCategoriesRepository(sqlxDB, tbl)

	assert.NotNil(t, repo)
	assert.Equal(t, tbl, repo.TBL)

	// Nil database case
	repositoryNil := repository.NewCategoriesRepository(nil, "categories")
	assert.Nil(t, repositoryNil)
}

func TestCategoriesRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewCategoriesRepository(sqlxDB, "categories")

	assert.Equal(t, sqlxDB, repo.InnerDB())

	// Nil receiver case
	var nilRepository *repository.CategoriesRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestCategoriesRepository_Ping(t *testing.T) {
	repo, mock := setupTestCategoriesRepo(t)
	mock.ExpectPing().WillReturnError(nil)

	err := repo.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repo.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repository.CategoriesRepository
	err = nilRepository.Ping()
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestCategoriesRepository_PingEdgeCase(t *testing.T) {
	repo, mock := setupTestCategoriesRepo(t)

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err := repo.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}

func TestCategoriesRepository_GetPublicCollectibleCategories_NilReceiver(t *testing.T) {
	var nilRepository *repository.CategoriesRepository

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersRepo := repository.NewUsersRepository(nil, "users")

	categories, err := nilRepository.GetPublicCollectibleCategories(context.TODO(), usersRepo)
	assert.Nil(t, categories)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestCategoriesRepository_GetPublicCollectibleCategories_EmptyResultSet(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := repository.NewCategoriesRepository(sqlxDB, "categories")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersRepo := mocks.NewMockUsersRepositoryInterface(ctrl)

	mock.ExpectQuery("SELECT .* FROM categories").
		WillReturnRows(sqlmock.NewRows(nil))

	categories, err := repo.GetPublicCollectibleCategories(context.TODO(), usersRepo)
	assert.NoError(t, err)
	assert.Empty(t, categories)
}
