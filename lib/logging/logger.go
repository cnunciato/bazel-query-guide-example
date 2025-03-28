package logging

import (
	"fmt"
	"time"
)

// Logger provides structured logging functionality
type Logger struct {
	serviceName string
	level       string
}

// New creates a new logger
func New(serviceName string) *Logger {
	return &Logger{
		serviceName: serviceName,
		level:       "info",
	}
}

// Info logs an informational message
func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.log("INFO", msg, fields)
}

// Error logs an error message
func (l *Logger) Error(msg string, err error, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["error"] = err.Error()
	l.log("ERROR", msg, fields)
}

// log implements the actual logging logic
func (l *Logger) log(level, msg string, fields map[string]interface{}) {
	// In a real implementation, this would format and write logs
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] %s [%s] %s %v\n", timestamp, level, l.serviceName, msg, fields)
}
