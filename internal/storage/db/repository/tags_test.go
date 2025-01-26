package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestConvertSlice(t *testing.T) {
	t.Run("Valid Conversion", func(t *testing.T) {
		in := []any{1, 2, 3}
		out := repository.ConvertSlice[int](in)

		// Assert the result is of the expected type and values
		assert.Equal(t, []int{1, 2, 3}, out)
	})

	t.Run("Partial Conversion", func(t *testing.T) {
		in := []any{1, "string", 2.5}
		out := repository.ConvertSlice[int](in)

		// Assert that only the values of the correct type (int) are included
		assert.Equal(t, []int{1}, out)
	})

	t.Run("Empty Input", func(t *testing.T) {
		in := []any{}
		out := repository.ConvertSlice[int](in)

		// Assert that the result is an empty slice
		assert.Empty(t, out)
	})

	t.Run("Different Type Conversion", func(t *testing.T) {
		in := []any{"1", "2", "3"}
		out := repository.ConvertSlice[string](in)

		// Assert that all values are converted correctly
		assert.Equal(t, []string{"1", "2", "3"}, out)
	})

	t.Run("No Matching Type", func(t *testing.T) {
		in := []any{1, 2, 3}
		out := repository.ConvertSlice[string](in)

		// Assert that the result is an empty slice since no values are of type string
		assert.Empty(t, out)
	})
}

func TestConvertInterfaceToSlice(t *testing.T) {
	t.Run("Valid Slice Input", func(t *testing.T) {
		in := []int{1, 2, 3}
		out := repository.ConvertInterfaceToSlice(in)

		// Assert that it returns a slice of interface{}
		assert.Equal(t, []interface{}{1, 2, 3}, out)
	})

	t.Run("Empty Slice", func(t *testing.T) {
		in := []string{}
		out := repository.ConvertInterfaceToSlice(in)

		// Assert that it returns an empty slice
		assert.Empty(t, out)
	})

	t.Run("Non-Slice Input", func(t *testing.T) {
		in := 42
		out := repository.ConvertInterfaceToSlice(in)

		// Assert that it returns an empty slice for non-slice input
		assert.Empty(t, out)
	})

	t.Run("Different Type Slice", func(t *testing.T) {
		in := []float64{1.1, 2.2, 3.3}
		out := repository.ConvertInterfaceToSlice(in)

		// Assert that it returns a slice of interface{}
		assert.Equal(t, []interface{}{1.1, 2.2, 3.3}, out)
	})
}

// TestNewUsersRepository tests the constructor for UsersRepository.
func TestNewTagsRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	tbl := "tags"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlogs := mocks.NewMockBlogsRepositoryInterface(ctrl)

	mockPages := mocks.NewMockPagesRepositoryInterface(ctrl)

	repo := repository.NewTagsRepository(sqlxDB, "tags", mockBlogs, mockPages)

	// Check if db is correctly set
	assert.Equal(t, *sqlxDB, repo.DB, "NewTagsRepository() db mismatch")
	// Check if table name is correctly set
	assert.Equal(t, tbl, repo.TBL, "NewTagsRepository() tbl mismatch")

	repo = repository.NewTagsRepository(nil, "tags", mockBlogs, mockPages)

	assert.Nil(t, repo, "NewTagsRepository() should return nil when db is nil")
}

// TestTagsRepository_InnerDB tests the InnerDB method.
func TestTagsRepository_InnerDB(t *testing.T) {
	db, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.TagsRepository{
		DB: *sqlxDB,
	}

	// Test case: repo is not nil
	assert.Equal(t, sqlxDB, repo.InnerDB(), "InnerDB() should return the correct *sqlx.DB")

	// Test case: repo is nil
	var nilRepo *repository.TagsRepository

	assert.Nil(t, nilRepo.InnerDB(), "InnerDB() on nil repository should return nil")
}

