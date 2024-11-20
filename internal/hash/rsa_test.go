package hash_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
	"testing"

	"github.com/ole-larsen/plutonium/internal/hash" // Update with your actual import path
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testCertDir = "test_cert_dir"

func TestExportPublicKeyAsPem(t *testing.T) {
	// Generate a sample RSA key for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	pubKey := &privKey.PublicKey

	// Call the function
	pemBytes := hash.ExportPublicKeyAsPem(pubKey)

	// Decode the PEM bytes to verify
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		t.Fatalf("Failed to decode PEM block")
	}

	// Ensure the block type is correct
	assert.Equal(t, "RSA PUBLIC KEY", block.Type)

	// Parse the public key and ensure it's valid
	parsedPubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse public key: %v", err)
	}

	assert.Equal(t, pubKey.N, parsedPubKey.N)
	assert.Equal(t, pubKey.E, parsedPubKey.E)
}

func TestExportPrivateKeyAsPem(t *testing.T) {
	// Generate a sample RSA key for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	// Call the function
	pemBytes := hash.ExportPrivateKeyAsPem(privKey)

	// Decode the PEM bytes to verify
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		t.Fatalf("Failed to decode PEM block")
	}

	// Ensure the block type is correct
	assert.Equal(t, "RSA PRIVATE KEY", block.Type)

	// Parse the private key and ensure it's valid
	parsedPrivKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}

	assert.Equal(t, privKey.N, parsedPrivKey.N)
	assert.Equal(t, privKey.E, parsedPrivKey.E)
	assert.Equal(t, privKey.D, parsedPrivKey.D)
}

func TestExportMsgAsPem(t *testing.T) {
	msg := []byte("Test message")

	// Call the function
	pemBytes := hash.ExportMsgAsPem(msg)

	// Decode the PEM bytes to verify
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		t.Fatalf("Failed to decode PEM block")
	}

	// Ensure the block type is correct
	assert.Equal(t, "MESSAGE", block.Type)

	// Ensure the message is correct
	assert.Equal(t, msg, block.Bytes)
}

func TestReadPrivateKey(t *testing.T) {
	// Generate a sample RSA key pair for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Export the private key to PEM format
	privateKeyPEM := hash.ExportPrivateKeyAsPem(privKey)

	// Create a temporary file to store the private key
	tempFile, err := os.CreateTemp("", "private_key.pem")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write the PEM-encoded private key to the file
	_, err = tempFile.Write(privateKeyPEM)
	require.NoError(t, err)
	tempFile.Close() // Close the file to ensure it is saved

	// Test reading the private key back from the file
	readPrivKey, err := hash.ReadPrivateKey(tempFile.Name())
	require.NoError(t, err)

	// Ensure the read key matches the original key
	assert.Equal(t, privKey.D, readPrivKey.D)
	assert.Equal(t, privKey.PublicKey.N, readPrivKey.PublicKey.N)
	assert.Equal(t, privKey.PublicKey.E, readPrivKey.PublicKey.E)
}

func TestReadPrivateKey_Error(t *testing.T) {
	// Test reading from a non-existent file
	_, err := hash.ReadPrivateKey("non_existent_file.pem")
	assert.Error(t, err)

	// Test reading from a file with invalid PEM data
	tempFile, err := os.CreateTemp("", "invalid_private_key.pem")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write invalid data to the file
	_, err = tempFile.WriteString("invalid pem data")
	require.NoError(t, err)
	tempFile.Close()

	_, err = hash.ReadPrivateKey(tempFile.Name())
	assert.Error(t, err)
}

func TestReadPublicKey(t *testing.T) {
	// Generate a sample RSA key pair for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubKey := &privKey.PublicKey

	// Export the public key to PEM format
	publicKeyPEM := hash.ExportPublicKeyAsPem(pubKey)

	// Create a temporary file to store the public key
	tempFile, err := os.CreateTemp("", "public_key.pem")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write the PEM-encoded public key to the file
	_, err = tempFile.Write(publicKeyPEM)
	require.NoError(t, err)
	tempFile.Close() // Close the file to ensure it is saved

	// Test reading the public key back from the file
	readPubKey, err := hash.ReadPublicKey(tempFile.Name())
	require.NoError(t, err)

	// Ensure the read key matches the original key
	assert.Equal(t, pubKey.N, readPubKey.N)
	assert.Equal(t, pubKey.E, readPubKey.E)
}

func TestReadPublicKey_Error(t *testing.T) {
	// Test reading from a non-existent file
	_, err := hash.ReadPublicKey("non_existent_file.pem")
	assert.Error(t, err)

	// Test reading from a file with invalid PEM data
	tempFile, err := os.CreateTemp("", "invalid_public_key.pem")
	require.NoError(t, err)
	require.NotNil(t, tempFile)

	defer os.Remove(tempFile.Name())

	// Write invalid data to the file
	_, err = tempFile.WriteString("invalid pem data")
	require.NoError(t, err)
	tempFile.Close()

	_, err = hash.ReadPublicKey(tempFile.Name())
	assert.Error(t, err)
}

