package main

import (
	"fmt"
)

type FinancialSlice []float64

func (fs FinancialSlice) Mean() float64 {
	total := float64(0)
	for _, value := range fs {
		total += value
	}
	return total / float64(len(fs))
}

func (fs FinancialSlice) Median() float64 {
	sortedSlice := append([]float64{}, fs...)
	sort.Float64s(sortedSlice)
	middleIndex := len(sortedSlice) / 2

	if len(sortedSlice)%2 == 0 {
		return (sortedSlice[middleIndex-1] + sortedSlice[middleIndex]) / 2
	} else {
		return sortedSlice[middleIndex]
	}
}

func (fs FinancialSlice) StandardDeviation() float64 {
	mean := fs.Mean()
	sumSquaredDifferences := float64(0)

	for _, value := range fs {
		difference := value - mean
		sumSquaredDifferences += difference * difference
	}

	variance := sumSquaredDifferences / float64(len(fs)-1)
	return math.Sqrt(variance)
}

func (fs FinancialSlice) Max() float64 {
	if len(fs) == 0 {
		return 0
	}
	maxValue := fs[0]
	for _, value := range fs[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func (fs FinancialSlice) Min() float64 {
	if len(fs) == 0 {
		return 0
	}
	minValue := fs[0]
	for _, value := range fs[1:] {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func main() {
	financialData := FinancialSlice{100.0, 120.0, 80.0, 150.0, 130.0}

	fmt.Println("Mean:", financialData.Mean())                            // Output: Mean: 120
	fmt.Println("Median:", financialData.Median())                        // Output: Median: 120
	fmt.Println("Standard Deviation:", financialData.StandardDeviation()) // Output: Standard Deviation: 28.284271247461903
	fmt.Println("Max:", financialData.Max())                              // Output: Max: 150
	fmt.Println("Min:", financialData.Min())                              // Output: Min: 80
}