// TestUsersRepository_Ping tests the Ping method.
func TestTagsRepository_Ping(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := &repository.TagsRepository{
		DB: *sqlxDB,
	}

	// Test successful ping
	mock.ExpectPing().WillReturnError(nil)

	err := repo.Ping()
	assert.NoError(t, err, "Ping() should not return an error")

	// Test ping with an error
	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repo.Ping()
	assert.Error(t, err, "Ping() should return an error")
	assert.Equal(t, "ping error", err.Error(), "Ping() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repository.TagsRepository
	err = nilRepo.Ping()
	assert.Error(t, err, "Ping() on nil repository should return an error")
	assert.Equal(t, repository.ErrDBNotInitialized, err, "Ping() should return ErrDBNotInitialized")
}

// TestTagsRepository_Create tests the Create method of TagsRepository.
func TestTagsRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := repository.NewTagsRepository(sqlxDB, "tags", nil, nil)

	tagMap := map[string]interface{}{
		"parent_id":     0,
		"title":         "test-tag",
		"slug":          "test-tag",
		"enabled":       true,
		"created_by_id": 1,
		"updated_by_id": 1,
	}

	// Mock INSERT INTO tags
	mock.ExpectExec("INSERT INTO tags").WithArgs(
		sqlmock.AnyArg(), // parent_id
		"test-tag",       // title
		"test-tag",       // slug
		true,             // enabled
		1,                // created_by_id
		1,                // updated_by_id
	).WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock SELECT id
	mock.ExpectQuery("SELECT id FROM tags WHERE title=").WithArgs("test-tag").WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1),
	)

	err := repo.Create(context.Background(), tagMap)
	assert.NoError(t, err, "Create() should not return an error")
	assert.Equal(t, int64(1), tagMap["id"], "Create() should set the correct ID in tagMap")

	// Test error handling during creation
	mock.ExpectExec("INSERT INTO tags").WillReturnError(errors.New("insert error"))

	err = repo.Create(context.Background(), tagMap)
	assert.Error(t, err, "Create() should return an error if INSERT fails")

	// Nil receiver case
	var nilRepository *repository.TagsRepository
	err = nilRepository.Create(context.Background(), tagMap)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestTagsRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := repository.NewTagsRepository(sqlxDB, "tags", nil, nil)
	tagMap := map[string]interface{}{
		"id":            int64(1),
		"title":         "updated-tag",
		"slug":          "updated-tag",
		"enabled":       true,
		"updated_by_id": 2,
		"parent_id":     0,
	}

	// Mock UPDATE query
	mock.ExpectExec("UPDATE tags SET").WithArgs(
		0,             // parent_id
		"updated-tag", // title
		"updated-tag", // slug
		true,          // enabled
		2,             // updated_by_id
		int64(1),      // id
	).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery("SELECT").WillReturnRows(
		sqlmock.NewRows([]string{
			"id", "parent_id", "title", "slug", "enabled",
			"created_by_id", "updated_by_id", "blog_id", "page_id",
		}).AddRow(
			int64(1), 0, "updated-tag", "updated-tag", true,
			2, 1, "{1,2}", "{3,4}",
		),
	)

	_, err := repo.Update(context.Background(), tagMap)
	assert.NoError(t, err, "Update() should not return an error")

	// Test error handling during update
	mock.ExpectExec("UPDATE tags SET").WillReturnError(errors.New("update error"))

	_, err = repo.Update(context.Background(), tagMap)
	assert.Error(t, err, "Update() should return an error if UPDATE fails")

	// Nil receiver case
	var nilRepository *repository.TagsRepository
	tag, err := nilRepository.Update(context.Background(), tagMap)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, tag)
}

