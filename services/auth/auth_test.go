package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bazel_query_example/lib/logging"
	"bazel_query_example/lib/metrics"
)

// TestLoginHandler tests the login handler functionality
func TestLoginHandler(t *testing.T) {
	logger := logging.New("test-auth-service")
	metricsCollector := metrics.New("test-auth-service")

	handler := loginHandler(logger, metricsCollector)

	// Create a test request
	req := httptest.NewRequest("POST", "/auth/login", nil)
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

	if _, ok := response["token"]; !ok {
		t.Error("Expected response to include 'token' field")
	}

	if _, ok := response["user_id"]; !ok {
		t.Error("Expected response to include 'user_id' field")
	}
}

// TestVerifyHandler tests the token verification handler
func TestVerifyHandler(t *testing.T) {
	logger := logging.New("test-auth-service")
	metricsCollector := metrics.New("test-auth-service")
	handler := verifyHandler(logger, nil, metricsCollector)

	// Create a test request
	req := httptest.NewRequest("POST", "/auth/verify", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify response contains expected fields
	var response map[string]bool
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if _, ok := response["valid"]; !ok {
		t.Error("Expected response to include 'valid' field")
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
