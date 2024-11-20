package storage_test

import (
	"errors"
	"testing"

	"github.com/ole-larsen/plutonium/internal/storage"
)

func TestError_Error(t *testing.T) {
	// Create a standard error
	stdErr := errors.New("something went wrong")
	// Create a custom Error instance
	customErr := storage.NewError(stdErr)

	// Use errors.As to perform the type assertion
	var storageErr *storage.Error
	if !errors.As(customErr, &storageErr) {
		// Type assertion failed
		t.Fatalf("expected *Error, got %T", customErr)
	}

	// Test Error method
	expectedMsg := "[storage]: something went wrong"
	if storageErr.Error() != expectedMsg {
		t.Errorf("expected error message %q, got %q", expectedMsg, storageErr.Error())
	}
}

func TestError_Unwrap(t *testing.T) {
	// Create a standard error
	stdErr := errors.New("something went wrong")
	// Create a custom Error instance
	customErr := storage.NewError(stdErr)

	// Use errors.As to perform the type assertion
	var storageErr *storage.Error
	if !errors.As(customErr, &storageErr) {
		// Type assertion failed
		t.Fatalf("expected *Error, got %T", customErr)
	}

	// Test Unwrap method
	if !errors.Is(storageErr.Unwrap(), stdErr) {
		t.Errorf("expected %v, got %v", stdErr, storageErr.Unwrap())
	}
}

func TestNewError(t *testing.T) {
	// Test with a non-nil error
	stdErr := errors.New("something went wrong")

	err := storage.NewError(stdErr)
	if err == nil {
		t.Fatal("expected non-nil error")
	}

	// Use errors.As to perform the type assertion
	var storageErr *storage.Error
	if !errors.As(err, &storageErr) {
		// Type assertion failed
		t.Fatalf("expected *Error, got %T", err)
	}

	// Ensure the underlying error is the same
	if !errors.Is(storageErr.Unwrap(), stdErr) {
		t.Errorf("expected %v, got %v", stdErr, storageErr.Unwrap())
	}

	// Test with a nil error
	err = storage.NewError(nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
