package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntegerSlice(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{}, expected: []int{}},
		{input: []int{42}, expected: []int{42}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 2, 3, 4, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{-5, -2, -3, -4, -1}, expected: []int{-5, -4, -3, -2, -1}},
		{input: []int{2, 5, 3, 2, 4, 5}, expected: []int{2, 2, 3, 4, 5, 5}},
	}

	t.Run("Test sort various slices", func(t *testing.T) {
		for _, tc := range testCases {
			sort.Ints(tc.input)
			got := tc.input
			assert.Equal(t, tc.expected, got, "Failed for input: %v", tc.input)
		}
	})
}
