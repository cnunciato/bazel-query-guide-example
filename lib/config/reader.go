package config

import (
	"context"
	"time"
)

// Reader provides configuration loading functionality
type Reader struct {
	environment string
	timeout     time.Duration
}

// New creates a new config reader
func New(environment string) *Reader {
	return &Reader{
		environment: environment,
		timeout:     5 * time.Second,
	}
}

// GetValue retrieves a configuration value
func (r *Reader) GetValue(ctx context.Context, key string) (string, error) {
	// In a real implementation, this would load from a config store
	return "default-value", nil
}

// GetIntValue retrieves a configuration value as an integer
func (r *Reader) GetIntValue(ctx context.Context, key string) (int, error) {
	// In a real implementation, this would load from a config store
	return 42, nil
}
