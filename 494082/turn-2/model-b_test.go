// Filename: main_test.go

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const (
	// Number of strings to concatenate
	numStrings = 100000
	// Length of each string
	stringLength = 20
)

// Generate random strings
func generateStrings() []string {
	strs := make([]string, numStrings)
	for i := range strs {
		b := make([]byte, stringLength)
		for j := 0; j < stringLength; j++ {
			b[j] = byte('a' + rand.Intn(26))
		}
		strs[i] = string(b)
	}
	return strs
}

// Benchmark using fmt.Sprintf
func BenchmarkSprintf(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var result string
		for _, str := range strs {
			result = fmt.Sprintf("%s%s", result, str)
		}
	}
}

// Benchmark using basic string concatenation
func BenchmarkConcatenation(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var result string
		for _, str := range strs {
			result += str
		}
	}
}

// Benchmark using bytes.Buffer
func BenchmarkBuffer(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		for _, str := range strs {
			buf.WriteString(str)
		}
	}
}

// Benchmark using strings.Builder
func BenchmarkBuilder(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var b strings.Builder
		for _, str := range strs {
			b.WriteString(str)
		}
	}
}

// Benchmark using by writing directly into a byte slice
func BenchmarkBufferPooling(b *testing.B) {
	strs := generateStrings()
	bytePool := make([][]byte, numStrings)
	for i := 0; i < numStrings; i++ {
		bytePool[i] = make([]byte, stringLength)
	}

	for n := 0; n < b.N; n++ {
		// Create a buffer of size equal to the combined length of all strings
		buf := make([]byte, numStrings*stringLength)
		var offset int
		for _, str := range strs {
			copy(buf[offset:], str)
			offset += len(str)
		}
	}
}
