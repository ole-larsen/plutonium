package repository_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	repo "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCategoriesRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "categories"
	repository := repo.NewCategoriesRepository(sqlxDB, tbl)

	assert.NotNil(t, repository)
	assert.Equal(t, tbl, repository.TBL)

	// Nil database case
	repositoryNil := repo.NewCategoriesRepository(nil, "categories")
	assert.Nil(t, repositoryNil)
}

func TestCategoriesRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewCategoriesRepository(sqlxDB, "categories")

	assert.Equal(t, sqlxDB, repository.InnerDB())

	// Nil receiver case
	var nilRepository *repo.CategoriesRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestCategoriesRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewCategoriesRepository(sqlxDB, "categories")

	mock.ExpectPing().WillReturnError(nil)

	err = repository.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repository.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repo.CategoriesRepository
	err = nilRepository.Ping()
	assert.Equal(t, repo.ErrDBNotInitialized, err)
}

func TestCategoriesRepository_PingEdgeCase(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewCategoriesRepository(sqlxDB, "categories")

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err = repository.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}
