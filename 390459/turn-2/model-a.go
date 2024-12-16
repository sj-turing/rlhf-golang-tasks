package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Metric struct represents a metric
type Metric struct {
	Name  string
	Value float64
}

// Global variables to store metrics
var (
	metrics     = make(chan Metric)
	metricData  = make(map[string]float64)
	metricsLock = sync.RWMutex{}
)

// CollectMetrics collects metrics and sends them to the metrics channel
func CollectMetrics(interval time.Duration) {
	for {
		time.Sleep(interval)
		// Simulating metric collection
		metrics <- Metric{Name: "CPU_Usage", Value: rand.Float64() * 100}
		metrics <- Metric{Name: "Memory_Usage", Value: rand.Float64() * 1024}
	}
}

// Metric Aggregator
func ProcessMetrics() {
	for metric := range metrics {
		metricsLock.Lock()
		metricData[metric.Name] = metric.Value
		metricsLock.Unlock()
	}
}

// Metrics HTTP Handler
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metricsLock.RLock()
	defer metricsLock.RUnlock()

	for name, value := range metricData {
		fmt.Fprintf(w, "%s %.2f\n", name, value)
	}
}

// Main function for starting the application
func main() {
	rand.Seed(time.Now().UnixNano())

	// Start Goroutines
	go CollectMetrics(1 * time.Second)
	go ProcessMetrics()

	// Set up HTTP server
	http.HandleFunc("/metrics", MetricsHandler)

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
