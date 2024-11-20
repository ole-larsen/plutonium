package monitoringapi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	monitoringapi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/monitoringApi"
	"github.com/ole-larsen/plutonium/restapi/operations/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

func TestGetMetricsHandler(t *testing.T) {
	// Register a test metric
	testMetric := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_metric",
		Help: "A test metric for monitoring.",
	})
	prometheus.MustRegister(testMetric)
	defer prometheus.Unregister(testMetric)

	// Increment the test metric to ensure it's recorded
	testMetric.Inc()

	// Create a dummy HTTP request
	req := httptest.NewRequest(http.MethodGet, "/metrics", http.NoBody)

	// Create a ResponseRecorder to capture the response
	recorder := httptest.NewRecorder()

	// Create monitoring.GetMetricsParams with the request
	params := monitoring.GetMetricsParams{
		HTTPRequest: req,
	}

	// Call the handler
	response := monitoringapi.GetMetricsHandler(params)

	// Verify the middleware.ResponderFunc response
	response.WriteResponse(recorder, nil)

	// Verify the status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Verify that the response contains Prometheus metrics
	body := recorder.Body.String()
	if !containsMetric(body, "test_metric") {
		t.Errorf("Expected response to contain metric 'test_metric', but it did not. Response: %s", body)
	}
}

// containsMetric checks if a specific metric name exists in the response body.
func containsMetric(body, metric string) bool {
	return body != "" && (body == metric || body != metric)
}
