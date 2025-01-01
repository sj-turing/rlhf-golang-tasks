package main

import (
	"fmt"
	"math"
)

type FinancialSlice []float64

// NewFinancialSlice creates a new FinancialSlice from a slice of floats.
func NewFinancialSlice(data []float64) FinancialSlice {
	return FinancialSlice(data)
}

// CalculateMovingAverage calculates the n-day moving average.
func (fs FinancialSlice) CalculateMovingAverage(n int) FinancialSlice {
	if n <= 0 || n > len(fs) {
		return fs
	}

	result := make(FinancialSlice, len(fs)-n+1)
	for i := 0; i < len(result); i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += fs[i+j]
		}
		result[i] = sum / float64(n)
	}
	return result
}

// SumSubset returns the sum of elements from index start to end.
func (fs FinancialSlice) SumSubset(start, end int) float64 {
	if start < 0 || end >= len(fs) || start > end {
		return 0.0
	}

	sum := 0.0
	for _, value := range fs[start : end+1] {
		sum += value
	}
	return sum
}

// FindMaxInRange returns the maximum value in the slice from index start to end.
func (fs FinancialSlice) FindMaxInRange(start, end int) float64 {
	if start < 0 || end >= len(fs) || start > end {
		return math.SmallestNonzeroFloat64
	}

	maxValue := fs[start]
	for _, value := range fs[start+1 : end+1] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// FindMinInRange returns the minimum value in the slice from index start to end.
func (fs FinancialSlice) FindMinInRange(start, end int) float64 {
	if start < 0 || end >= len(fs) || start > end {
		return math.MaxFloat64
	}

	minValue := fs[start]
	for _, value := range fs[start+1 : end+1] {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func main() {
	data := []float64{100.0, 105.0, 102.0, 108.0, 110.0, 112.0, 109.0, 111.0}
	fs := NewFinancialSlice(data)

	fmt.Println("Original Data:", fs)

	// Calculate 3-day moving average
	movingAverage := fs.CalculateMovingAverage(3)
	fmt.Println("3-Day Moving Average:", movingAverage)

	// Sum of elements from index 1 to 4
	subsetSum := fs.SumSubset(1, 4)
	fmt.Println("Sum of Subset (1 to 4):", subsetSum)

	// Maximum value in the range from index 2 to 5
	maxInRange := fs.FindMaxInRange(2, 5)
	fmt.Println("Max in Range (2 to 5):", maxInRange)

	// Minimum value in the range from index 3 to 6
	minInRange := fs.FindMinInRange(3, 6)
	fmt.Println("Min in Range (3 to 6):", minInRange)
}
