// Package hash_test provides tests for the hash package.
package hash_test

import (
	"bytes"
	"fmt"
	"testing"

	"crypto/cipher"

	"github.com/ole-larsen/plutonium/internal/hash"
	"github.com/stretchr/testify/mock"
)

// Mock for GenerateRandomFunc.
type MockGenerateRandom struct {
	mock.Mock
}

func (m *MockGenerateRandom) GenerateRandom(size int) ([]byte, error) {
	args := m.Called(size)

	randomBytes, ok := args.Get(0).([]byte)
	if !ok && args.Get(0) != nil {
		return nil, fmt.Errorf("invalid type assertion for []byte")
	}

	return randomBytes, args.Error(1)
}

// Mock for NewCipherFunc.
type MockCipher struct {
	mock.Mock
}

const secretErrMsg = "secret key cannot be nil"
const secretNilMsg = "secret key cannot be empty"

const errNilMsg = "message cannot be nil"
const errEmptyMsg = "message cannot be empty"

func (m *MockCipher) NewCipher(key []byte) (cipher.Block, error) {
	args := m.Called(key)

	block, ok := args.Get(0).(cipher.Block)
	if !ok && args.Get(0) != nil {
		return nil, fmt.Errorf("invalid type assertion for cipher.Block")
	}

	return block, args.Error(1)
}

// Mock for NewGCMFunc.
type MockGCM struct {
	mock.Mock
}

func (m *MockGCM) NewGCM(block cipher.Block) (cipher.AEAD, error) {
	args := m.Called(block)

	aead, ok := args.Get(0).(cipher.AEAD)
	if !ok && args.Get(0) != nil {
		return nil, fmt.Errorf("invalid type assertion for cipher.AEAD")
	}

	return aead, args.Error(1)
}

// MockAEAD is a mock implementation of cipher.AEAD for testing purposes.
type MockAEAD struct {
	mock.Mock
}

func (m *MockAEAD) NonceSize() int {
	return 12
}

func (m *MockAEAD) Overhead() int {
	return 16
}

func (m *MockAEAD) Seal(_, _, _, _ []byte) []byte {
	return nil // Not needed for this test
}

func (m *MockAEAD) Open(_, _, _, _ []byte) ([]byte, error) {
	return nil, fmt.Errorf("some aesgcm.Open error") // Force an error
}

