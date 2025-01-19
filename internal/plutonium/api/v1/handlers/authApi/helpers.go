package authapi

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const cost = 14

func hashPassword(password, secret string) (string, error) {
	hashed := password + secret
	bytes, err := bcrypt.GenerateFromPassword([]byte(hashed), cost)

	return string(bytes), err
}

func generateNonce() string {
	id := uuid.New()
	return id.String()
}
