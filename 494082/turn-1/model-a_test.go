package main

import (
	"fmt"
	"testing"
)

var shortStrings = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func TestBenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result string
		for _, s := range shortStrings {
			result += s
		}
	}
}

func TestBenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := string.Concat(shortStrings...)
	}
}

func TestBenchmarkSprintfSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := fmt.Sprintf("%s", "Hello, World!")
	}
}

func TestBenchmarkSprintfComplex(b *testing.B) {
	const num = 1234567890
	for i := 0; i < b.N; i++ {
		result := fmt.Sprintf("The number is %d", num)
	}
}

func BenchmarkMemoryAllocation(b *testing.B) {
	b.ResetTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result string
		for _, s := range shortStrings {
			result += s
		}
	}
}
