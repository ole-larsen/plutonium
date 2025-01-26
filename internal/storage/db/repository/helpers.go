package repository

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/ole-larsen/plutonium/internal/hash"
)

const nonceLength = 16

func GenerateNonce() (string, error) {
	b := make([]byte, nonceLength)
	_, err := rand.Read(b)

	return fmt.Sprintf("%x", b), err
}

func SetPassword(password interface{}) (string, error) {
	var pwd string

	if plainPwd, ok := password.(string); ok {
		if plainPwd == "" {
			return pwd, errors.New("empty password not allowed")
		}

		return hash.Password([]byte(plainPwd))
	}

	return pwd, errors.New("password must be a string")
}
