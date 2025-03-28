package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bazel_query_example/lib/logging"
	"bazel_query_example/lib/metrics"
	db "bazel_query_example/services/payment/db"
)

// MockDB implements the db.Database interface for testing
type MockDB struct{}

// SavePayment mocks saving a payment
func (m *MockDB) SavePayment(ctx context.Context, payment *db.Payment) error {
	return nil
}

// GetPayment mocks retrieving a payment
func (m *MockDB) GetPayment(ctx context.Context, paymentID string) (*db.Payment, error) {
	return &db.Payment{
		ID:     paymentID,
		Amount: 100.0,
		Status: "completed",
	}, nil
}

// Close mocks closing the database connection
func (m *MockDB) Close() error {
	return nil
}

// TestProcessPaymentHandler tests the payment processing handler
func TestProcessPaymentHandler(t *testing.T) {
	logger := logging.New("test-payment-service")
	metricsCollector := metrics.New("test-payment-service")
	mockDB := &MockDB{}

	handler := processPaymentHandler(logger, mockDB, metricsCollector)

	// Create a test request
	req := httptest.NewRequest("POST", "/payments/process", nil)
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

	if _, ok := response["payment_id"]; !ok {
		t.Error("Expected response to include 'payment_id' field")
	}

	if status, ok := response["status"]; !ok || status != "processed" {
		t.Errorf("Expected status 'processed', got '%s'", status)
	}
}

// [rest of the file remains unchanged]
