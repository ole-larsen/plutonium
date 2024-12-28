package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFilesRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "files"
	repo := repository.NewFilesRepository(sqlxDB, tbl)

	assert.NotNil(t, repo)
	assert.Equal(t, tbl, repo.TBL)

	// Nil database case
	repositoryNil := repository.NewFilesRepository(nil, "files")
	assert.Nil(t, repositoryNil)
}

func TestFilesRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewFilesRepository(sqlxDB, "files")

	assert.Equal(t, sqlxDB, repo.InnerDB())

	// Nil receiver case
	var nilRepository *repository.FilesRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestFilesRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewFilesRepository(sqlxDB, "files")

	mock.ExpectPing().WillReturnError(nil)

	err = repo.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repo.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	err = nilRepository.Ping()
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestFilesRepository_PingEdgeCase(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewFilesRepository(sqlxDB, "files")

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err = repo.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}

func setupTestFilesRepo(t *testing.T) (*repository.FilesRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewFilesRepository(sqlxDB, "files")

	return repo, mock
}

func TestFilesRepository_Create(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	fileMap := map[string]interface{}{
		"name":          "file1",
		"alt":           "alt text",
		"caption":       "caption",
		"hash":          "hash1",
		"mime":          "image/png",
		"ext":           ".png",
		"size":          1024,
		"width":         100,
		"height":        200,
		"provider":      "local",
		"url":           "/url",
		"created_by_id": 1,
		"updated_by_id": 1,
	}

	mock.ExpectExec(`
INSERT INTO files`).WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(ctx, fileMap)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	err = nilRepository.Create(ctx, fileMap)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestFilesRepository_Update(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	fileMap := map[string]interface{}{
		"id":       int64(1),
		"name":     "test-file",
		"alt":      "test-alt",
		"hash":     "test-hash",
		"caption":  "test-caption",
		"ext":      ".png",
		"mime":     "image/png",
		"size":     1024.0,
		"width":    100,
		"height":   100,
		"provider": "test-provider",
	}

	expectedURL := "/api/v1/files/test-file.png"
	fileMap["url"] = expectedURL

	// Mock NamedExecContext for the update query
	mock.ExpectExec("UPDATE files SET").
		WithArgs(
			fileMap["name"], fileMap["alt"], fileMap["hash"], fileMap["caption"], fileMap["ext"],
			fileMap["mime"], fileMap["size"], fileMap["width"], fileMap["height"], fileMap["url"], fileMap["provider"], fileMap["id"],
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock GetFiles
	mockRows := sqlmock.NewRows([]string{
		"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url", "created_by_id", "updated_by_id",
	}).
		AddRow(
			int64(1), "test-file", "test-alt", "test-caption", "test-hash", "image/png", ".png",
			1024.0, 100, 100, "test-provider", expectedURL, int64(1), int64(1),
		)
	mock.ExpectQuery("SELECT id, name, alt, caption, hash, mime, ext, size, width, height, provider, url, created_by_id, updated_by_id from files").
		WillReturnRows(mockRows)

	// Call the method
	files, err := repo.Update(ctx, fileMap)

	// Assertions
	require.NoError(t, err)
	require.Len(t, files, 1)
	assert.Equal(t, int64(1), files[0].ID)
	assert.Equal(t, "test-file.png", files[0].Name)
	assert.Equal(t, "test-alt", files[0].Alt)
	assert.Equal(t, "test-caption", files[0].Caption)
	assert.Equal(t, "test-hash", files[0].Hash)
	assert.Equal(t, "image/png", files[0].Type)
	assert.Equal(t, ".png", files[0].Ext)
	assert.Equal(t, float64(1024), files[0].Size)
	assert.Equal(t, int64(100), files[0].Width)
	assert.Equal(t, int64(100), files[0].Height)
	assert.Equal(t, "test-provider", files[0].Provider)
	assert.Equal(t, expectedURL, files[0].Thumb)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	mock.ExpectQuery("SELECT id, name, alt, caption, hash, mime, ext, size, width, height, provider, url, created_by_id, updated_by_id from files").
		WillReturnError(errors.New("failed"))

	// Call the method
	files, err = repo.Update(ctx, fileMap)
	assert.Error(t, err)
	assert.Nil(t, files)

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	file, err := nilRepository.Update(ctx, fileMap)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, file)
}

func TestFilesRepository_GetFiles(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	rows := sqlmock.NewRows([]string{"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url", "created_by_id", "updated_by_id"}).
		AddRow(1, "file1", "alt text", "caption", "hash1", "image/png", ".png", 1024, 100, 200, "local", "/url", 1, 2)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WillReturnRows(rows)

	files, err := repo.GetFiles(ctx)
	assert.NoError(t, err)
	assert.Len(t, files, 1)
	assert.Equal(t, "file1.png", files[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WillReturnError(errors.New("failed to get files"))

	files, err = repo.GetFiles(ctx)
	assert.Error(t, err)
	assert.Nil(t, files)

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	files, err = nilRepository.GetFiles(ctx)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, files)

	// Mock invalid rows to trigger rows.Scan error
	invalidRows := sqlmock.NewRows([]string{"id", "name", "alt"}). // Missing required columns
									AddRow(1, "file1", "alt text")

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WillReturnRows(invalidRows)

	files, err = repo.GetFiles(ctx)
	assert.Error(t, err) // Expect an error due to scan failure
	assert.Nil(t, files)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFilesRepository_GetFileByName(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	rows := sqlmock.NewRows([]string{"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url", "created_by_id", "updated_by_id"}).
		AddRow(1, "file1", "alt text", "caption", "hash1", "image/png", ".png", 1024, 100, 200, "local", "/url", 1, 2)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("file1").WillReturnRows(rows)

	file, err := repo.GetFileByName(ctx, "file1")
	assert.NoError(t, err)
	assert.Equal(t, "file1.png", file.Name)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	file, err = nilRepository.GetFileByName(ctx, "file1")
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, file)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("file1").WillReturnError(sql.ErrNoRows)

	file, err = repo.GetFileByName(ctx, "file1")
	assert.Error(t, err)
	assert.Equal(t, "file not found", err.Error())
	assert.Nil(t, file)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("file1").WillReturnError(errors.New("some error"))

	file, err = repo.GetFileByName(ctx, "file1")
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
	assert.Nil(t, file)
}

func TestFilesRepository_GetFileByID(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	rows := sqlmock.NewRows([]string{"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url", "created_by_id", "updated_by_id"}).
		AddRow(1, "file1", "alt text", "caption", "hash1", "image/png", ".png", 1024, 100, 200, "local", "/url", 1, 2)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs(1).WillReturnRows(rows)

	file, err := repo.GetFileByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "file1", file.Name)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	file, err = nilRepository.GetFileByID(ctx, 1)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, file)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs(1).WillReturnError(sql.ErrNoRows)

	file, err = repo.GetFileByID(ctx, 1)
	assert.Error(t, err)
	assert.Equal(t, "file not found", err.Error())
	assert.Nil(t, file)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs(1).WillReturnError(errors.New("some error"))

	file, err = repo.GetFileByID(ctx, 1)
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
	assert.Nil(t, file)
}