func TestGenerateRandom(t *testing.T) {
	t.Run("valid size", func(t *testing.T) {
		size := 16

		data, err := hash.GenerateRandom(size)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(data) != size {
			t.Fatalf("expected length %d, got %d", size, len(data))
		}
	})

	t.Run("invalid size", func(t *testing.T) {
		size := -1

		_, err := hash.GenerateRandom(size)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := fmt.Sprintf("size must be a positive integer, got %d", size)
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}

func TestCreate32BytesKey(t *testing.T) {
	t.Run("valid secret", func(t *testing.T) {
		secret := []byte("test-secret")

		key, err := hash.Create32BytesKey(secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(key) != 32 {
			t.Fatalf("expected key length 32, got %d", len(key))
		}
	})

	t.Run("nil secret", func(t *testing.T) {
		_, err := hash.Create32BytesKey(nil)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := secretErrMsg
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}

func TestEncrypt(t *testing.T) {
	t.Run("successful encryption", func(t *testing.T) {
		msg := []byte("message")
		secret := []byte("super-secret-key")

		encrypted, err := hash.Encrypt(msg, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(encrypted) <= len(msg) {
			t.Fatalf("encrypted message should be longer than original message")
		}
	})

	t.Run("nil message", func(t *testing.T) {
		secret := []byte("super-secret-key")

		_, err := hash.Encrypt(nil, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := errNilMsg
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("empty message", func(t *testing.T) {
		secret := []byte("super-secret-key")

		_, err := hash.Encrypt([]byte(""), secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := errEmptyMsg
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("nil secret", func(t *testing.T) {
		msg := []byte("message")

		_, err := hash.Encrypt(msg, nil)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := secretErrMsg
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("empty secret", func(t *testing.T) {
		msg := []byte("message")

		_, err := hash.Encrypt(msg, []byte(""))
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := secretNilMsg
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("CreateKeyFunc error", func(t *testing.T) {
		originalCreateKeyFunc := hash.CreateKeyFunc
		defer func() { hash.CreateKeyFunc = originalCreateKeyFunc }()

		hash.CreateKeyFunc = func(_ []byte) ([]byte, error) {
			return nil, fmt.Errorf("create key error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Encrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "create key error" {
			t.Fatalf("expected error 'create key error', got %v", err)
		}
	})

	t.Run("NewCipherFunc error", func(t *testing.T) {
		originalNewCipherFunc := hash.NewCipherFunc
		defer func() { hash.NewCipherFunc = originalNewCipherFunc }()

		hash.NewCipherFunc = func(_ []byte) (cipher.Block, error) {
			return nil, fmt.Errorf("new cipher error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Encrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "new cipher error" {
			t.Fatalf("expected error 'new cipher error', got %v", err)
		}
	})

	t.Run("NewGCMFunc error", func(t *testing.T) {
		originalNewGCMFunc := hash.NewGCMFunc
		defer func() { hash.NewGCMFunc = originalNewGCMFunc }()

		hash.NewGCMFunc = func(_ cipher.Block) (cipher.AEAD, error) {
			return nil, fmt.Errorf("new gcm error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Encrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "new gcm error" {
			t.Fatalf("expected error 'new gcm error', got %v", err)
		}
	})

	t.Run("GenerateRandomFunc error", func(t *testing.T) {
		originalGenerateRandomFunc := hash.GenerateRandomFunc
		defer func() { hash.GenerateRandomFunc = originalGenerateRandomFunc }()

		hash.GenerateRandomFunc = func(_ int) ([]byte, error) {
			return nil, fmt.Errorf("generate random error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Encrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "generate random error" {
			t.Fatalf("expected error 'generate random error', got %v", err)
		}
	})
}

func TestDecrypt(t *testing.T) {
	t.Run("successful decryption", func(t *testing.T) {
		msg := []byte("message")
		secret := []byte("super-secret-key")

		encrypted, err := hash.Encrypt(msg, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		decrypted, err := hash.Decrypt(encrypted, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !bytes.Equal(decrypted, msg) {
			t.Fatalf("expected decrypted message %s, got %s", msg, decrypted)
		}
	})

	t.Run("nil message", func(t *testing.T) {
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(nil, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := "nothing to decrypt"
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("nil secret", func(t *testing.T) {
		msg := []byte("message")

		_, err := hash.Decrypt(msg, nil)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := secretErrMsg
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("message too short", func(t *testing.T) {
		msg := []byte("short")
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := "message too short"
		if err.Error() != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("CreateKeyFunc error", func(t *testing.T) {
		originalCreateKeyFunc := hash.CreateKeyFunc
		defer func() { hash.CreateKeyFunc = originalCreateKeyFunc }()

		hash.CreateKeyFunc = func(_ []byte) ([]byte, error) {
			return nil, fmt.Errorf("create key error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "create key error" {
			t.Fatalf("expected error 'create key error', got %v", err)
		}
	})

	t.Run("NewCipherFunc error", func(t *testing.T) {
		originalNewCipherFunc := hash.NewCipherFunc
		defer func() { hash.NewCipherFunc = originalNewCipherFunc }()

		hash.NewCipherFunc = func(_ []byte) (cipher.Block, error) {
			return nil, fmt.Errorf("new cipher error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "new cipher error" {
			t.Fatalf("expected error 'new cipher error', got %v", err)
		}
	})

	t.Run("NewGCMFunc error", func(t *testing.T) {
		originalNewGCMFunc := hash.NewGCMFunc
		defer func() { hash.NewGCMFunc = originalNewGCMFunc }()

		hash.NewGCMFunc = func(_ cipher.Block) (cipher.AEAD, error) {
			return nil, fmt.Errorf("new gcm error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "new gcm error" {
			t.Fatalf("expected error 'new gcm error', got %v", err)
		}
	})

	t.Run("GenerateRandomFunc error", func(t *testing.T) {
		originalGenerateRandomFunc := hash.GenerateRandomFunc
		defer func() { hash.GenerateRandomFunc = originalGenerateRandomFunc }()

		hash.GenerateRandomFunc = func(_ int) ([]byte, error) {
			return nil, fmt.Errorf("generate random error")
		}

		msg := []byte("message")
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("Decrypt aesgcm.Open error", func(t *testing.T) {
		// Backup the original function and restore it after the test
		originalNewGCMFunc := hash.NewGCMFunc
		defer func() { hash.NewGCMFunc = originalNewGCMFunc }()

		// Set NewGCMFunc to return a MockAEAD that will return an error
		hash.NewGCMFunc = func(_ cipher.Block) (cipher.AEAD, error) {
			return &MockAEAD{}, nil
		}

		// Encrypted message with nonce and payload (valid but will fail on decryption)
		msg := []byte("some-encrypted-message")
		secret := []byte("super-secret-key")

		_, err := hash.Decrypt(msg, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expectedErr := "some aesgcm.Open error" // Ensure this matches the error returned by MockAEAD
		if err.Error() != expectedErr {
			t.Fatalf("expected error '%v', got '%v'", expectedErr, err)
		}
	})
}
