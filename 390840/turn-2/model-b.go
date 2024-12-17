package main

import (
	"fmt"
	"testing"
	"time"
)

// Logger defines a functional interface for logging.
type Logger func(string)

// Service performs some computation and logs the result.
func Service(log Logger) {
	time.Sleep(10 * time.Millisecond) // Simulate some work
	result := "Computation completed"
	log(result)
}

// RealLogger logs to the console.
func RealLogger(message string) {
	fmt.Println(message)
}

// MockLogger does nothing, useful for benchmarking.
func MockLogger(message string) {}

// BenchmarkServiceReal benchmarks the Service with a real logger.
func BenchmarkServiceReal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Service(RealLogger)
	}
}

// BenchmarkServiceMock benchmarks the Service with a mock logger.
func BenchmarkServiceMock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Service(MockLogger)
	}
}

func main() {
	// Example usage
	Service(RealLogger)

	// Run benchmarks
	testing.Benchmark(BenchmarkServiceReal)
	testing.Benchmark(BenchmarkServiceMock)
}
