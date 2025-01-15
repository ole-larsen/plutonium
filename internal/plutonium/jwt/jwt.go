package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token string `json:"token"`
}

func Sign(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateToken(login, secret string) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("could not sign token")
	}

	claims := jwt.MapClaims{
		"login":     login,
		"timestamp": time.Now(),
	}

	return Sign(claims, secret)
}

func Verify(token, secret string) (map[string]interface{}, error) {
	jwToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return jwToken.Claims.(jwt.MapClaims), nil
}

func GetBearerToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("an authorization header is required")
	}

	token := strings.Split(header, " ")
	tknLength := 2

	if len(token) != tknLength {
		return "", fmt.Errorf("malformed bearer token")
	}

	return token[1], nil
}
