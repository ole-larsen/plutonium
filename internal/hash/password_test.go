package hash_test

import (
	"testing"

	"github.com/ole-larsen/plutonium/internal/hash"
)

func TestHashPassword(t *testing.T) {
	// Test case 1: Successful hashing
	password := []byte("securepassword")

	hashedPassword, err := hash.Password(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("Expected non-empty hashed password, got empty string")
	}

	// Test case 2: bcrypt.GenerateFromPassword fails
	password = []byte("securepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepasswordsecurepassword")

	hashedPassword, err = hash.Password(password)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	if hashedPassword != "" {
		t.Fatalf("Expected empty hashed password, got empty string")
	}
}

func TestComparePassword(t *testing.T) {
	// Test case 1: Successful password comparison
	password := []byte("securepassword")

	hashedPassword, err := hash.Password(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	match, err := hash.ComparePassword(hashedPassword, password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !match {
		t.Fatalf("Expected passwords to match, but they didn't")
	}

	// Test case 2: Incorrect password comparison
	wrongPassword := []byte("wrongpassword")

	match, err = hash.ComparePassword(hashedPassword, wrongPassword)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	if match {
		t.Fatalf("Expected passwords not to match, but they did")
	}

	// Test case 3: bcrypt.CompareHashAndPassword fails with invalid hash
	invalidHash := "invalidHash"

	match, err = hash.ComparePassword(invalidHash, password)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	if match {
		t.Fatalf("Expected passwords not to match due to invalid hash, but they did")
	}
}
