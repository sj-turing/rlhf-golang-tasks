package main

import (
	"fmt"
	"testing"
)

func BenchmarkSprintfConcatenation(b *testing.B) {
	strs := make([]string, 100)
	for i := 0; i < 100; i++ {
		strs[i] = fmt.Sprintf("string%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result string
		for _, str := range strs {
			result = fmt.Sprintf("%s%s", result, str)
		}
	}
}

func BenchmarkStringConcatenation(b *testing.B) {
	strs := make([]string, 100)
	for i := 0; i < 100; i++ {
		strs[i] = fmt.Sprintf("string%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result string
		for _, str := range strs {
			result += str
		}
	}
}

func BenchmarkBufferConcatenation(b *testing.B) {
	strs := make([]string, 100)
	for i := 0; i < 100; i++ {
		strs[i] = fmt.Sprintf("string%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer strings.Builder
		for _, str := range strs {
			buffer.WriteString(str)
		}
		result := buffer.String()
	}
}