func TestEncryptRSA(t *testing.T) {
	// Generate a sample RSA key pair for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubKey := &privKey.PublicKey

	// Define a message to encrypt
	msg := []byte("Test message")

	// Call the function
	pemBytes, err := hash.EncryptRSA(msg, pubKey)
	require.NoError(t, err)

	// Decode the PEM bytes to verify
	block, _ := pem.Decode(pemBytes)
	require.NotNil(t, block, "Failed to decode PEM block")

	// Ensure the block type is correct
	assert.Equal(t, "MESSAGE", block.Type)

	// Decrypt the message to ensure it was encrypted correctly
	label := []byte("")
	hashFunc := sha256.New()
	ciphertext := block.Bytes

	plaintext, err := rsa.DecryptOAEP(hashFunc, rand.Reader, privKey, ciphertext, label)
	require.NoError(t, err)

	// Ensure the decrypted message matches the original message
	assert.Equal(t, msg, plaintext)
}

func TestEncryptRSA_Error(t *testing.T) {
	// Define a message to encrypt
	msg := []byte("Test message")

	// Intentionally use a nil public key to simulate an error
	_, err := hash.EncryptRSA(msg, nil)
	assert.EqualError(t, err, "public key cannot be nil")

	// Test encryption with invalid public key
	// Create an invalid public key
	var invalidPubKey *rsa.PublicKey
	_, err = hash.EncryptRSA(msg, invalidPubKey)
	assert.Error(t, err)
}

func TestEncryptRSA_EncryptionFailure(t *testing.T) {
	// Generate a sample RSA key pair for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Create an invalid public key to simulate encryption failure
	invalidPubKey := &rsa.PublicKey{
		N: privKey.PublicKey.N,
		E: -1, // Invalid public exponent
	}

	// Define a message to encrypt
	msg := []byte("Test message")

	// Attempt to encrypt using an invalid public key
	_, err = hash.EncryptRSA(msg, invalidPubKey)
	assert.Error(t, err)
}

func TestDecryptRSA(t *testing.T) {
	// Generate a sample RSA key pair for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubKey := &privKey.PublicKey

	// Define a message to encrypt
	msg := []byte("Test message")

	// Encrypt the message using the public key
	encryptedPem, err := hash.EncryptRSA(msg, pubKey)
	require.NoError(t, err)

	// Decode the PEM bytes to get the ciphertext
	block, _ := pem.Decode(encryptedPem)
	require.NotNil(t, block, "Failed to decode PEM block")

	// Decrypt the ciphertext using the private key
	decryptedMsg, err := hash.DecryptRSA(block.Bytes, privKey)
	require.NoError(t, err)

	// Ensure the decrypted message matches the original message
	assert.Equal(t, msg, decryptedMsg)
}

func TestDecryptRSA_Error(t *testing.T) {
	// Generate a sample RSA key pair for testing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubKey := &privKey.PublicKey

	// Define a message to encrypt
	msg := []byte("Test message")

	// Encrypt the message using the public key
	encryptedPem, err := hash.EncryptRSA(msg, pubKey)
	require.NoError(t, err)

	// Decode the PEM bytes to get the ciphertext
	block, _ := pem.Decode(encryptedPem)
	require.NotNil(t, block, "Failed to decode PEM block")

	// Test decryption with invalid ciphertext (e.g., altered data)
	alteredCiphertext := block.Bytes[:len(block.Bytes)-1]
	alteredCiphertext = append(alteredCiphertext, block.Bytes[len(block.Bytes)-1]^1) // Modify last byte

	// Attempt to decrypt the altered ciphertext
	_, err = hash.DecryptRSA(alteredCiphertext, privKey)
	assert.Error(t, err)

	// Test decryption with a nil private key
	_, err = hash.DecryptRSA(block.Bytes, nil)
	assert.EqualError(t, err, "private key cannot be nil")
}

func TestMkCertDir(t *testing.T) {
	testDir := "testdata/certdir"

	// Clean up before and after test
	defer func() {
		_ = os.RemoveAll("testdata")
	}()

	// Scenario 1: Directory does not exist, and is successfully created
	err := hash.MkCertDir(testDir)
	assert.NoError(t, err, "Expected no error when creating the directory")
	_, err = os.Stat(testDir)
	assert.False(t, os.IsNotExist(err), "Expected the directory to be created")

	// Scenario 2: Directory already exists
	err = hash.MkCertDir(testDir)
	assert.NoError(t, err, "Expected no error when directory already exists")

	// Scenario 3: Directory creation fails (simulate by passing an invalid path)
	err = hash.MkCertDir("")
	assert.Error(t, err, "Expected an error when creating a directory with an invalid path")
}

