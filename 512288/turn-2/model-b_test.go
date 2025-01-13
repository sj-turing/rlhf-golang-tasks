package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntegerSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element slice",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Sorted slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted slice",
			input:    []int{5, 2, 3, 4, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Slice with negative numbers",
			input:    []int{-5, -2, -3, -4, -1},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			name:     "Slice with duplicate numbers",
			input:    []int{2, 5, 3, 2, 4, 5},
			expected: []int{2, 2, 3, 4, 5, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := make([]int, len(tc.input))
			copy(got, tc.input)
			sort.Ints(got)
			assert.Equal(t, tc.expected, got, "Input slice should be sorted correctly")
		})
	}
}
