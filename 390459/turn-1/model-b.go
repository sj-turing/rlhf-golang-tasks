package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Metric struct represents a metric
type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// CollectMetrics collects metrics and sends them to the metrics channel
func CollectMetrics(metrics chan<- Metric, interval time.Duration) {
	defer close(metrics)
	for {
		time.Sleep(interval)
		// Example metrics
		metrics <- Metric{Name: "CPU Usage", Value: rand.Float64() * 100, Timestamp: time.Now()}
		metrics <- Metric{Name: "Memory Usage", Value: rand.Float64() * 1024, Timestamp: time.Now()}
	}
}

// ProcessMetrics processes the metrics
func ProcessMetrics(metrics <-chan Metric, processed chan<- Metric) {
	for metric := range metrics {
		// Simple processing: adding a random offset
		metric.Value += rand.Float64() * 5
		processed <- metric
	}
	close(processed)
}

// StoreMetrics stores the processed metrics
func StoreMetrics(processed <-chan Metric) {
	for metric := range processed {
		fmt.Printf("Storing Metric: %+v\n", metric)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create channels
	metrics := make(chan Metric)
	processed := make(chan Metric)

	// Start goroutines
	go CollectMetrics(metrics, 1*time.Second)
	go ProcessMetrics(metrics, processed)
	go StoreMetrics(processed)

	// Wait for everything to finish (this will never happen in a real application)
	<-processed
}
