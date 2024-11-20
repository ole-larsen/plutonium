package prometheus_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ole-larsen/plutonium/internal/plutonium/api/v1/prometheus"
)

func TestNewPrometheusResponseWriter(t *testing.T) {
	// Create a dummy ResponseWriter
	dummyWriter := httptest.NewRecorder()

	// Create a PrometheusResponseWriter
	prometheusWriter := prometheus.NewResponseWriter(dummyWriter)

	// Verify the wrapped writer
	if prometheusWriter.ResponseWriter != dummyWriter {
		t.Errorf("Expected ResponseWriter to be %v, got %v", dummyWriter, prometheusWriter.ResponseWriter)
	}

	// Verify the default StatusCode
	if prometheusWriter.StatusCode != http.StatusOK {
		t.Errorf("Expected default StatusCode to be %d, got %d", http.StatusOK, prometheusWriter.StatusCode)
	}
}

func TestPrometheusResponseWriter_WriteHeader(t *testing.T) {
	// Create a dummy ResponseWriter
	dummyWriter := httptest.NewRecorder()

	// Create a PrometheusResponseWriter
	prometheusWriter := prometheus.NewResponseWriter(dummyWriter)

	// Write a custom status code
	customStatusCode := http.StatusNotFound
	prometheusWriter.WriteHeader(customStatusCode)

	// Access and defer closing of the ResponseRecorder's result body
	response := dummyWriter.Result()
	defer response.Body.Close()

	// Verify the StatusCode is updated
	if prometheusWriter.StatusCode != customStatusCode {
		t.Errorf("Expected StatusCode to be %d, got %d", customStatusCode, prometheusWriter.StatusCode)
	}

	// Verify the dummy ResponseWriter received the correct status code
	if response.StatusCode != customStatusCode {
		t.Errorf("Expected wrapped ResponseWriter StatusCode to be %d, got %d", customStatusCode, response.StatusCode)
	}
}
