package main

import (
	"benchmark"
)

func BenchmarkAppendNoPrealloc(b *benchmark.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < 100000; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAppendWithPrealloc(b *benchmark.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < 100000; j++ {
			s = append(s, j)
		}
	}
}

func main() {
	benchmark.Main()
}