// TestTagsRepository_GetTags tests the GetTags method of TagsRepository.
func TestTagsRepository_GetTags(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := repository.NewTagsRepository(sqlxDB, "tags", nil, nil)

	mock.ExpectQuery("SELECT t.id, t.parent_id, t.title").WillReturnRows(
		sqlmock.NewRows([]string{
			"id", "parent_id", "title", "slug", "enabled", "created_by_id", "updated_by_id",
			"blog_id", "page_id",
		}).AddRow(
			1, 0, "test-tag", "test-slug", true, 1, 1,
			"{1,2}", "{3,4}",
		),
	)

	tags, err := repo.GetTags(context.Background())
	assert.NoError(t, err, "GetTagByID() should not return an error")
	assert.NotNil(t, tags, "GetTagByID() should return a valid tag")

	// Nil receiver case
	var nilRepository *repository.TagsRepository
	tags, err = nilRepository.GetTags(context.Background())
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, tags)

	t.Run("Tags Found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "parent_id", "title", "slug", "enabled", "created_by", "updated_by", "blog_id", "page_id"}).
			AddRow(1, 0, "Sample Tag", "sample-tag", true, 1, 1, nil, nil)

		mock.ExpectQuery("SELECT").
			WillReturnRows(rows)

		tags, err := repo.GetTags(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, tags)
	})

	t.Run("Tags Not Found", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WillReturnError(sql.ErrNoRows)

		tags, err := repo.GetTags(context.Background())
		assert.Error(t, err)
		assert.Nil(t, tags)
		assert.Contains(t, err.Error(), "sql: no rows in result set")
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WillReturnError(errors.New("database error"))

		tags, err := repo.GetTags(context.Background())
		assert.Error(t, err)
		assert.Nil(t, tags)
		assert.Contains(t, err.Error(), "database error")
	})

	t.Run("Scan Error", func(t *testing.T) {
		// Simulate a row scan error by returning incompatible data type
		rows := sqlmock.NewRows([]string{"id", "parent_id", "title", "slug", "enabled", "created_by", "updated_by", "blog_id", "page_id"}).
			AddRow("invalid-id", "invalid-parent-id", "invalid-title", "invalid-slug", "invalid-enabled", "invalid-created_by", "invalid-updated_by", "invalid-blog_id", "invalid-page_id")

		mock.ExpectQuery("SELECT").
			WillReturnRows(rows)

		tags, err := repo.GetTags(context.Background())
		assert.Error(t, err)
		assert.Nil(t, tags)
		assert.Contains(t, err.Error(), "sql: Scan error")
	})

	// Ensure all expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestTagsRepository_GetTagByID tests the GetTagByID method of TagsRepository.
func TestTagsRepository_GetTagByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repo := repository.NewTagsRepository(sqlxDB, "tags", nil, nil)

	mock.ExpectQuery("SELECT t.id, t.parent_id, t.title").WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{
			"id", "parent_id", "title", "slug", "enabled", "created_by_id", "updated_by_id",
			"blog_id", "page_id",
		}).AddRow(
			1, 0, "test-tag", "test-slug", true, 1, 1,
			"{1,2}", "{3,4}",
		),
	)

	tag, err := repo.GetTagByID(context.Background(), 1)
	assert.NoError(t, err, "GetTagByID() should not return an error")
	assert.NotNil(t, tag, "GetTagByID() should return a valid tag")

	// Nil receiver case
	var nilRepository *repository.TagsRepository
	tag, err = nilRepository.GetTagByID(context.Background(), 1)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, tag)

	t.Run("Tag Found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "parent_id", "title", "slug", "enabled", "created_by", "updated_by", "blog_id", "page_id"}).
			AddRow(1, 0, "Sample Tag", "sample-tag", true, 1, 1, nil, nil)

		mock.ExpectQuery("SELECT").
			WithArgs(1).
			WillReturnRows(rows)

		tag, err := repo.GetTagByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, tag)
		assert.Equal(t, int64(1), tag.ID)
		assert.Equal(t, "Sample Tag", tag.Title)
	})

	t.Run("Tag Not Found", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(2).
			WillReturnError(sql.ErrNoRows)

		tag, err := repo.GetTagByID(context.Background(), 2)
		assert.Error(t, err)
		assert.Nil(t, tag)
		assert.Contains(t, err.Error(), "tag not found")
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(3).
			WillReturnError(errors.New("database error"))

		tag, err := repo.GetTagByID(context.Background(), 3)
		assert.Error(t, err)
		assert.Nil(t, tag)
		assert.Contains(t, err.Error(), "database error")
	})

	// Ensure all expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}
