package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ole-larsen/plutonium/internal/plutonium/api/v1/middleware"
	"github.com/stretchr/testify/require"
)

func TestCorsMiddleware(t *testing.T) {
	handler := middleware.CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		require.NoError(t, err)
	}))

	tests := []struct {
		headers         map[string]string
		expectedHeaders map[string]string
		name            string
		method          string
		body            string
		expectedStatus  int
	}{
		{
			name:   "OPTIONS request",
			method: http.MethodOptions,
			headers: map[string]string{
				"Origin": "http://example.com",
			},
			expectedStatus: http.StatusOK,
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "http://example.com",
				"Content-Type":                     "application/json, multipart/form-data, application/x-www-form-urlencoded",
				"Access-Control-Allow-Methods":     "POST, GET, OPTIONS, PUT, PATCH, DELETE",
				"Access-Control-Allow-Headers":     "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Token",
				"Access-Control-Expose-Headers":    "*",
				"Access-Control-Allow-Credentials": "true",
			},
			body: "",
		},
		{
			name:   "GET request",
			method: http.MethodGet,
			headers: map[string]string{
				"Origin": "http://example.com",
			},
			expectedStatus: http.StatusOK,
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "http://example.com",
				"Content-Type":                     "application/json, multipart/form-data, application/x-www-form-urlencoded",
				"Access-Control-Allow-Methods":     "POST, GET, OPTIONS, PUT, PATCH, DELETE",
				"Access-Control-Allow-Headers":     "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Token",
				"Access-Control-Expose-Headers":    "*",
				"Access-Control-Allow-Credentials": "true",
			},
			body: "OK",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "http://example.com", http.NoBody)
			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			resp := rec.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			for k, v := range tt.expectedHeaders {
				if got := resp.Header.Get(k); got != v {
					t.Errorf("expected header %s to be %q, got %q", k, v, got)
				}
			}

			body := rec.Body.String()
			if strings.TrimSpace(body) != tt.body {
				t.Errorf("expected body %q, got %q", tt.body, body)
			}
		})
	}
}
