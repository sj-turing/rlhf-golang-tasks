// filename: main_test.go
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

const (
	alphanum = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var inputString = generateStrings()

// Generate a slice of 100 strings for benchmarking
func generateStrings() []string {
	result := make([]string, 100)
	for index := range result {
		result[index] = getString(rand.Intn(50))
	}
	return result
}

func getString(length int) string {
	str := ""

	for length > 0 {
		str += string(alphanum[rand.Intn(len(alphanum))])
		length--
	}
	return str
}

func concatBySprintf(strings []string) string {
	result := ""
	for _, str := range strings {
		result = fmt.Sprintf("%s%s", result, str)
	}
	return result
}

func concatByBuffer(strings []string) string {
	var buf = &bytes.Buffer{}
	defer buf.Reset()
	for _, s := range strings {
		buf.WriteString(s)
	}
	return buf.String()
}

func concatByPlusOperator(strings []string) string {
	result := ""
	for _, str := range strings {
		result += str
	}
	return result
}

// Benchmark using fmt.Sprintf
func BenchmarkConcatBySprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatBySprintf(inputString)
	}
}

// Benchmark using traditional string concatenation
func BenchmarkConcatByPlusOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByPlusOperator(inputString)
	}
}

// Benchmark using bytes buffer for optimized concatenation
func BenchmarkConcatByBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByBuffer(inputString)
	}
}
