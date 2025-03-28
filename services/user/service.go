package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bazel_query_example/lib/config"
	"bazel_query_example/lib/logging"
	"bazel_query_example/lib/metrics"
)

const serviceName = "user-service"

func main() {
	// Initialize dependencies
	logger := logging.New(serviceName)
	logger.Info("Starting service", map[string]interface{}{
		"service": serviceName,
	})

	configReader := config.New("production")
	metricsCollector := metrics.New(serviceName)

	// Create and start the server
	mux := http.NewServeMux()
	setupRoutes(mux, logger, configReader, metricsCollector)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		logger.Info("Shutting down server", nil)
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Server shutdown failed", err, nil)
		}
	}()

	// Start the server
	logger.Info("Server listening", map[string]interface{}{
		"address": server.Addr,
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("Server error", err, nil)
		os.Exit(1)
	}

	logger.Info("Server stopped", nil)
}
