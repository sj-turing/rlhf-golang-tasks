package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntegerSlice(t *testing.T) {
	t.Run("Test sort empty slice", func(t *testing.T) {
		input := []int{}
		sort.Ints(input)
		assert.Equal(t, []int{}, input, "Empty slice should remain empty after sorting")
	})

	t.Run("Test sort single element slice", func(t *testing.T) {
		input := []int{42}
		sort.Ints(input)
		assert.Equal(t, []int{42}, input, "Single element slice should remain unchanged after sorting")
	})

	t.Run("Test sort sorted slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		sort.Ints(input)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, input, "Sorted slice should remain unchanged after sorting")
	})

	t.Run("Test sort unsorted slice", func(t *testing.T) {
		input := []int{5, 2, 3, 4, 1}
		sort.Ints(input)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, input, "Unsorted slice should be sorted correctly")
	})

	t.Run("Test sort slice with negative numbers", func(t *testing.T) {
		input := []int{-5, -2, -3, -4, -1}
		sort.Ints(input)
		assert.Equal(t, []int{-5, -4, -3, -2, -1}, input, "Slice with negative numbers should be sorted correctly")
	})

	t.Run("Test sort slice with duplicate numbers", func(t *testing.T) {
		input := []int{2, 5, 3, 2, 4, 5}
		sort.Ints(input)
		assert.Equal(t, []int{2, 2, 3, 4, 5, 5}, input, "Slice with duplicate numbers should be sorted correctly")
	})
}
