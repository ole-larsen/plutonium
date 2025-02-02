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

	// Valid JSON as []byte
	validJSON := []byte(`{"name":"Menu1","items":[{"name":"Item1"}]}`)
	err := attr.Scan(validJSON)
	assert.NoError(t, err)
	assert.Equal(t, "Menu1", attr.Name)

	// Valid JSON as string
	validJSONString := `{"name":"Menu2","items":[{"name":"Item2"}]}`
	err = attr.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Equal(t, "Menu2", attr.Name)

	// Invalid JSON
	invalidJSON := []byte(`{"name":"Menu3",`)
	err = attr.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type (non-string/non-byte slice)
	err = attr.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedMenuAttributesJSON_Value(t *testing.T) {
	attr := repo.AggregatedMenuAttributesJSON{
		Name:  "Menu1",
		Items: []*models.PublicMenu{},
	}

	// Test Value method
	val, err := attr.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(attr)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedSliderItemJSON_Scan(t *testing.T) {
	var sliderItems repo.AggregatedSliderItemJSON

	// Valid JSON as []byte
	validJSON := []byte(`[
		{"bg": {"url": "bg1.jpg"}, "image": {"url": "image1.jpg"}, "btnLink1": "https://example.com/1", "description": "Slider item 1", "heading": "Heading 1", "id": 1}
	]`)
	err := sliderItems.Scan(validJSON)
	assert.NoError(t, err)
	assert.Len(t, sliderItems, 1)
	assert.Equal(t, "Slider item 1", sliderItems[0].Description)

	// Valid JSON as string
	validJSONString := `[{"bg": {"url": "bg2.jpg"}, "image": {"url": "image2.jpg"}, "btnLink1": "https://example.com/2", "description": "Slider item 2", "heading": "Heading 2", "id": 2}]`
	err = sliderItems.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Len(t, sliderItems, 1)
	assert.Equal(t, "Slider item 2", sliderItems[0].Description)

	// Invalid JSON
	invalidJSON := []byte(`[{"bg": {"url": "bg3.jpg"}, "image": {"url": "image3.jpg"}]`)
	err = sliderItems.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = sliderItems.Scan(123)
	assert.NoError(t, err)
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
			Description: "Slider item 1",
			Heading:     "Heading 1",
			ID:          1,
		},
	}

	// Test Value method
	val, err := sliderItems.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(sliderItems)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedCategoryCollectionsCollectiblesJSON_Scan(t *testing.T) {
	var collectibles repo.AggregatedCategoryCollectionsCollectiblesJSON

	// Valid JSON as []byte
	validJSON := []byte(`[{"id": 1, "name": "Collectible 1"}, {"id": 2, "name": "Collectible 2"}]`)
	err := collectibles.Scan(validJSON)
	assert.NoError(t, err)
	assert.Len(t, collectibles, 2)

	// Valid JSON as string
	validJSONString := `[{"id": 3, "name": "Collectible 3"}, {"id": 4, "name": "Collectible 4"}]`
	err = collectibles.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Len(t, collectibles, 2)

	// Invalid JSON
	invalidJSON := []byte(`[{"id": 5, "name": "Collectible 5"}`)
	err = collectibles.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = collectibles.Scan(123)
	assert.NoError(t, err)
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

	// Test Value method
	val, err := collectibles.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(collectibles)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedCategoryCollectionsJSON_Scan(t *testing.T) {
	var collections repo.AggregatedCategoryCollectionsJSON

	// Valid JSON as []byte
	validJSON := []byte(`[{"id": 1, "title": "Category 1"}, {"id": 2, "title": "Category 2"}]`)
	err := collections.Scan(validJSON)
	assert.NoError(t, err)
	assert.Len(t, collections, 2)

	// Valid JSON as string
	validJSONString := `[{"id": 3, "title": "Category 3"}, {"id": 4, "title": "Category 4"}]`
	err = collections.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Len(t, collections, 2)

	// Invalid JSON
	invalidJSON := []byte(`[{"id": 5, "title": "Category 5"}`)
	err = collections.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = collections.Scan(123)
	assert.NoError(t, err)
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

	// Test Value method
	val, err := collections.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(collections)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedImageJSON_Scan(t *testing.T) {
	var image repo.AggregatedImageJSON

	// Valid JSON as []byte
	validJSON := []byte(`{"url": "image.jpg", "size": 1024}`)
	err := image.Scan(validJSON)
	assert.NoError(t, err)

	// Valid JSON as string
	validJSONString := `{"url": "image2.jpg", "size": 2048}`
	err = image.Scan(validJSONString)
	assert.NoError(t, err)

	// Invalid JSON
	invalidJSON := []byte(`{"url": "image3.jpg"}`)
	err = image.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = image.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedImageJSON_Value(t *testing.T) {
	image := repo.AggregatedImageJSON{
		Attributes: &models.PublicFileAttributes{
			URL: "image.jpg", Size: 1024,
		},
	}

	// Test Value method
	val, err := image.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(image)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedWallet_Scan(t *testing.T) {
	var wallet repo.AggregatedWallet

	// Valid JSON as []byte
	validJSON := []byte(`[{"id": 1, "address": "0xabc", "balance": 1000}]`)
	err := wallet.Scan(validJSON)
	assert.NoError(t, err)
	assert.Len(t, wallet, 1)
	assert.Equal(t, "0xabc", wallet[0].Address)

	// Valid JSON as string
	validJSONString := `[{"id": 2, "address": "0xdef", "balance": 2000}]`
	err = wallet.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Len(t, wallet, 1)
	assert.Equal(t, "0xdef", wallet[0].Address)

	// Invalid JSON
	invalidJSON := []byte(`[{"id": 3, "address": "0xghi"}`)
	err = wallet.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = wallet.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedWallet_Value(t *testing.T) {
	wallet := repo.AggregatedWallet{
		&models.Wallet{
			Address: "0xabc",
			Name:    "Metamask",
		},
	}

	// Test Value method
	val, err := wallet.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(wallet)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedSocial_Scan(t *testing.T) {
	var social repo.AggregatedSocial

	// Valid JSON as []byte
	validJSON := []byte(`[{"icon": "Twitter", "username": "user1", "followers": 1000}]`)
	err := social.Scan(validJSON)
	assert.NoError(t, err)
	assert.Len(t, social, 1)
	assert.Equal(t, "Twitter", social[0].Icon)

	// Valid JSON as string
	validJSONString := `[{"icon": "Facebook", "username": "user2", "followers": 2000}]`
	err = social.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Len(t, social, 1)
	assert.Equal(t, "Facebook", social[0].Icon)

	// Invalid JSON
	invalidJSON := []byte(`[{"icon": "Instagram", "username": "user3"}`)
	err = social.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = social.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedSocial_Value(t *testing.T) {
	social := repo.AggregatedSocial{
		&models.Social{
			Icon: "Twitter",
		},
	}

	// Test Value method
	val, err := social.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(social)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedAuthorJSON_Scan(t *testing.T) {
	var author repo.AggregatedAuthorJSON

	// Valid JSON as []byte
	validJSON := []byte(`{"id": 1, "attributes": {"title": "Author 1", "bio": "Author bio"}}`)
	err := author.Scan(validJSON)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), author.ID)

	// Valid JSON as string
	validJSONString := `{"id": 2,  "attributes": {"title": "Author 2", "bio": "Author bio"}}`
	err = author.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Equal(t, "Author 2", author.Attributes.Title)

	// Invalid JSON
	invalidJSON := []byte(`{"id": 3, "name": "Author 3"}`)
	err = author.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = author.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedAuthorJSON_Value(t *testing.T) {
	author := repo.AggregatedAuthorJSON{
		ID: 1,
		Attributes: &models.PublicAuthorAttributes{
			Name: "Author 1",
		},
	}

	// Test Value method
	val, err := author.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(author)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedAuthorAttributes_Scan(t *testing.T) {
	var authorAttributes repo.AggregatedAuthorAttributes

	// Valid JSON as []byte
	validJSON := []byte(`{"name": "Author 1", "bio": "Author bio"}`)
	err := authorAttributes.Scan(validJSON)
	assert.NoError(t, err)
	assert.Equal(t, "Author 1", authorAttributes.Name)

	// Valid JSON as string
	validJSONString := `{"name": "Author 2", "bio": "Author bio 2"}`
	err = authorAttributes.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Equal(t, "Author 2", authorAttributes.Name)

	// Invalid JSON
	invalidJSON := []byte(`{"name": "Author 3"}`)
	err = authorAttributes.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = authorAttributes.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedAuthorAttributes_Value(t *testing.T) {
	authorAttributes := repo.AggregatedAuthorAttributes{
		Name: "Author 1",
	}

	// Test Value method
	val, err := authorAttributes.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(authorAttributes)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedPageAttributes_Scan(t *testing.T) {
	var pageAttributes repo.AggregatedPageAttributes

	// Valid JSON as []byte
	validJSON := []byte(`{"title": "Page 1", "description": "Page description"}`)
	err := pageAttributes.Scan(validJSON)
	assert.NoError(t, err)
	assert.Equal(t, "Page 1", pageAttributes.Title)

	// Valid JSON as string
	validJSONString := `{"title": "Page 2", "description": "Page description 2"}`
	err = pageAttributes.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Equal(t, "Page 2", pageAttributes.Title)

	// Invalid JSON
	invalidJSON := []byte(`{"title": "Page 3"}`)
	err = pageAttributes.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = pageAttributes.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedPageAttributes_Value(t *testing.T) {
	pageAttributes := repo.AggregatedPageAttributes{
		Title:       "Page 1",
		Description: "Page description",
	}

	// Test Value method
	val, err := pageAttributes.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(pageAttributes)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedContactAttributes_Scan(t *testing.T) {
	var contactAttributes repo.AggregatedContactAttributes

	// Valid JSON as []byte
	validJSON := []byte(`{"heading": "contact@example.com", "phone": "123456789"}`)
	err := contactAttributes.Scan(validJSON)
	assert.NoError(t, err)
	assert.Equal(t, "contact@example.com", contactAttributes.Heading)

	// Valid JSON as string
	validJSONString := `{"heading": "contact2@example.com", "phone": "987654321"}`
	err = contactAttributes.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Equal(t, "contact2@example.com", contactAttributes.Heading)

	// Invalid JSON
	invalidJSON := []byte(`{"email": "contact3@example.com"}`)
	err = contactAttributes.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = contactAttributes.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedContactAttributes_Value(t *testing.T) {
	contactAttributes := repo.AggregatedContactAttributes{
		Heading:    "contact@example.com",
		SubHeading: "123456789",
	}

	// Test Value method
	val, err := contactAttributes.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(contactAttributes)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}

func TestAggregatedTagJSON_Scan(t *testing.T) {
	var tags repo.AggregatedTagJSON

	// Valid JSON as []byte
	validJSON := []byte(`[{"title": "Tag 1", "description": "Description 1"}]`)
	err := tags.Scan(validJSON)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "Tag 1", tags[0].Title)

	// Valid JSON as string
	validJSONString := `[{"title": "Tag 2", "description": "Description 2"}]`
	err = tags.Scan(validJSONString)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "Tag 2", tags[0].Title)

	// Invalid JSON
	invalidJSON := []byte(`[{"title": "Tag 3", "description": "Description 3"}`)
	err = tags.Scan(invalidJSON)
	assert.NoError(t, err)

	// Invalid type
	err = tags.Scan(123)
	assert.NoError(t, err)
}

func TestAggregatedTagJSON_Value(t *testing.T) {
	tags := repo.AggregatedTagJSON{
		{Title: "Tag 1", Link: "Description 1"},
	}

	// Test Value method
	val, err := tags.Value()
	assert.NoError(t, err)

	valBytes, ok := val.([]byte)
	assert.True(t, ok)

	expectedJSON, err := json.Marshal(tags)
	require.NoError(t, err)
	assert.Equal(t, expectedJSON, valBytes)
}
