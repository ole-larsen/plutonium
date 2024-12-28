package repository_test

import (
	"encoding/json"
	"testing"

	repo "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAggregatedMenuAttributesJSON_Scan(t *testing.T) {
	var attr repo.AggregatedMenuAttributesJSON

	// Case: Valid JSON as []byte
	validJSON := []byte(`{"name":"Menu1","items":[{"name":"Item1"}]}`)
	err := attr.Scan(validJSON)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON")
	assert.Equal(t, "Menu1", attr.Name, "Scan() should correctly parse the JSON")

	// Case: Valid JSON as string
	validJSONString := `{"name":"Menu2","items":[{"name":"Item2"}]}`
	err = attr.Scan(validJSONString)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON string")
	assert.Equal(t, "Menu2", attr.Name, "Scan() should correctly parse the JSON string")

	// Case: Invalid JSON
	invalidJSON := []byte(`{"name":"Menu3",`)
	err = attr.Scan(invalidJSON)
	assert.NoError(t, err, "Scan() should ignore invalid JSON and not return an error")

	// Case: Invalid type
	err = attr.Scan(123) // Not a []byte or string
	assert.NoError(t, err, "Scan() should not return an error for unsupported types")
}

func TestAggregatedMenuAttributesJSON_Value(t *testing.T) {
	attr := repo.AggregatedMenuAttributesJSON{
		Name:  "Menu1",
		Items: []*models.PublicMenu{},
	}

	// Marshal to JSON
	val, err := attr.Value()
	assert.NoError(t, err, "Value() should not return an error")

	// Convert `val` to `[]byte` for comparison
	valBytes, ok := val.([]byte)
	assert.True(t, ok, "Value() should return a []byte")

	expectedJSON, err := json.Marshal(attr)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes, "Value() should return the correct JSON representation")
}

func TestAggregatedSliderItemJSON_Scan(t *testing.T) {
	var sliderItems repo.AggregatedSliderItemJSON

	// Case: Valid JSON as []byte
	validJSON := []byte(`[
		{
			"bg": {"url": "bg1.jpg"},
			"image": {"url": "image1.jpg"},
			"btnLink1": "https://example.com/1",
			"btnLink2": "https://example.com/2",
			"btnText1": "Button 1",
			"btnText2": "Button 2",
			"description": "Slider item 1",
			"heading": "Heading 1",
			"id": 1
		}
	]`)
	err := sliderItems.Scan(validJSON)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON")
	assert.Len(t, sliderItems, 1, "Scan() should correctly parse the JSON into a slice of items")
	item := sliderItems[0]
	assert.Equal(t, "Slider item 1", item.Description, "Scan() should correctly parse the description")
	assert.Equal(t, "Heading 1", item.Heading, "Scan() should correctly parse the heading")
	assert.Equal(t, int64(1), item.ID, "Scan() should correctly parse the ID")
	assert.Equal(t, "https://example.com/1", item.BtnLink1, "Scan() should correctly parse the btnLink1")
	assert.Equal(t, "Button 1", item.BtnText1, "Scan() should correctly parse the btnText1")
	assert.NotNil(t, item.Bg, "Scan() should parse the Bg field correctly")
	assert.NotNil(t, item.Image, "Scan() should parse the Image field correctly")

	// Case: Valid JSON as string
	validJSONString := `[{
		"bg": {"url": "bg2.jpg"},
		"image": {"url": "image2.jpg"},
		"btnLink1": "https://example.com/3",
		"btnLink2": "https://example.com/4",
		"btnText1": "Button 3",
		"btnText2": "Button 4",
		"description": "Slider item 2",
		"heading": "Heading 2",
		"id": 2
	}]`
	err = sliderItems.Scan(validJSONString)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON string")
	assert.Len(t, sliderItems, 1, "Scan() should correctly parse the JSON string into a slice of items")
	item = sliderItems[0]
	assert.Equal(t, "Slider item 2", item.Description, "Scan() should correctly parse the description")
	assert.Equal(t, "Heading 2", item.Heading, "Scan() should correctly parse the heading")
	assert.Equal(t, int64(2), item.ID, "Scan() should correctly parse the ID")
	assert.Equal(t, "https://example.com/3", item.BtnLink1, "Scan() should correctly parse the btnLink1")
	assert.Equal(t, "Button 3", item.BtnText1, "Scan() should correctly parse the btnText1")
	assert.NotNil(t, item.Bg, "Scan() should parse the Bg field correctly")
	assert.NotNil(t, item.Image, "Scan() should parse the Image field correctly")

	// Case: Invalid JSON
	invalidJSON := []byte(`[{"bg": {"url": "bg3.jpg"}, "image": {"url": "image3.jpg"}]`)
	err = sliderItems.Scan(invalidJSON)
	assert.NoError(t, err, "Scan() should ignore invalid JSON and not return an error")

	// Case: Invalid type
	err = sliderItems.Scan(123) // Not a []byte or string
	assert.NoError(t, err, "Scan() should not return an error for unsupported types")
}

