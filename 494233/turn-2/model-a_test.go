package main

import (
	"testing"
)

func BenchmarkRangeLoop(b *testing.B) {
	slice := make([]int, 100000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for _, value := range slice {
			sum += value
		}
	}
}

func BenchmarkIndexLoop(b *testing.B) {
	slice := make([]int, 100000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < len(slice); i++ {
			sum += slice[i]
		}
	}
}