func TestFilesRepository_GetPublicFilesByProvider(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	rows := sqlmock.NewRows([]string{"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url"}).
		AddRow(1, "file1", "alt text", "caption", "hash1", "image/png", ".png", 1024, 100, 200, "local", "/url")

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("local").WillReturnRows(rows)

	files, err := repo.GetPublicFilesByProvider(ctx, "local")
	assert.NoError(t, err)
	assert.Len(t, files, 1)
	assert.Equal(t, "file1", files[0].Attributes.Name)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	files, err = nilRepository.GetPublicFilesByProvider(ctx, "local")
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, files)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("local").WillReturnError(sql.ErrNoRows)

	files, err = repo.GetPublicFilesByProvider(ctx, "local")
	assert.Error(t, err)
	assert.Equal(t, "sql: no rows in result set", err.Error())
	assert.Nil(t, files)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("local").WillReturnError(errors.New("some error"))

	files, err = repo.GetPublicFilesByProvider(ctx, "local")
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
	assert.Nil(t, files)

	// Test general query error
	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("local").WillReturnError(errors.New("some error"))

	files, err = repo.GetPublicFilesByProvider(ctx, "local")
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
	assert.Nil(t, files)

	// Test scan error due to invalid rows
	invalidRows := sqlmock.NewRows([]string{"id", "name", "alt"}). // Missing required columns
									AddRow(1, "file1", "alt text")

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("local").WillReturnRows(invalidRows)

	files, err = repo.GetPublicFilesByProvider(ctx, "local")
	assert.Error(t, err) // Expect a scan error
	assert.Nil(t, files)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFilesRepository_GetPublicFileByName(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	rows := sqlmock.NewRows([]string{"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url", "created_by_id", "updated_by_id"}).
		AddRow(1, "file1", "alt text", "caption", "hash1", "image/png", ".png", 1024, 100, 200, "local", "/url", 1, 2)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs("file1").WillReturnRows(rows)

	// Happy path test
	publicFile, err := repo.GetPublicFileByName(ctx, "file1")
	assert.NoError(t, err)
	assert.NotNil(t, publicFile)
	assert.Equal(t, "file1.png", publicFile.Attributes.Name)

	// Simulate error from GetFileByName
	mock.ExpectQuery(`SELECT id`).WithArgs("nonexistent").WillReturnError(errors.New("file not found"))

	publicFile, err = repo.GetPublicFileByName(ctx, "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, publicFile)

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	publicFile, err = nilRepository.GetPublicFileByName(ctx, "file1")
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, publicFile)
}

func TestFilesRepository_GetPublicFileByID(t *testing.T) {
	repo, mock := setupTestFilesRepo(t)
	ctx := context.TODO()

	rows := sqlmock.NewRows([]string{"id", "name", "alt", "caption", "hash", "mime", "ext", "size", "width", "height", "provider", "url", "created_by_id", "updated_by_id"}).
		AddRow(1, "file1", "alt text", "caption", "hash1", "image/png", ".png", 1024, 100, 200, "local", "/url", 1, 2)

	mock.ExpectQuery(`SELECT id, name, alt, caption`).WithArgs(1).WillReturnRows(rows)

	// Happy path test
	publicFile, err := repo.GetPublicFileByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, publicFile)
	assert.Equal(t, "file1", publicFile.Attributes.Name)

	// Simulate error from GetFileByID
	mock.ExpectQuery(`SELECT .* FROM files WHERE id = .*`).WithArgs(999).WillReturnError(errors.New("file not found"))

	publicFile, err = repo.GetPublicFileByID(ctx, 999)
	assert.Error(t, err)
	assert.Nil(t, publicFile)

	// Nil receiver case
	var nilRepository *repository.FilesRepository
	publicFile, err = nilRepository.GetPublicFileByID(ctx, 1)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, publicFile)
}
