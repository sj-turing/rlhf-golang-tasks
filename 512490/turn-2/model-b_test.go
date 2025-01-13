package main

import (
	"testing"
)

func BenchmarkSequentialJoin(b *testing.B) {
	strList := generateStrings(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sequentialJoin(strList)
	}
}

func BenchmarkConcurrentJoin(b *testing.B) {
	strList := generateStrings(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrentJoin(strList, 100)
	}
}
