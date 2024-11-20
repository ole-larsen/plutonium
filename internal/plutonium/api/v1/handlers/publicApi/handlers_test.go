package publicapi_test

import (
	"testing"
	"time"

	publicapi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/publicApi"
	"github.com/ole-larsen/plutonium/restapi/operations/public"
)

func TestGetPingHandler(t *testing.T) {
	// Call the handler
	response := publicapi.GetPingHandler(public.GetPingParams{})

	// Verify the response type
	responder, ok := response.(*public.GetPingOK)
	if !ok {
		t.Fatalf("Expected response of type *public.GetPingOK, got %T", response)
	}

	// Verify the payload
	payload := responder.Payload
	if payload == nil {
		t.Fatal("Expected payload to be non-nil")
	}

	// Check the Message field
	expectedMessage := "pong"
	if payload.Message != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, payload.Message)
	}

	// Check the Timestamp field
	if payload.Timestamp.IsZero() {
		t.Error("Expected Timestamp to be a non-zero value")
	}

	// Verify the timestamp is close to now
	now := time.Now()

	timestamp, err := time.Parse(time.RFC3339, payload.Timestamp.String())
	if err != nil {
		t.Fatalf("Failed to parse timestamp: %v", err)
	}

	if timestamp.Before(now.Add(-1*time.Second)) || timestamp.After(now.Add(1*time.Second)) {
		t.Errorf("Expected timestamp to be close to now, got %v", timestamp)
	}
}
