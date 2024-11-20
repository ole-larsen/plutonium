package hash_test

import (
	"testing"

	"github.com/ole-larsen/plutonium/internal/hash"
	"github.com/stretchr/testify/assert"
)

func TestGetMD5Hash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple text",
			input:    "hello",
			expected: "5d41402abc4b2a76b9719d911017c592", // MD5 for "hello"
		},
		{
			name:     "empty string",
			input:    "",
			expected: "d41d8cd98f00b204e9800998ecf8427e", // MD5 for empty string
		},
		{
			name:     "numeric input",
			input:    "123456",
			expected: "e10adc3949ba59abbe56e057f20f883e", // MD5 for "123456"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hash.GetMD5Hash(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateMD5Hash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hash     string
		expected bool
	}{
		{
			name:     "valid hash",
			input:    "hello",
			hash:     "5d41402abc4b2a76b9719d911017c592", // Correct MD5 for "hello"
			expected: true,
		},
		{
			name:     "invalid hash",
			input:    "hello",
			hash:     "invalidhash",
			expected: false,
		},
		{
			name:     "empty input",
			input:    "",
			hash:     "d41d8cd98f00b204e9800998ecf8427e", // MD5 for empty string
			expected: true,
		},
		{
			name:     "wrong hash for input",
			input:    "hello",
			hash:     "e10adc3949ba59abbe56e057f20f883e", // MD5 for "123456"
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hash.ValidateMD5Hash(tt.input, tt.hash)
			assert.Equal(t, tt.expected, result)
		})
	}
}