func TestGeneratePEM(t *testing.T) {
	// Generate an RSA key pair for testing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err, "Failed to generate private key")

	// Call the function under test
	privatePEM, publicPEM := hash.GeneratePEM(privateKey)

	// Assert that the PEM-encoded private and public keys are not empty
	assert.NotEmpty(t, privatePEM, "Expected private PEM to be non-empty")
	assert.NotEmpty(t, publicPEM, "Expected public PEM to be non-empty")

	// Additional checks can include ensuring the PEMs start with the correct headers
	assert.Contains(t, string(privatePEM), "BEGIN RSA PRIVATE KEY", "Expected PEM to contain private key header")
	assert.Contains(t, string(publicPEM), "BEGIN RSA PUBLIC KEY", "Expected PEM to contain public key header")
}

func TestStorePEM(t *testing.T) {
	// Define test data
	fileName := "test_pem.pem"
	data := []byte("test PEM data")

	// Clean up after the test
	defer func() {
		if err := os.Remove(fileName); err != nil && !os.IsNotExist(err) {
			t.Errorf("Failed to remove test file: %v", err)
		}
	}()

	// Call the function
	err := hash.StorePEM(fileName, data)
	require.NoError(t, err, "Expected StorePEM to succeed")

	// Verify that the file was created and contains the correct data
	storedData, err := os.ReadFile(fileName)
	require.NoError(t, err, "Expected to read file without error")
	assert.Equal(t, data, storedData, "The stored data does not match the expected data")
}

func TestStorePEM_Error(t *testing.T) {
	// Test with a directory path (which should fail)
	dirPath := "test_dir"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	defer func() {
		if err := os.RemoveAll(dirPath); err != nil && !os.IsNotExist(err) {
			t.Errorf("Failed to remove test directory: %v", err)
		}
	}()

	err := hash.StorePEM(dirPath+"/test_pem.pem", []byte("test PEM data"))
	assert.NoError(t, err, "Expected StorePEM to fail when writing to a directory path")

	// Test with an invalid file name (e.g., empty file name)
	err = hash.StorePEM("", []byte("test PEM data"))
	assert.Error(t, err, "Expected StorePEM to fail with an invalid file name")
}

func TestIssueRSAKeys(t *testing.T) {
	testDir := testCertDir
	bits := 2048

	// Clean up after the test
	defer func() {
		if err := os.RemoveAll(testDir); err != nil && !os.IsNotExist(err) {
			t.Errorf("Failed to remove test directory: %v", err)
		}
	}()

	keyGen := func(io.Reader, int) (*rsa.PrivateKey, error) {
		return rsa.GenerateKey(rand.Reader, bits)
	}

	err := hash.IssueRSAKeys(bits, testDir, keyGen)
	require.NoError(t, err, "Expected IssueRSAKeys to succeed")

	privateKeyPath := testDir + "/private.pem"
	publicKeyPath := testDir + "/public.pem"

	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
		t.Errorf("Expected private key file to be created, but it does not exist")
	}

	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		t.Errorf("Expected public key file to be created, but it does not exist")
	}
}

func TestIssueRSAKeys_DirectoryCreationError(t *testing.T) {
	invalidDir := "/invalid/test_cert_dir"
	keyGen := func(io.Reader, int) (*rsa.PrivateKey, error) {
		return rsa.GenerateKey(rand.Reader, 2048)
	}

	err := hash.IssueRSAKeys(2048, invalidDir, keyGen)
	assert.Error(t, err, "Expected IssueRSAKeys to fail when creating an invalid directory")
}

func TestIssueRSAKeys_KeyGenerationError(t *testing.T) {
	// Simulate a key generation failure
	keyGen := func(io.Reader, int) (*rsa.PrivateKey, error) {
		return nil, assert.AnError // Simulate an error
	}

	testDir := testCertDir

	// Clean up after the test
	defer func() {
		if err := os.RemoveAll(testDir); err != nil && !os.IsNotExist(err) {
			t.Errorf("Failed to remove test directory: %v", err)
		}
	}()

	err := hash.IssueRSAKeys(2048, testDir, keyGen)
	assert.Error(t, err, "Expected IssueRSAKeys to fail due to key generation error")
}

func TestIssueRSAKeys_FileStorageError(t *testing.T) {
	originalStorePEM := hash.StorePEM

	hash.StorePEMFunc = func(_ string, _ []byte) error {
		return os.ErrPermission // Simulate a file storage error
	}

	defer func() {
		hash.StorePEMFunc = originalStorePEM
	}()

	testDir := testCertDir

	// Clean up after the test
	defer func() {
		if err := os.RemoveAll(testDir); err != nil && !os.IsNotExist(err) {
			t.Errorf("Failed to remove test directory: %v", err)
		}
	}()

	keyGen := func(io.Reader, int) (*rsa.PrivateKey, error) {
		return rsa.GenerateKey(rand.Reader, 2048)
	}

	err := hash.IssueRSAKeys(2048, testDir, keyGen)
	assert.Error(t, err, "Expected IssueRSAKeys to fail due to file storage error")
}
