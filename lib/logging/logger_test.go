package logging

import (
	"errors"
	"testing"
)

func TestLoggerInfo(t *testing.T) {
	logger := New("test-service")

	// This won't actually assert anything since our logger implementation
	// just prints to stdout, but at least it ensures the code runs without panicking
	logger.Info("test message", map[string]interface{}{
		"key": "value",
	})
}

func TestLoggerError(t *testing.T) {
	logger := New("test-service")

	err := errors.New("test error")
	logger.Error("error occurred", err, map[string]interface{}{
		"key": "value",
	})
}

func TestLoggerWithNilFields(t *testing.T) {
	logger := New("test-service")

	// Test with nil fields to ensure no panic
	logger.Info("test message", nil)
	logger.Error("error occurred", errors.New("test error"), nil)
}
