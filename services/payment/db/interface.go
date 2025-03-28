package db

import (
	"context"
)

// Payment represents payment information
type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

// Database defines the interface for payment storage
type Database interface {
	// SavePayment stores payment information
	SavePayment(ctx context.Context, payment *Payment) error

	// GetPayment retrieves payment information by ID
	GetPayment(ctx context.Context, paymentID string) (*Payment, error)

	// Close closes the database connection
	Close() error
}
