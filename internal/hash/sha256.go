// Package hash implements hash for requests using AES.
// Copyright 2024 The Oleg Nazarov. All rights reserved.
// Package hash implements hash for requests using AES.
// Copyright 2024 The Oleg Nazarov. All rights reserved.
package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

var (
	CreateKeyFunc      = Create32BytesKey
	NewCipherFunc      = aes.NewCipher
	NewGCMFunc         = cipher.NewGCM
	GenerateRandomFunc = GenerateRandom
)

func GenerateRandom(size int) ([]byte, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size must be a positive integer, got %d", size)
	}

	b := make([]byte, size)
	_, err := rand.Read(b)

	return b, err
}

func Create32BytesKey(secret []byte) ([]byte, error) {
	if secret == nil {
		return nil, fmt.Errorf("secret key cannot be nil")
	}

	h := sha256.New()
	_, err := h.Write(secret)

	return h.Sum(nil), err
}

func Encrypt(msg, secret []byte) ([]byte, error) {
	if msg == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}

	if len(msg) == 0 {
		return nil, fmt.Errorf("message cannot be empty")
	}

	if secret == nil {
		return nil, fmt.Errorf("secret key cannot be nil")
	}

	if len(secret) == 0 {
		return nil, fmt.Errorf("secret key cannot be empty")
	}

	key, err := CreateKeyFunc(secret)
	if err != nil {
		return nil, err
	}

	aesblock, err := NewCipherFunc(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := NewGCMFunc(aesblock)
	if err != nil {
		return nil, err
	}

	nonce, err := GenerateRandomFunc(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nonce, nonce, msg, nil), nil
}

func Decrypt(msg, secret []byte) ([]byte, error) {
	if msg == nil {
		return nil, fmt.Errorf("nothing to decrypt")
	}

	if secret == nil {
		return nil, fmt.Errorf("secret key cannot be nil")
	}

	// Create key from secret, sized 32 bytes
	key, err := CreateKeyFunc(secret)
	if err != nil {
		return nil, err
	}

	// Create cipher.Block with 32 bytes key to get AES-256
	aesblock, err := NewCipherFunc(key)
	if err != nil {
		return nil, err
	}

	// NewGCM returns the specified 128-bit block cipher
	aesgcm, err := NewGCMFunc(aesblock)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(msg) < nonceSize {
		return nil, fmt.Errorf("message too short")
	}

	nonce, msg := msg[:nonceSize], msg[nonceSize:]

	decrypted, err := aesgcm.Open(nil, nonce, msg, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
