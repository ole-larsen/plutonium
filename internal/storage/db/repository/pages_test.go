package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	repo "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPagesRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "pages"
	repository := repo.NewPagesRepository(sqlxDB, tbl)

	assert.NotNil(t, repository)
	assert.Equal(t, tbl, repository.TBL)

	// Nil database case
	repositoryNil := repo.NewPagesRepository(nil, "pages")
	assert.Nil(t, repositoryNil)
}

func TestPagesRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewPagesRepository(sqlxDB, "pages")

	assert.Equal(t, sqlxDB, repository.InnerDB())

	// Nil receiver case
	var nilRepository *repo.PagesRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestPagesRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewPagesRepository(sqlxDB, "pages")

	mock.ExpectPing().WillReturnError(nil)

	err = repository.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repository.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repo.PagesRepository
	err = nilRepository.Ping()
	assert.Equal(t, repo.ErrDBNotInitialized, err)
}

func TestPagesRepository_PingEdgeCase(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewPagesRepository(sqlxDB, "pages")

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err = repository.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}

// TestPagesRepository_MigrateContext tests the MigrateContext method of PagesRepository.
func TestPagesRepository_MigrateContext(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repository := &repo.PagesRepository{
		DB:  *sqlxDB,
		TBL: "pages",
	}

	ctx := context.Background()

	// Test successful migration
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS pages`).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repository.MigrateContext(ctx)
	assert.NoError(t, err, "MigrateContext() should not return an error")

	// Test migration with an error
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS pages`).WillReturnError(errors.New("exec error"))

	err = repository.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() should return an error")
	assert.Equal(t, "exec error", err.Error(), "MigrateContext() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repo.PagesRepository
	err = nilRepo.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() on nil repository should return error")
}
