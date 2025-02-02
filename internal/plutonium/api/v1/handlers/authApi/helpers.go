package authapi

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/ole-larsen/plutonium/internal/hash"
)

func generateNonce() string {
	id := uuid.New()
	return id.String()
}

func gravatar(email string, size int) string {
	gravatarURL := "https://gravatar.com/avatar/"
	if email != "" {
		return gravatarURL + hash.GetMD5Hash(email) + "?s=" + strconv.Itoa(size) + "&d=retro"
	}

	return gravatarURL + "?s=" + strconv.Itoa(size) + "&d=retro"
}
