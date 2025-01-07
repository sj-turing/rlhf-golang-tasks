package main

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = rand.Intn(100000)
	}
	for n := 0; n < b.N; n++ {
		bubbleSort(arr)
	}
}
