package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const (
	numStrings   = 100000
	stringLength = 10
)

// Generate random strings for benchmarking
func generateStrings() []string {
	strs := make([]string, numStrings)
	for i := range strs {
		bytes := make([]byte, stringLength)
		for j := 0; j < stringLength; j++ {
			bytes[j] = byte('a' + rand.Intn(26))
		}
		strs[i] = string(bytes)
	}
	return strs
}

// Benchmark fmt.Sprintf
func BenchmarkSprintf(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var result string
		for _, str := range strs {
			result = fmt.Sprintf("%s %s", result, str)
		}
	}
}

// Benchmark string.Concat
func BenchmarkConcat(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var result string
		for _, str := range strs {
			result = strings.Concat(result, str)
		}
	}
}

func main() {
	// Run the benchmarks
	testing.Benchmark(BenchmarkSprintf)
	testing.Benchmark(BenchmarkConcat)
}
