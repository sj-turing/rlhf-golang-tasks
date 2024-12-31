// filename: string_benchmarks_test.go
package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
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

// Benchmark String Concatenation with +
func BenchmarkConcatenatePlus(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var result string
		for _, str := range strs {
			result += str
		}
	}
}

// Benchmark String Concatenation with bytes.Buffer
func BenchmarkConcatenateBuffer(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		for _, str := range strs {
			buffer.WriteString(str)
		}
	}
}

// Benchmark String Formatting with fmt.Sprintf
func BenchmarkFormatSprintf(b *testing.B) {
	strs := generateStrings()
	for n := 0; n < b.N; n++ {
		var result string
		for _, str := range strs {
			result = fmt.Sprintf("%s %s", result, str)
		}
	}
}

// Benchmark String Replacement with strings.Replace
func BenchmarkReplace(b *testing.B) {
	str := generateStrings()[0]
	for n := 0; n < b.N; n++ {
		result := strings.Replace(str, "a", "b", -1)
	}
}

// Benchmark String Slicing
func BenchmarkSlicing(b *testing.B) {
	str := generateStrings()[0]
	for n := 0; n < b.N; n++ {
		result := str[:5]
	}
}
