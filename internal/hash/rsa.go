// Package hash provides functions for RSA key generation, encryption, decryption,
// and PEM encoding/decoding, as well as functions for file operations related
// to certificate storage.

// The package includes functions to generate RSA keys, export keys to PEM format,
// import keys from PEM format, encrypt and decrypt messages using RSA, and handle
// file operations for storing and reading keys.
package hash

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

const perm = 0o600

var (
	StorePEMFunc = StorePEM
)

// ExportPublicKeyAsPem encodes an RSA public key into PEM format.
// The PEM block will have the type "RSA PUBLIC KEY".
//
// Parameters:
// - pubkey: The RSA public key to be encoded.
//
// Returns:
// - []byte: The PEM-encoded public key.
func ExportPublicKeyAsPem(pubkey *rsa.PublicKey) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(pubkey)})
}

// ExportPrivateKeyAsPem encodes an RSA private key into PEM format.
// The PEM block will have the type "RSA PRIVATE KEY".
//
// Parameters:
// - privatekey: The RSA private key to be encoded.
//
// Returns:
// - []byte: The PEM-encoded private key.
func ExportPrivateKeyAsPem(privatekey *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)})
}

// ExportMsgAsPem encodes a message into PEM format.
// The PEM block will have the type "MESSAGE".
//
// Parameters:
// - msg: The message to be encoded.
//
// Returns:
// - []byte: The PEM-encoded message.
func ExportMsgAsPem(msg []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "MESSAGE", Bytes: msg})
}

// Type definition for the key generation function rsa.GenerateKey(rand.Reader, bits).
type KeyGenerator func(io.Reader, int) (*rsa.PrivateKey, error)

// IssueRSAKeys generates RSA private and public keys and stores them as PEM files.
// If the directory specified by 'root' does not exist, it is created.
//
// Parameters:
// - bits: The number of bits for the RSA key (e.g., 2048, 4096).
// - root: The directory where the PEM files will be stored.
//
// Returns:
// - error: An error if key generation or file writing fails.
func IssueRSAKeys(bits int, root string, keyGen KeyGenerator) error {
	if err := MkCertDir(root); err != nil {
		return err
	}

	privateKey, err := keyGen(rand.Reader, bits)
	if err != nil {
		return err
	}

	privateKeyPEM, publicKeyPEM := GeneratePEM(privateKey)

	var errs []error

	if err := StorePEMFunc(root+"/private.pem", privateKeyPEM); err != nil {
		errs = append(errs, err)
	}

	if err := StorePEMFunc(root+"/public.pem", publicKeyPEM); err != nil {
		errs = append(errs, err)
	}

	err = errors.Join(errs...)

	if err != nil {
		return err
	}

	return nil
}

// ReadPrivateKey reads an RSA private key from a PEM-encoded file.
//
// Parameters:
// - path: The path to the PEM file containing the private key.
//
// Returns:
// - *rsa.PrivateKey: The parsed RSA private key.
// - error: An error if reading the file or parsing the key fails.
func ReadPrivateKey(path string) (*rsa.PrivateKey, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(file)

	if block == nil {
		return nil, errors.New("invalid PEM block")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// ReadPublicKey reads an RSA public key from a PEM-encoded file.
//
// Parameters:
// - path: The path to the PEM file containing the public key.
//
// Returns:
// - *rsa.PublicKey: The parsed RSA public key.
// - error: An error if reading the file or parsing the key fails.
func ReadPublicKey(path string) (*rsa.PublicKey, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(file)
	if block == nil {
		return nil, errors.New("invalid PEM block")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// EncryptRSA encrypts a message using RSA and encodes it into PEM format.
//
// Parameters:
// - msg: The message to be encrypted.
// - publicKey: The RSA public key used for encryption.
//
// Returns:
// - []byte: The PEM-encoded ciphertext.
// - error: An error if encryption fails.
func EncryptRSA(msg []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("public key cannot be nil")
	}

	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, msg, label)
	if err != nil {
		return nil, err
	}

	return ExportMsgAsPem(ciphertext), nil
}

// DecryptRSA decrypts a PEM-encoded ciphertext using RSA.
//
// Parameters:
// - ciphertext: The PEM-encoded ciphertext to be decrypted.
// - privateKey: The RSA private key used for decryption.
//
// Returns:
// - []byte: The decrypted message.
// - error: An error if decryption fails.
func DecryptRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}

	label := []byte("")
	hash := sha256.New()

	return rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
}

// MkCertDir - Creates directory for certs if not exists.
func MkCertDir(root string) error {
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(root, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func GeneratePEM(privateKey *rsa.PrivateKey) (priv, pub []byte) {
	publicKey := &privateKey.PublicKey

	return ExportPrivateKeyAsPem(privateKey), ExportPublicKeyAsPem(publicKey)
}

func StorePEM(name string, data []byte) error {
	return os.WriteFile(name, data, perm)
}
