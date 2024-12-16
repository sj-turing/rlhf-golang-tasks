package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Metric struct represents a metric
type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// MetricsCollection holds a collection of metrics
type MetricsCollection []Metric

// CollectMetrics collects metrics and sends them to the metrics channel
func CollectMetrics(metrics chan<- Metric, interval time.Duration) {
	defer close(metrics)
	for {
		time.Sleep(interval)
		metrics <- Metric{Name: "CPU Usage", Value: rand.Float64() * 100, Timestamp: time.Now()}
		metrics <- Metric{Name: "Memory Usage", Value: rand.Float64() * 1024, Timestamp: time.Now()}
	}
}

// ServeMetricsHandler handles HTTP requests to the /metrics endpoint
func ServeMetricsHandler(metrics <-chan Metric) http.HandlerFunc {
	var collectedMetrics MetricsCollection

	go func() {
		for metric := range metrics {
			collectedMetrics = append(collectedMetrics, metric)
			// Optionally, trim the list to avoid growing indefinitely
			if len(collectedMetrics) > 100 {
				collectedMetrics = collectedMetrics[1:]
			}
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(collectedMetrics)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create metrics channel
	metrics := make(chan Metric)

	// Start goroutines to collect metrics
	go CollectMetrics(metrics, 1*time.Second)

	// Start HTTP server
	http.HandleFunc("/metrics", ServeMetricsHandler(metrics))
	fmt.Println("Metrics server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
