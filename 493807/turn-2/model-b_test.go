package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello, %s! Your score is %d.", "Alice", 95)
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("Hello, %s! Your score is %d.\n", "Alice", 95)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join([]string{"Hello, ", "Alice", "! Your score is ", "95", "."}, "")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString("Hello, ")
		buf.WriteString("Alice")
		buf.WriteString("! Your score is ")
		buf.WriteString("95")
		buf.WriteString(".")
		_ = buf.String()
	}
}

func BenchmarkFprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "Hello, %s! Your score is %d.", "Alice", 95)
		_ = buf.String()
	}
}
