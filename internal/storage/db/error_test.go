package db_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/internal/storage/db"
)

// Helper function to format time.
func formatTime(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}

// TestNewError tests the NewError function.
func TestNewError(t *testing.T) {
	// Test with a non-nil error and retry > 0
	err := errors.New("test error")
	retry := 3

	dbErr := db.NewError(err, retry)
	if dbErr == nil {
		t.Errorf("NewError() returned nil, expected non-nil")
	}

	if !errors.Is(dbErr, err) {
		t.Errorf("NewError() returned wrong error, expected %v, got %v", err, dbErr)
	}

	// Test with a non-nil error and retry = 0
	err = errors.New("test error")
	retry = 0

	dbErr = db.NewError(err, retry)
	if dbErr == nil {
		t.Errorf("NewError() returned nil, expected non-nil")
	}

	if !errors.Is(dbErr, err) {
		t.Errorf("NewError() returned wrong error, expected %v, got %v", err, dbErr)
	}

	// Test with a nil error
	dbErr = db.NewError(nil, 1)
	if dbErr != nil {
		t.Errorf("NewError() returned %v, expected nil", dbErr)
	}
}

// TestErrorMethod tests the Error method of the Error struct.
func TestErrorMethod(t *testing.T) {
	// Mock time
	currentTime := time.Now()
	formattedTime := formatTime(currentTime)

	// Test with retry > 0
	err := errors.New("test error")
	retry := 3
	dbErr := &db.Error{
		Err:   err,
		Retry: retry,
	}
	expectedMessage := fmt.Sprintf("%v attempt: %d %s", err, retry, formattedTime)

	if dbErr.Error() != expectedMessage {
		t.Errorf("Error() returned %v, expected %v", dbErr.Error(), expectedMessage)
	}

	// Test with retry = 0
	dbErr = &db.Error{
		Err:   err,
		Retry: 0,
	}
	expectedMessage = fmt.Sprintf("%v %s", err, formattedTime)

	if dbErr.Error() != expectedMessage {
		t.Errorf("Error() returned %v, expected %v", dbErr.Error(), expectedMessage)
	}
}

// TestUnwrapMethod tests the Unwrap method of the Error struct.
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
