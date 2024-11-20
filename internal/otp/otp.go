package otp

import (
	"crypto/rand"
	"encoding/base32"

	"github.com/dgryski/dgoogauth"
)

func CreateOTPSecret(readRandom func(b []byte) (n int, err error)) string {
	const length = 10
	random := make([]byte, length)

	_, err := readRandom(random)
	if err != nil {
		panic(err)
	}

	return base32.StdEncoding.EncodeToString(random)
}

func CreateOTP() *dgoogauth.OTPConfig {
	return &dgoogauth.OTPConfig{
		Secret:      CreateOTPSecret(rand.Read),
		WindowSize:  3,
		HotpCounter: 0,
	}
}

func GetOTP(secret string) *dgoogauth.OTPConfig {
	return &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}
}