func TestAggregatedSliderItemJSON_Value(t *testing.T) {
	sliderItems := repo.AggregatedSliderItemJSON{
		&models.PublicSliderItem{
			Bg: &models.PublicFile{
				Attributes: &models.PublicFileAttributes{
					URL: "bg1.jpg",
				},
			},
			Image: &models.PublicFile{
				Attributes: &models.PublicFileAttributes{
					URL: "image1.jpg",
				},
			},
			BtnLink1:    "https://example.com/1",
			BtnLink2:    "https://example.com/2",
			BtnText1:    "Button 1",
			BtnText2:    "Button 2",
			Description: "Slider item 1",
			Heading:     "Heading 1",
			ID:          1,
		},
	}

	// Marshal to JSON
	val, err := sliderItems.Value()
	assert.NoError(t, err, "Value() should not return an error")

	// Convert `val` to `[]byte` for comparison
	valBytes, ok := val.([]byte)
	assert.True(t, ok, "Value() should return a []byte")

	expectedJSON, err := json.Marshal(sliderItems)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes, "Value() should return the correct JSON representation")
}

func TestAggregatedSliderItemJSON_Empty(t *testing.T) {
	var sliderItems repo.AggregatedSliderItemJSON

	// Case: Empty slice
	val, err := sliderItems.Value()
	assert.NoError(t, err, "Value() should not return an error for an empty slice")

	// Convert `val` to `[]byte` for comparison
	valBytes, ok := val.([]byte)
	assert.True(t, ok, "Value() should return a []byte")

	// Expected empty JSON array
	expectedJSON := `null`
	assert.Equal(t, expectedJSON, string(valBytes), "Value() should return an empty JSON array for an empty slice")
}

func TestAggregatedCategoryCollectionsCollectiblesJSON_Scan(t *testing.T) {
	var collectibles repo.AggregatedCategoryCollectionsCollectiblesJSON

	// Case: Valid JSON as []byte
	validJSON := []byte(`[{"id": 1, "name": "Collectible 1"}, {"id": 2, "name": "Collectible 2"}]`)
	err := collectibles.Scan(validJSON)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON")
	assert.Len(t, collectibles, 2, "Scan() should correctly parse the JSON into a slice of collectibles")
	assert.Equal(t, int64(1), collectibles[0].ID, "Scan() should correctly parse the ID of the first collectible")
	assert.Equal(t, int64(2), collectibles[1].ID, "Scan() should correctly parse the ID of the second collectible")

	// Case: Valid JSON as string
	validJSONString := `[{"id": 3, "name": "Collectible 3"}, {"id": 4, "name": "Collectible 4"}]`
	err = collectibles.Scan(validJSONString)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON string")
	assert.Len(t, collectibles, 2, "Scan() should correctly parse the JSON string into a slice of collectibles")
	assert.Equal(t, int64(3), collectibles[0].ID, "Scan() should correctly parse the ID of the first collectible")

	// Case: Invalid JSON
	invalidJSON := []byte(`[{"id": 5, "name": "Collectible 5"}`)
	err = collectibles.Scan(invalidJSON)
	assert.NoError(t, err, "Scan() should ignore invalid JSON and not return an error")

	// Case: Invalid type
	err = collectibles.Scan(123) // Not a []byte or string
	assert.NoError(t, err, "Scan() should not return an error for unsupported types")
}

