package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bazel_query_example/lib/logging"
	"bazel_query_example/lib/metrics"
)

// TestProfileHandler tests the user profile handler
func TestProfileHandler(t *testing.T) {
	logger := logging.New("test-user-service")
	metricsCollector := metrics.New("test-user-service")

	handler := profileHandler(logger, metricsCollector)

	// Create a test request
	req := httptest.NewRequest("GET", "/users/profile?id=user-123", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify response contains expected fields
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if _, ok := response["user_id"]; !ok {
		t.Error("Expected response to include 'user_id' field")
	}

	if _, ok := response["name"]; !ok {
		t.Error("Expected response to include 'name' field")
	}

	if _, ok := response["email"]; !ok {
		t.Error("Expected response to include 'email' field")
	}
}

// TestUpdateUserHandler tests the user update handler
func TestUpdateUserHandler(t *testing.T) {
	logger := logging.New("test-user-service")
	metricsCollector := metrics.New("test-user-service")

	handler := updateUserHandler(logger, nil, metricsCollector)

	// Create a test request
	req := httptest.NewRequest("POST", "/users/update", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify response contains expected fields
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if _, ok := response["user_id"]; !ok {
		t.Error("Expected response to include 'user_id' field")
	}

	if _, ok := response["updated_at"]; !ok {
		t.Error("Expected response to include 'updated_at' field")
	}

	if success, ok := response["success"].(bool); !ok || !success {
		t.Error("Expected 'success' field to be true")
	}
}

// TestHealthHandler tests the health check endpoint
func TestHealthHandler(t *testing.T) {
	handler := healthHandler()

	// Create a test request
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify response contains expected fields
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	status, ok := response["status"]
	if !ok {
		t.Error("Expected response to include 'status' field")
	}

	if status != "ok" {
		t.Errorf("Expected status to be 'ok', got '%s'", status)
	}
}
