package main

import (
	"fmt"
	"testing"
)

// DataSource defines a functional interface for data sources.
type DataSource func() []int

// ProcessData processes the data from the provided data source.
func ProcessData(source DataSource) int {
	data := source()
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

// RealDataSource simulates a real data source.
func RealDataSource() []int {
	return []int{1, 2, 3, 4, 5}
}

// MockDataSource simulates a mock data source for testing.
func MockDataSource() []int {
	return []int{10, 20, 30}
}

// Benchmarking function for a real data source.
func BenchmarkProcessDataReal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessData(RealDataSource)
	}
}

// Benchmarking function for a mock data source.
func BenchmarkProcessDataMock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessData(MockDataSource)
	}
}

func main() {
	// Example usage
	total := ProcessData(RealDataSource)
	fmt.Printf("Total from RealDataSource: %d\n", total)

	totalMock := ProcessData(MockDataSource)
	fmt.Printf("Total from MockDataSource: %d\n", totalMock)
}