func TestAggregatedCategoryCollectionsCollectiblesJSON_Value(t *testing.T) {
	collectibles := repo.AggregatedCategoryCollectionsCollectiblesJSON{
		{
			ID: 1,
			Attributes: repo.CategoryCollectionCollectibleAttributes{
				Metadata: repo.CollectibleMetadata{
					Name: "Collectible 1",
				},
			},
		},
		{
			ID: 2,
			Attributes: repo.CategoryCollectionCollectibleAttributes{
				Metadata: repo.CollectibleMetadata{
					Name: "Collectible 2",
				},
			},
		},
	}

	// Marshal to JSON
	val, err := collectibles.Value()
	assert.NoError(t, err, "Value() should not return an error")

	// Convert `val` to `[]byte` for comparison
	valBytes, ok := val.([]byte)
	assert.True(t, ok, "Value() should return a []byte")

	expectedJSON, err := json.Marshal(collectibles)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes, "Value() should return the correct JSON representation")
}

func TestAggregatedCategoryCollectionsJSON_Scan(t *testing.T) {
	var collections repo.AggregatedCategoryCollectionsJSON

	// Case: Valid JSON as []byte
	validJSON := []byte(`[{"id": 1, "title": "Category 1"}, {"id": 2, "title": "Category 2"}]`)
	err := collections.Scan(validJSON)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON")
	assert.Len(t, collections, 2, "Scan() should correctly parse the JSON into a slice of collections")
	assert.Equal(t, int64(1), collections[0].ID, "Scan() should correctly parse the ID of the first collection")
	assert.Equal(t, int64(2), collections[1].ID, "Scan() should correctly parse the ID of the second collection")

	// Case: Valid JSON as string
	validJSONString := `[{"id": 3, "title": "Category 3"}, {"id": 4, "title": "Category 4"}]`
	err = collections.Scan(validJSONString)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON string")
	assert.Len(t, collections, 2, "Scan() should correctly parse the JSON string into a slice of collections")
	assert.Equal(t, int64(3), collections[0].ID, "Scan() should correctly parse the ID of the first collection")

	// Case: Invalid JSON
	invalidJSON := []byte(`[{"id": 5, "title": "Category 5"}`)
	err = collections.Scan(invalidJSON)
	assert.NoError(t, err, "Scan() should ignore invalid JSON and not return an error")

	// Case: Invalid type
	err = collections.Scan(123) // Not a []byte or string
	assert.NoError(t, err, "Scan() should not return an error for unsupported types")
}

func TestAggregatedCategoryCollectionsJSON_Value(t *testing.T) {
	collections := repo.AggregatedCategoryCollectionsJSON{
		{
			ID: 1,
			Attributes: repo.CategoryCollectionAttributes{
				Name: "Category 1",
			},
		},
		{
			ID: 2,
			Attributes: repo.CategoryCollectionAttributes{
				Name: "Category 2",
			},
		},
	}

	// Marshal to JSON
	val, err := collections.Value()
	assert.NoError(t, err, "Value() should not return an error")

	// Convert `val` to `[]byte` for comparison
	valBytes, ok := val.([]byte)
	assert.True(t, ok, "Value() should return a []byte")

	expectedJSON, err := json.Marshal(collections)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes, "Value() should return the correct JSON representation")
}

func TestAggregatedImageJSON_Scan(t *testing.T) {
	var image repo.AggregatedImageJSON

	// Case: Valid JSON as []byte
	validJSON := []byte(`{"url": "image.jpg", "size": 1024}`)
	err := image.Scan(validJSON)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON")

	// Case: Valid JSON as string
	validJSONString := `{"url": "image2.jpg", "size": 2048}`
	err = image.Scan(validJSONString)
	assert.NoError(t, err, "Scan() should not return an error for valid JSON string")

	// Case: Invalid JSON
	invalidJSON := []byte(`{"url": "image3.jpg"}`)
	err = image.Scan(invalidJSON)
	assert.NoError(t, err, "Scan() should ignore invalid JSON and not return an error")

	// Case: Invalid type
	err = image.Scan(123) // Not a []byte or string
	assert.NoError(t, err, "Scan() should not return an error for unsupported types")
}

func TestAggregatedImageJSON_Value(t *testing.T) {
	image := repo.AggregatedImageJSON{
		Attributes: &models.PublicFileAttributes{
			URL: "image.jpg", Size: 1024,
		},
	}

	// Marshal to JSON
	val, err := image.Value()
	assert.NoError(t, err, "Value() should not return an error")

	// Convert `val` to `[]byte` for comparison
	valBytes, ok := val.([]byte)
	assert.True(t, ok, "Value() should return a []byte")

	expectedJSON, err := json.Marshal(image)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes, "Value() should return the correct JSON representation")
}
