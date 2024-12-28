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

func TestNewSlidersRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "sliders"
	repo := repository.NewSlidersRepository(sqlxDB, tbl)

	assert.NotNil(t, repo)
	assert.Equal(t, tbl, repo.TBL)

	// Nil database case
	repositoryNil := repository.NewSlidersRepository(nil, "sliders")
	assert.Nil(t, repositoryNil)
}

func TestSlidersRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewSlidersRepository(sqlxDB, "sliders")

	assert.Equal(t, sqlxDB, repo.InnerDB())

	// Nil receiver case
	var nilRepository *repository.SlidersRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestSlidersRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewSlidersRepository(sqlxDB, "sliders")

	mock.ExpectPing().WillReturnError(nil)

	err = repo.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repo.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	err = nilRepository.Ping()
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestSlidersRepository_PingEdgeCase(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewSlidersRepository(sqlxDB, "sliders")

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err = repo.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}

func setupTestSlidersRepo(t *testing.T) (*repository.SlidersRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewSlidersRepository(sqlxDB, "sliders")

	return repo, mock
}

func TestSlidersRepository_Create(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	// Test data
	sliderMap := map[string]interface{}{
		"provider":      "provider1",
		"title":         "Slider Title",
		"description":   "Slider Description",
		"enabled":       true,
		"created_by_id": 1,
		"updated_by_id": 1,
	}

	mock.ExpectExec(`INSERT INTO sliders`).
		WithArgs(sliderMap["provider"], sliderMap["title"], sliderMap["description"], sliderMap["enabled"], sliderMap["created_by_id"], sliderMap["updated_by_id"]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(ctx, sliderMap)
	assert.NoError(t, err)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	err = nilRepository.Create(ctx, sliderMap)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
}

func TestSlidersRepository_Update(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	// Test data
	sliderMap := map[string]interface{}{
		"id":            1,
		"provider":      "provider1",
		"title":         "Updated Slider Title",
		"description":   "Updated Slider Description",
		"enabled":       false,
		"updated_by_id": 2,
	}

	mock.ExpectExec(`UPDATE sliders SET`).
		WithArgs(sliderMap["provider"], sliderMap["title"], sliderMap["description"], sliderMap["enabled"], sliderMap["updated_by_id"], sliderMap["id"]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "title", "description", "enabled", "created_by_id", "updated_by_id"}).
			AddRow(1, "provider1", "Updated Slider Title", "Updated Slider Description", false, 1, 2))

	// Perform the update and verify
	sliders, err := repo.Update(ctx, sliderMap)
	assert.NoError(t, err)
	assert.Len(t, sliders, 1)
	assert.Equal(t, "Updated Slider Title", sliders[0].Title)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	sliders, err = nilRepository.Update(ctx, sliderMap)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, sliders)

	mock.ExpectExec(`UPDATE sliders SET`).
		WithArgs(sliderMap["provider"], sliderMap["title"], sliderMap["description"], sliderMap["enabled"], sliderMap["updated_by_id"], sliderMap["id"]).
		WillReturnError(errors.New("some error"))

	sliders, err = repo.Update(ctx, sliderMap)
	assert.Error(t, err)
	assert.Nil(t, sliders)
}

func TestSlidersRepository_GetSliders(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	// Test case for successful query
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "title", "description", "enabled", "created_by_id", "updated_by_id"}).
			AddRow(1, "provider1", "Slider Title", "Slider Description", true, 1, 2).
			AddRow(2, "provider2", "Another Slider", "Another Description", false, 2, 3))

	sliders, err := repo.GetSliders(ctx)
	assert.NoError(t, err)
	assert.Len(t, sliders, 2)
	assert.Equal(t, "Slider Title", sliders[0].Title)
	assert.Equal(t, "Another Slider", sliders[1].Title)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	sliders, err = nilRepository.GetSliders(ctx)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, sliders)

	// Test case for database query error
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WillReturnError(errors.New("some db error"))

	sliders, err = repo.GetSliders(ctx)
	assert.Error(t, err)
	assert.Nil(t, sliders)

	// Test case for scanning error
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "title", "description", "enabled", "created_by_id", "updated_by_id"}).
			AddRow(1, "provider1", "Slider Title", "Slider Description", true, 1, 2).
			AddRow(2, "provider2", nil, "Another Description", false, 2, 3)) // Simulating a scanning error (nil Title)

	sliders, err = repo.GetSliders(ctx)
	assert.Error(t, err)
	assert.Nil(t, sliders)
}

func TestSlidersRepository_GetSliderByTitle(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	// Test case where the slider is found
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WithArgs("Slider Title").
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "title", "description", "enabled", "created_by_id", "updated_by_id"}).
			AddRow(1, "provider1", "Slider Title", "Slider Description", true, 1, 2))

	slider, err := repo.GetSliderByTitle(ctx, "Slider Title")
	assert.NoError(t, err)
	assert.NotNil(t, slider)
	assert.Equal(t, "Slider Title", slider.Title)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	slider, err = nilRepository.GetSliderByTitle(ctx, "Slider Title")
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, slider)

	// Test case for sql.ErrNoRows (no slider found)
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WithArgs("Nonexistent Slider Title").
		WillReturnError(sql.ErrNoRows) // Simulating no rows found

	slider, err = repo.GetSliderByTitle(ctx, "Nonexistent Slider Title")
	assert.Error(t, err)
	assert.Nil(t, slider)
	assert.Equal(t, "slider not found", err.Error())

	// Test case for general database query error
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WithArgs("Slider Title").
		WillReturnError(errors.New("some db error"))

	slider, err = repo.GetSliderByTitle(ctx, "Slider Title")
	assert.Error(t, err)
	assert.Nil(t, slider)
}

