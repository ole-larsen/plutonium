package repository_test

import (
	"errors"
	"testing"

	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

func TestError_Error(t *testing.T) {
	// Create a standard error
	stdErr := errors.New("something went wrong")
	// Create a custom Error instance
	customErr := repository.NewError(stdErr)

	// Use errors.As to perform the type assertion
	var repositoryErr *repository.Error
	if !errors.As(customErr, &repositoryErr) {
		// Type assertion failed
		t.Fatalf("expected *Error, got %T", customErr)
	}

	// Test Error method
	expectedMsg := "[repository]: something went wrong"
	if repositoryErr.Error() != expectedMsg {
		t.Errorf("expected error message %q, got %q", expectedMsg, repositoryErr.Error())
	}
}

func TestError_Unwrap(t *testing.T) {
	// Create a standard error
	stdErr := errors.New("something went wrong")
	// Create a custom Error instance
	customErr := repository.NewError(stdErr)

	// Use errors.As to perform the type assertion
	var serverErr *repository.Error
	if !errors.As(customErr, &serverErr) {
		// Type assertion failed
		t.Fatalf("expected *Error, got %T", customErr)
	}

	// Test Unwrap method
	if !errors.Is(serverErr.Unwrap(), stdErr) {
		t.Errorf("expected %v, got %v", stdErr, serverErr.Unwrap())
	}
}

func TestNewError(t *testing.T) {
	// Test with a non-nil error
	stdErr := errors.New("something went wrong")

	err := repository.NewError(stdErr)
	if err == nil {
		t.Fatal("expected non-nil error")
	}

	// Use errors.As to perform the type assertion
	var serverErr *repository.Error
	if !errors.As(err, &serverErr) {
		// Type assertion failed
		t.Fatalf("expected *Error, got %T", err)
	}

	// Ensure the underlying error is the same
	if !errors.Is(serverErr.Unwrap(), stdErr) {
		t.Errorf("expected %v, got %v", stdErr, serverErr.Unwrap())
	}

	// Test with a nil error
	err = repository.NewError(nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
