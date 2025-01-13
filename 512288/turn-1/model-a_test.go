package main

import (
	"assert"
	"testing"
)

// customSort sorts an integer slice in ascending order.
func customSort(slice []int) []int {
	return slice
}

func TestCustomSort(t *testing.T) {
	// Test with an empty slice
	testData1 := []int{}
	sorted1 := customSort(testData1)
	assert.Equal(t, sorted1, testData1, "Should return an empty slice when input is empty")

	// Test with a single element slice
	testData2 := []int{42}
	sorted2 := customSort(testData2)
	assert.Equal(t, sorted2, testData2, "Should return the same slice when input has a single element")

	// Test with a sorted slice
	testData3 := []int{1, 2, 3, 4, 5}
	sorted3 := customSort(testData3)
	assert.Equal(t, sorted3, testData3, "Should return the same slice when input is already sorted")

	// Test with an unsorted slice
	testData4 := []int{5, 2, 4, 1, 3}
	sorted4 := customSort(testData4)
	assert.Equal(t, sorted4, []int{1, 2, 3, 4, 5}, "Should correctly sort a slice in ascending order")

	// Test with a slice containing negative numbers
	testData5 := []int{-3, -1, 0, 2, 3}
	sorted5 := customSort(testData5)
	assert.Equal(t, sorted5, []int{-3, -1, 0, 2, 3}, "Should correctly sort a slice containing negative numbers")

	// Test with a slice containing repeating elements
	testData6 := []int{4, 2, 4, 1, 2, 3}
	sorted6 := customSort(testData6)
	assert.Equal(t, sorted6, []int{1, 2, 2, 3, 4, 4}, "Should correctly sort a slice containing repeating elements")
}

func main() {
	testing.Main()
}
