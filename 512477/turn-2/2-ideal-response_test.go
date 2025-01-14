// filename: main_test.go
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const (
	alphanum  = "abcdefghijklmnopqrstuvwxyz0123456789"
	sliceSize = 200
)

var inputStrings = generateStrings(sliceSize)

// Generate a slice of random strings
func generateStrings(n int) []string {
	result := make([]string, n)
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

func concatBySprintf(inpStr []string) string {
	result := ""
	for _, str := range inpStr {
		result = fmt.Sprintf("%s%s", result, str)
	}
	return result
}

func concatByPlusOperator(inpStr []string) string {
	result := ""
	for _, str := range inpStr {
		result += str
	}
	return result
}

func concatByBuffer(inpStr []string) string {
	var buf = &bytes.Buffer{}
	defer buf.Reset()
	for _, s := range inpStr {
		buf.WriteString(s)
	}
	return buf.String()
}

func concatByBuilder(inpStr []string) string {
	var sb strings.Builder
	for _, str := range inpStr {
		sb.WriteString(str)
	}
	return sb.String()
}

// Benchmark using fmt.Sprintf
func BenchmarkConcatBySprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatBySprintf(inputStrings)
	}
}

// Benchmark using traditional string concatenation
func BenchmarkConcatByPlusOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByPlusOperator(inputStrings)
	}
}

// Benchmark using bytes buffer for optimized concatenation
func BenchmarkConcatByBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByBuffer(inputStrings)
	}
}

// Benchmark using strings.Builder for optimized concatenation
func BenchmarkConcatByBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatByBuilder(inputStrings)
	}
}
