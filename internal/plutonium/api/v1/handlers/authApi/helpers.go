package authapi

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password, secret string) (string, error) {
	hashed := password + secret
	bytes, err := bcrypt.GenerateFromPassword([]byte(hashed), 14)
	return string(bytes), err
}

func generateNonce() string {
	id := uuid.New()
	return id.String()
}
