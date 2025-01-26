package authapi

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const cost = 14

func createHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hashPassword(password, secret string) (string, error) {
	hashed := password + secret
	bytes, err := bcrypt.GenerateFromPassword([]byte(hashed), cost)

	return string(bytes), err
}

func generateNonce() string {
	id := uuid.New()
	return id.String()
}

func gravatar(email string, size int) string {
	gravatarURL := "https://gravatar.com/avatar/"
	if email != "" {
		return gravatarURL + createHash(email) + "?s=" + strconv.Itoa(size) + "&d=retro"
	}

	return gravatarURL + "?s=" + strconv.Itoa(size) + "&d=retro"
}
