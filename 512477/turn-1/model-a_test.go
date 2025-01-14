package main

import (
	"fmt"
	"testing"
)

// Generate a slice of 100 strings for benchmarking
func generateStrings() []string {
	return make([]string, 100, 100)
}

// Benchmark using fmt.Sprintf
func BenchmarkSprintf(b *testing.B) {
	strings := generateStrings()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", strings...)
	}
}

// Benchmark using traditional string concatenation
func BenchmarkConcatenation(b *testing.B) {
	strings := generateStrings()
	for i := 0; i < b.N; i++ {
		var result string
		for _, s := range strings {
			result += s
		}
		_ = result
	}
}

// Benchmark using strings.Builder for optimized concatenation
func BenchmarkBuilder(b *testing.B) {
	strings := generateStrings()
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		for _, s := range strings {
			buf.WriteString(s)
		}
		_ = buf.String()
	}
}

func main() {
	testing.Run(nil, func(t *testing.T) {
		testing.Benchmark(BenchmarkSprintf)
		testing.Benchmark(BenchmarkConcatenation)
		testing.Benchmark(BenchmarkBuilder)
	})
}
