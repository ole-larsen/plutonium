package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ole-larsen/plutonium/internal/plutonium/api/v1/middleware"
)

func TestCsrfMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Next handler called"))
	})

	handler := middleware.CsrfMiddleware(nextHandler)

	t.Run("CSRF cookie is set", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Check for the CSRF cookie
		cookies := resp.Cookies()
		if len(cookies) == 0 {
			t.Fatalf("Expected a CSRF cookie to be set, but found none")
		}

		csrfCookie := cookies[0]
		if csrfCookie.Name != "_csrf" {
			t.Errorf("Expected cookie name '_csrf', got '%s'", csrfCookie.Name)
		}

		if csrfCookie.Value == "" {
			t.Error("Expected a non-empty CSRF token value")
		}

		if csrfCookie.Path != "/" {
			t.Errorf("Expected cookie path '/', got '%s'", csrfCookie.Path)
		}

		if time.Until(csrfCookie.Expires) <= 0 {
			t.Error("Expected cookie expiration to be in the future")
		}
	})

	t.Run("Next handler is called", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		body := rec.Body.String()
		if !strings.Contains(body, "Next handler called") {
			t.Error("Next handler was not called")
		}
	})
}

func TestGenerateNonce(t *testing.T) {
	t.Run("Generates valid nonce", func(t *testing.T) {
		nonce, err := middleware.GenerateNonce()
		if err != nil {
			t.Fatalf("Unexpected error generating nonce: %v", err)
		}

		if len(nonce) != middleware.NonceLen*2 { // Hex encoding doubles the length
			t.Errorf("Expected nonce length %d, got %d", middleware.NonceLen*2, len(nonce))
		}
	})
}

func TestCsrfMiddlewarePanic(t *testing.T) {
	// Save the original function and restore it after the test
	originalGenerateNonce := middleware.GenerateNonceFunc
	defer func() { middleware.GenerateNonceFunc = originalGenerateNonce }()

	// Mock the function to return an error
	middleware.GenerateNonceFunc = func() (string, error) {
		return "", fmt.Errorf("simulated error")
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Next handler called"))
	})

	middlewareHandler := middleware.CsrfMiddleware(nextHandler)

	// Use defer-recover to assert that panic occurs
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but none occurred")
		}
	}()

	// This should trigger a panic
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	middlewareHandler.ServeHTTP(rec, req)
}
