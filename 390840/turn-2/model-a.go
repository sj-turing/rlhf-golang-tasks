package main

import (
	"fmt"
	"testing"
)

// DataSource is a functional interface for fetching data.
type DataSource func() []int

// ProcessData processes integers from the data source and returns their sum.
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

// Benchmark for processing data from the real data source.
func BenchmarkProcessDataReal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessData(RealDataSource)
	}
}

// Benchmark for processing data from the mock data source.
func BenchmarkProcessDataMock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessData(MockDataSource)
	}
}

func main() {
	// Example usage of the ProcessData function
	total := ProcessData(RealDataSource)
	fmt.Printf("Total from RealDataSource: %d\n", total)

	totalMock := ProcessData(MockDataSource)
	fmt.Printf("Total from MockDataSource: %d\n", totalMock)
}
