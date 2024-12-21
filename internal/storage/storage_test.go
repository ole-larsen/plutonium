package storage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

// MockStorage is a mock implementation of the Storage interface.
type MockStorage struct {
	ShouldError bool
}

// Init is a mock implementation of the Init method.
func (m *MockStorage) Init(_ context.Context) error {
	if m.ShouldError {
		return errors.New("mock init error")
	}

	return nil
}

// ConnectUsersRepository is a mock implementation of the ConnectUsersRepository method.
func (m *MockStorage) ConnectUsersRepository(_ *sqlx.DB) error {
	if m.ShouldError {
		return errors.New("mock connect error")
	}

	return nil
}

// ConnectContractsRepository is a mock implementation of the ConnectContractsRepository method.
func (m *MockStorage) ConnectContractsRepository(_ *sqlx.DB) error {
	if m.ShouldError {
		return errors.New("mock connect error")
	}

	return nil
}

// Ping is a mock implementation of the Ping method.
func (m *MockStorage) Ping() error {
	if m.ShouldError {
		return errors.New("mock ping error")
	}

	return nil
}

// GetUser is a mock implementation of the GetUser method.
func (m *MockStorage) GetUser(_ context.Context, _ string) (*repository.User, error) {
	if m.ShouldError {
		return nil, errors.New("mock get user error")
	}

	return &repository.User{}, nil
}

// TestStorageInitSuccess tests the successful case of the Init method.
func TestStorageInitSuccess(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: false}

	err := mockStorage.Init(context.TODO())
	if err != nil {
		t.Errorf("Init() returned an error: %v, expected nil", err)
	}
}

// TestStorageInitError tests the error case of the Init method.
func TestStorageInitError(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: true}

	err := mockStorage.Init(context.TODO())
	if err == nil {
		t.Errorf("Init() returned nil, expected an error")
	}
}

// TestStorageConnectUsersRepositorySuccess tests the successful case of the ConnectUsersRepository method.
func TestStorageConnectUsersRepositorySuccess(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: false}
	mockSQLxDB := &sqlx.DB{}

	err := mockStorage.ConnectUsersRepository(mockSQLxDB)
	if err != nil {
		t.Errorf("ConnectUsersRepository() returned an error: %v, expected nil", err)
	}
}

// TestStorageConnectUsersRepositoryError tests the error case of the ConnectUsersRepository method.
func TestStorageConnectUsersRepositoryError(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: true}
	mockSQLxDB := &sqlx.DB{}

	err := mockStorage.ConnectUsersRepository(mockSQLxDB)
	if err == nil {
		t.Errorf("ConnectUsersRepository() returned nil, expected an error")
	}
}

// TestStoragePingSuccess tests the successful case of the Ping method.
func TestStoragePingSuccess(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: false}

	err := mockStorage.Ping()
	if err != nil {
		t.Errorf("Ping() returned an error: %v, expected nil", err)
	}
}

// TestStoragePingError tests the error case of the Ping method.
func TestStoragePingError(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: true}

	err := mockStorage.Ping()
	if err == nil {
		t.Errorf("Ping() returned nil, expected an error")
	}
}

// CreateUser is a mock implementation of the CreateUser method.
func (m *MockStorage) CreateUser(_ context.Context, _ map[string]interface{}) error {
	if m.ShouldError {
		return errors.New("mock create user error")
	}

	return nil
}

// TestStorageCreateUserSuccess tests the successful case of the CreateUser method.
func TestStorageCreateUserSuccess(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: false}

	err := mockStorage.CreateUser(context.TODO(), map[string]interface{}{"name": "John Doe"})
	if err != nil {
		t.Errorf("CreateUser() returned an error: %v, expected nil", err)
	}
}

// TestStorageCreateUserError tests the error case of the CreateUser method.
func TestStorageCreateUserError(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: true}

	err := mockStorage.CreateUser(context.TODO(), map[string]interface{}{"name": "John Doe"})
	if err == nil {
		t.Errorf("CreateUser() returned nil, expected an error")
	}
}

// TestStorageGetUserSuccess tests the successful case of the GetUser method.
func TestStorageGetUserSuccess(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: false}
	ctx := context.Background()
	email := "test@example.com"

	user, err := mockStorage.GetUser(ctx, email)
	if err != nil {
		t.Errorf("GetUser() returned an error: %v, expected nil", err)
	}

	if user == nil {
		t.Errorf("GetUser() returned nil user, expected non-nil")
	}
}

// TestStorageGetUserError tests the error case of the GetUser method.
func TestStorageGetUserError(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: true}
	ctx := context.Background()
	email := "test@example.com"

	user, err := mockStorage.GetUser(ctx, email)
	if err == nil {
		t.Errorf("GetUser() returned nil, expected an error")
	}

	if user != nil {
		t.Errorf("GetUser() returned non-nil user, expected nil")
	}
}

// TestStorageConnectContractsRepositorySuccess tests the successful case of the ConnectContractsRepository method.
func TestStorageConnectContractsRepositorySuccess(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: false}
	mockSQLxDB := &sqlx.DB{}

	err := mockStorage.ConnectContractsRepository(mockSQLxDB)
	if err != nil {
		t.Errorf("ConnectcontractsRepository() returned an error: %v, expected nil", err)
	}
}

// TestStorageConnectContractsRepositoryError tests the error case of the ConnectContractsRepository method.
func TestStorageConnectContractsRepositoryError(t *testing.T) {
	mockStorage := &MockStorage{ShouldError: true}
	mockSQLxDB := &sqlx.DB{}

	err := mockStorage.ConnectContractsRepository(mockSQLxDB)
	if err == nil {
		t.Errorf("ConnectContractsRepository() returned nil, expected an error")
	}
}
