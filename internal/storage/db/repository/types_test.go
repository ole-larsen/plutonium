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
