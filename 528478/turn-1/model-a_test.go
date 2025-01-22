package benchmarkfmt

import (
	"fmt"
	"testing"
)

// Benchmark for fmt.Sprintf with different data types
func BenchmarkSprintf(b *testing.B) {
	var (
		intVal    = 42
		floatVal  = 3.14159
		stringVal = "Hello, World!"
	)

	// Benchmark with int
	b.Run("Int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%d", intVal)
		}
	})

	// Benchmark with float
	b.Run("Float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%.2f", floatVal)
		}
	})

	// Benchmark with string
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", stringVal)
		}
	})
}

// Benchmark for fmt.Printf with different data types
func BenchmarkPrintf(b *testing.B) {
	var (
		intVal    = 42
		floatVal  = 3.14159
		stringVal = "Hello, World!"
	)

	// Benchmark with int
	b.Run("Int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Printf("%d", intVal)
		}
	})

	// Benchmark with float
	b.Run("Float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Printf("%.2f", floatVal)
		}
	})

	// Benchmark with string
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Printf("%s", stringVal)
		}
	})
}
