package metrics

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func TestNewCollector(t *testing.T) {
	collector := New("test-service")
	if collector == nil {
		t.Fatal("Expected non-nil collector")
	}

	if collector.serviceName != "test-service" {
		t.Errorf("Expected service name 'test-service', got '%s'", collector.serviceName)
	}
}

func TestIncrementCounter(t *testing.T) {
	collector := New("test-service")

	// Set up a test counter
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_counter",
		Help: "Test counter for unit testing",
	})
	collector.counters["test_counter"] = counter

	// This should not panic
	collector.IncrementCounter("test_counter", map[string]string{
		"label": "value",
	})

	// Incrementing a non-existent counter should not panic
	collector.IncrementCounter("non_existent", nil)
}

func TestRecordDuration(t *testing.T) {
	collector := New("test-service")

	// Set up a test histogram
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "test_histogram",
		Help:    "Test histogram for unit testing",
		Buckets: prometheus.DefBuckets,
	})
	collector.histograms["test_histogram"] = histogram

	// This should not panic
	collector.RecordDuration("test_histogram", 100*time.Millisecond, map[string]string{
		"label": "value",
	})

	// Recording to a non-existent histogram should not panic
	collector.RecordDuration("non_existent", 100*time.Millisecond, nil)
}
