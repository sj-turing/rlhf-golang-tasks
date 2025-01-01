package main

import (
	"fmt"
)

type FinancialDataSlice []float64

func (fs FinancialDataSlice) FindMaxProfit(start, end int) float64 {
	if start < 0 || end >= len(fs) || start > end {
		return 0.0
	}

	maxProfit := 0.0
	minPrice := fs[start]

	for i := start + 1; i <= end; i++ {
		currentProfit := fs[i] - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
		if fs[i] < minPrice {
			minPrice = fs[i]
		}
	}

	return maxProfit
}

func main() {
	data := FinancialDataSlice{100.0, 105.0, 102.0, 108.0, 110.0, 112.0, 109.0, 111.0}

	fmt.Println("Original Data:", data)
	maxProfit := data.FindMaxProfit(2, 6)
	fmt.Println("Maximum Profit between indices 2 and 6:", maxProfit) // Output: 13.0
}
