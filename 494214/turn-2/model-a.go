package main

import (
	"fmt"
	"math"
)

// FinancialSlice represents a slice of financial data (e.g., stock prices).
type FinancialSlice []float64

// NewFinancialSlice creates a new FinancialSlice from a slice of floats.
func NewFinancialSlice(data []float64) FinancialSlice {
	return FinancialSlice(data)
}

// CalculateMaxProfit finds the maximum profit possible in the given interval.
func (fs FinancialSlice) CalculateMaxProfit(start, end int) float64 {
	if start < 0 || end >= len(fs) || start > end {
		return 0.0
	}

	minPrice := fs[start]
	maxProfit := 0.0

	for _, price := range fs[start+1 : end+1] {
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	}

	return maxProfit
}

func main() {
	data := []float64{7.0, 1.0, 5.0, 3.0, 6.0, 4.0}
	fs := NewFinancialSlice(data)

	fmt.Println("Original Data:", fs)

	// Calculate max profit in the interval from index 1 to 4
	maxProfit := fs.CalculateMaxProfit(1, 4)
	fmt.Println("Max Profit in Interval (1 to 4):", maxProfit)
}
