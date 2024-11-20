package otp_test

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"testing"

	"github.com/ole-larsen/plutonium/internal/otp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOTPSecret(t *testing.T) {
	// Generate a new OTP secret
	secret := otp.CreateOTPSecret(rand.Read)

	// Ensure the secret is a valid base32-encoded string
	_, err := base32.StdEncoding.DecodeString(secret)
	require.NoError(t, err, "The OTP secret should be a valid base32-encoded string")

	// Optionally check the length of the secret
	assert.True(t, secret != "", "The OTP secret should not be empty")
}

func TestCreateOTP(t *testing.T) {
	// Create a new OTP configuration
	otpConfig := otp.CreateOTP()

	// Ensure that the OTPConfig is created with a valid secret
	_, err := base32.StdEncoding.DecodeString(otpConfig.Secret)
	require.NoError(t, err, "The OTP secret should be a valid base32-encoded string")

	// Check other fields to ensure they are set correctly
	assert.Equal(t, 3, otpConfig.WindowSize, "The WindowSize should be set to 3")
	assert.Equal(t, 0, otpConfig.HotpCounter, "The HotpCounter should be set to 0")
}

func TestGetOTP(t *testing.T) {
	// Create a new OTP secret
	secret := otp.CreateOTPSecret(rand.Read)

	// Get the OTP configuration using the created secret
	otpConfig := otp.GetOTP(secret)

	// Ensure that the OTPConfig is created with the provided secret
	assert.Equal(t, secret, otpConfig.Secret, "The OTPConfig secret should match the provided secret")

	// Check other fields to ensure they are set correctly
	assert.Equal(t, 3, otpConfig.WindowSize, "The WindowSize should be set to 3")
	assert.Equal(t, 0, otpConfig.HotpCounter, "The HotpCounter should be set to 0")
}

// MockRead simulates an error during rand.Read.
func MockRead(_ []byte) (int, error) {
	return 0, fmt.Errorf("mock error")
}

func TestCreateOTPSecret_Panic(t *testing.T) {
	// Use require.Panics to assert that the function panics
	require.Panics(t, func() {
		otp.CreateOTPSecret(MockRead)
	}, "The function did not panic as expected")
}

func TestCreateOTPSecret_Success(t *testing.T) {
	secret := otp.CreateOTPSecret(rand.Read)

	// Assert that the secret is not empty and has the expected length
	assert.NotEmpty(t, secret)
	assert.Equal(t, 16, len(secret), "Expected a 16-character OTP secret") // base32 encoding 10 bytes yields 16 characters
}
