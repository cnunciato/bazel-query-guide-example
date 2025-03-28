package main

import (
	"encoding/json"
	"net/http"
	"time"

	"bazel_query_example/lib/config"
	"bazel_query_example/lib/logging"
	"bazel_query_example/lib/metrics"
	db "bazel_query_example/services/payment/db"
)

// setupRoutes configures all HTTP routes for the service
func setupRoutes(mux *http.ServeMux, logger *logging.Logger, configReader *config.Reader, metrics *metrics.Collector, database db.Database) {
	mux.HandleFunc("/payments/process", processPaymentHandler(logger, database, metrics))
	mux.HandleFunc("/payments/status", paymentStatusHandler(logger, configReader, database, metrics))
	mux.HandleFunc("/health", healthHandler())
}

// processPaymentHandler handles payment processing requests
func processPaymentHandler(logger *logging.Logger, database db.Database, metrics *metrics.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer metrics.RecordDuration("request_duration", time.Since(start), map[string]string{
			"path":   "/payments/process",
			"method": r.Method,
		})

		// Payment processing logic would go here
		response := map[string]string{
			"payment_id": "payment-789",
			"status":     "processed",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		logger.Info("Payment processed", map[string]interface{}{
			"payment_id": "payment-789",
		})
	}
}

// paymentStatusHandler handles payment status check requests
func paymentStatusHandler(logger *logging.Logger, configReader *config.Reader, database db.Database, metrics *metrics.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer metrics.RecordDuration("request_duration", time.Since(start), map[string]string{
			"path":   "/payments/status",
			"method": r.Method,
		})

		// Payment status check logic would go here
		response := map[string]string{
			"payment_id": "payment-789",
			"status":     "completed",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		logger.Info("Payment status checked", map[string]interface{}{
			"payment_id": "payment-789",
		})
	}
}

// healthHandler provides a basic health check endpoint
func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}
}
