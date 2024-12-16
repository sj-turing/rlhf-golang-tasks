package main

/*

import (
	"fmt"
	"math/rand"
	"time"
)

// Metric represents a simple metric data structure
type Metric struct {
	Name  string
	Value float64
}

// Metric source that generates random metrics
func metricSource(metrics chan<- Metric, stop chan<- struct{}) {
	for {
		select {
		case <-stop:
			return
		default:
			metric := Metric{
				Name:  "request_count",
				Value: rand.Float64() * 100, // Random value to simulate metric
			}
			metrics <- metric
			time.Sleep(1 * time.Second) // Simulate time taken to collect metrics
		}
	}
}

// Metric processor that processes incoming metrics
func metricProcessor(metrics <-chan Metric, stop <-chan struct{}) {
	for {
		select {
		case metric := <-metrics:
			// Process the metric (for demonstration, we just print it)
			fmt.Printf("Processing metric: %s = %f\n", metric.Name, metric.Value)
		case <-stop:
			return
		}
	}
}

func main() {
	metrics := make(chan Metric)
	stop := make(chan struct{})

	// Start the metric source goroutine
	go metricSource(metrics, stop)

	// Start the metric processor goroutine
	go metricProcessor(metrics, stop)

	// Let the system run for 10 seconds
	time.Sleep(10 * time.Second)

	// Stop the goroutines
	close(stop)

	fmt.Println("Metric collection stopped.")
}
*/
