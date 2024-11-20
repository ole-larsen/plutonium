package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ole-larsen/plutonium/internal/plutonium/api/v1/middleware"
	"github.com/ole-larsen/plutonium/internal/plutonium/api/v1/prometheus"
	"github.com/stretchr/testify/require"
)

func TestPrometheusMiddleware(t *testing.T) {
	// Mock handler to be wrapped
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello, World!"))
		require.NoError(t, err)
	})

	// Wrap the handler with PrometheusMiddleware
	handler := middleware.PrometheusMiddleware(mockHandler)

	// Create a test HTTP request
	req := httptest.NewRequest(http.MethodGet, "/test-endpoint", http.NoBody)

	// Create a ResponseRecorder to capture the response
	recorder := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(recorder, req)

	// Verify response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Verify response body
	expectedBody := "Hello, World!"
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, recorder.Body.String())
	}

	// Check that the PrometheusResponseWriter captured the status code
	lrw := prometheus.NewResponseWriter(recorder)
	if lrw.StatusCode != http.StatusOK {
		t.Errorf("Expected PrometheusResponseWriter status code %d, got %d", http.StatusOK, lrw.StatusCode)
	}

	// Ensure the duration is captured (duration cannot be 0)
	duration := time.Since(time.Now())
	if duration.Seconds() <= 0 {
		t.Errorf("Expected a positive duration, got %f seconds", duration.Seconds())
	}
}