func TestSlidersRepository_GetSliderByID(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	// Test case where the slider is found
	mock.ExpectQuery("SELECT id, provider, title, description, enabled, created_by_id, updated_by_id FROM sliders").
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "title", "description", "enabled", "created_by_id", "updated_by_id"}).
			AddRow(1, "provider1", "Slider Title", "Slider Description", true, 1, 2))

	slider, err := repo.GetSliderByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, slider)
	assert.Equal(t, "Slider Title", slider.Title)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	slider, err = nilRepository.GetSliderByID(ctx, 1)
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, slider)

	// Test case for sql.ErrNoRows (no slider found)
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WithArgs(int64(1)).
		WillReturnError(sql.ErrNoRows) // Simulating no rows found

	slider, err = repo.GetSliderByID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, slider)
	assert.Equal(t, "slider not found", err.Error()) // Checking for the "slider not found" error message

	// Test case for general database query error
	mock.ExpectQuery(`SELECT id, provider, title, description, enabled, created_by_id, updated_by_id`).
		WithArgs(int64(1)).
		WillReturnError(errors.New("some db error"))

	slider, err = repo.GetSliderByID(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, slider)
}

func TestSlidersRepository_Ping_Error(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	mock.ExpectPing().WillReturnError(errors.New("ping failed"))

	err := repo.Ping()
	assert.Error(t, err)
}

func TestSlidersRepository_GetSliderByProvider(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	sqlStatement := `
SELECT
	s.id`

	mock.ExpectQuery(sqlStatement).
		WithArgs("provider1").
		WillReturnError(sql.ErrNoRows) // Simulating no rows found

	slider, err := repo.GetSliderByProvider(ctx, "provider1")
	assert.Error(t, err)
	assert.Equal(t, "[repository]: slider not found", err.Error()) // Checking for the "slider not found" error message

	assert.Nil(t, slider)

	mock.ExpectQuery(sqlStatement).
		WithArgs("provider1").
		WillReturnError(errors.New("some error")) // Simulating no rows found

	slider, err = repo.GetSliderByProvider(ctx, "provider1")
	assert.Error(t, err)
	assert.Nil(t, slider)

	// Nil receiver case
	var nilRepository *repository.SlidersRepository
	slider, err = nilRepository.GetSliderByProvider(ctx, "provider1")
	assert.Equal(t, repository.ErrDBNotInitialized, err)
	assert.Nil(t, slider)
}

func TestSlidersRepository_GetSliderByProvider_ValidData(t *testing.T) {
	repo, mock := setupTestSlidersRepo(t)
	ctx := context.Background()

	// Mock the query and result
	mock.ExpectQuery(`
SELECT
	s.id`).
		WithArgs("provider1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "attributes"}).
			AddRow(1, `[{
                "heading": "Slide 1 Heading",
                "description": "Slide 1 Description",
                "btnLink1": "https://example.com/link1",
                "btnText1": "Button 1",
                "btnLink2": "https://example.com/link2",
                "btnText2": "Button 2",
                "image": {
                    "id": 101,
                    "attributes": {
                        "name": "image1.jpg",
                        "alt": "Alt Text",
                        "caption": "Image Caption",
                        "ext": ".jpg",
                        "provider": "local",
                        "width": 800,
                        "height": 600,
                        "size": 12345,
                        "url": "https://example.com/image1.jpg"
                    }
                },
                "bg": {
                    "id": 201,
                    "attributes": {
                        "name": "bg1.jpg",
                        "alt": "Background Alt Text",
                        "caption": "Background Caption",
                        "ext": ".jpg",
                        "provider": "local",
                        "width": 1920,
                        "height": 1080,
                        "size": 45678,
                        "url": "https://example.com/bg1.jpg"
                    }
                }
            }]`))

	// Call the function
	slider, err := repo.GetSliderByProvider(ctx, "provider1")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, slider)
	assert.Equal(t, int64(1), slider.ID)
	assert.NotNil(t, slider.Attributes)
	assert.Len(t, slider.Attributes.SlidesItem, 1)

	slide := slider.Attributes.SlidesItem[0]
	assert.Equal(t, "Slide 1 Heading", slide.Heading)
	assert.Equal(t, "Slide 1 Description", slide.Description)
	assert.Equal(t, "https://example.com/link1", slide.BtnLink1)
	assert.Equal(t, "Button 1", slide.BtnText1)
	assert.Equal(t, "https://example.com/link2", slide.BtnLink2)
	assert.Equal(t, "Button 2", slide.BtnText2)

	assert.NotNil(t, slide.Image)
	assert.Equal(t, int64(101), slide.Image.ID)
	assert.Equal(t, "image1.jpg", slide.Image.Attributes.Name)

	assert.NotNil(t, slide.Bg)
	assert.Equal(t, int64(201), slide.Bg.ID)
	assert.Equal(t, "bg1.jpg", slide.Bg.Attributes.Name)
}
