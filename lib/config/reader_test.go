package config

import (
	"context"
	"testing"
)

func TestGetValue(t *testing.T) {
	reader := New("test")
	value, err := reader.GetValue(context.Background(), "test-key")
	if err != nil {
		t.Errorf("GetValue returned an error: %v", err)
	}

	if value == "" {
		t.Error("Expected non-empty value")
	}
}

func TestGetIntValue(t *testing.T) {
	reader := New("test")
	value, err := reader.GetIntValue(context.Background(), "test-key")
	if err != nil {
		t.Errorf("GetIntValue returned an error: %v", err)
	}

	if value != 42 {
		t.Errorf("Expected value 42, got %d", value)
	}
}

func TestNewReader(t *testing.T) {
	reader := New("production")
	if reader == nil {
		t.Fatal("Expected non-nil reader")
	}

	if reader.environment != "production" {
		t.Errorf("Expected environment 'production', got '%s'", reader.environment)
	}

	if reader.timeout.Seconds() != 5 {
		t.Errorf("Expected timeout 5 seconds, got %v", reader.timeout)
	}
}
