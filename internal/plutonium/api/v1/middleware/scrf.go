package middleware

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"
)

const NonceLen = 16

var GenerateNonceFunc = GenerateNonce // Function variable for mocking

func CsrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrf, err := GenerateNonceFunc()
		if err != nil {
			panic(err)
		}

		expiration := time.Now().Add(time.Hour)
		cookie := http.Cookie{Name: "_csrf", Value: csrf, Expires: expiration, Path: "/"}
		http.SetCookie(w, &cookie)
		next.ServeHTTP(w, r)
	})
}

func GenerateNonce() (string, error) {
	b := make([]byte, NonceLen)
	_, err := rand.Read(b)

	return fmt.Sprintf("%x", b), err
}
