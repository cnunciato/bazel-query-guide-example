package main

import (
	"encoding/json"
	"net/http"
	"time"

	"bazel_query_example/lib/config"
	"bazel_query_example/lib/logging"
	"bazel_query_example/lib/metrics"
)

// setupRoutes configures all HTTP routes for the service
func setupRoutes(mux *http.ServeMux, logger *logging.Logger, configReader *config.Reader, metrics *metrics.Collector) {
	mux.HandleFunc("/auth/login", loginHandler(logger, metrics))
	mux.HandleFunc("/auth/verify", verifyHandler(logger, configReader, metrics))
	mux.HandleFunc("/health", healthHandler())
}

// loginHandler handles user authentication requests
func loginHandler(logger *logging.Logger, metrics *metrics.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer metrics.RecordDuration("request_duration", time.Since(start), map[string]string{
			"path":   "/auth/login",
			"method": r.Method,
		})

		// Simple auth logic would go here
		response := map[string]string{
			"token":   "fake-jwt-token",
			"user_id": "user-123",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		logger.Info("Login successful", map[string]interface{}{
			"user_id": "user-123",
		})
	}
}

// verifyHandler verifies authentication tokens
func verifyHandler(logger *logging.Logger, configReader *config.Reader, metrics *metrics.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer metrics.RecordDuration("request_duration", time.Since(start), map[string]string{
			"path":   "/auth/verify",
			"method": r.Method,
		})

		// Token verification logic would go here
		response := map[string]bool{
			"valid": true,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		logger.Info("Token verified", map[string]interface{}{
			"token_id": "token-456",
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
