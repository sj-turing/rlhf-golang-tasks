package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func BenchmarkStringFormatting(t *testing.B) {
	const count = 1000000
	var s string

	t.Run("fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = fmt.Sprintf("String with %d arguments", i)
		}
	})
}

func main() {
	// Simulate running a lot of benchmarks and saving the results to a file.
	result, err := os.Create("benchmark_result.txt")
	if err != nil {
		log.Fatalf("Failed to create result file: %v", err)
	}
	defer result.Close()

	testStart := time.Now()

	// Run the benchmarks
	BenchmarkStringFormatting()

	testEnd := time.Now()
	duration := testEnd.Sub(testStart)
	fmt.Fprintf(result, "Test execution time: %v\n", duration)
}
