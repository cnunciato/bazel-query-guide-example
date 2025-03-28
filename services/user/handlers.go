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
	mux.HandleFunc("/users/profile", profileHandler(logger, metrics))
	mux.HandleFunc("/users/update", updateUserHandler(logger, configReader, metrics))
	mux.HandleFunc("/health", healthHandler())
}

// profileHandler handles user profile requests
func profileHandler(logger *logging.Logger, metrics *metrics.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer metrics.RecordDuration("request_duration", time.Since(start), map[string]string{
			"path":   "/users/profile",
			"method": r.Method,
		})

		// User profile logic would go here
		response := map[string]interface{}{
			"user_id":    "user-123",
			"name":       "John Doe",
			"email":      "john.doe@example.com",
			"created_at": "2023-01-01T00:00:00Z",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		logger.Info("Profile retrieved", map[string]interface{}{
			"user_id": "user-123",
		})
	}
}

// updateUserHandler handles user update requests
func updateUserHandler(logger *logging.Logger, configReader *config.Reader, metrics *metrics.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer metrics.RecordDuration("request_duration", time.Since(start), map[string]string{
			"path":   "/users/update",
			"method": r.Method,
		})

		// User update logic would go here
		response := map[string]interface{}{
			"user_id":    "user-123",
			"updated_at": time.Now().Format(time.RFC3339),
			"success":    true,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		logger.Info("User updated", map[string]interface{}{
			"user_id": "user-123",
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
