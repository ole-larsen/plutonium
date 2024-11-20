package hash_test

import (
	"testing"
	"unicode"

	"github.com/ole-larsen/plutonium/internal/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandStringBytes_Length(t *testing.T) {
	// Test cases for different string lengths
	lengths := []int{0, 1, 5, 10, 50, 100}

	for _, length := range lengths {
		t.Run("Length_"+string(rune(length)), func(t *testing.T) {
			randomString := hash.RandStringBytes(length)
			require.Len(t, randomString, length, "Generated string should have the correct length")
		})
	}
}

func TestRandStringBytes_ValidCharacters(t *testing.T) {
	// Generate a string of length 100
	randomString := hash.RandStringBytes(100)

	// Ensure that all characters are within the valid character set
	for _, char := range randomString {
		assert.True(t, unicode.IsLetter(char), "String should only contain letters")
	}
}

func TestRandStringBytes_Randomness(t *testing.T) {
	// Generate two strings of the same length
	randomString1 := hash.RandStringBytes(20)
	randomString2 := hash.RandStringBytes(20)

	// Ensure that the two strings are not the same
	assert.NotEqual(t, randomString1, randomString2, "Randomly generated strings should be different")
}

func TestRandStringBytes_ZeroLength(t *testing.T) {
	// Test generating a string of length 0
	randomString := hash.RandStringBytes(0)
	assert.Equal(t, "", randomString, "String with zero length should be empty")
}
