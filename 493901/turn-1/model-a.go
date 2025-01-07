package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// The -t or --test parameter in go test is used to run test.
func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}

	if result := bubbleSort(arr); !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
