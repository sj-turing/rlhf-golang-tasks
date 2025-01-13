package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkJoinSequential(b *testing.B) {
	strings := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		strings[i] = fmt.Sprintf("string%d", i)
	}
	for n := 0; n < b.N; n++ {
		_ = strings.Join(strings, " ")
	}
}

func BenchmarkJoinConcurrent(b *testing.B) {
	numStrings := b.N
	parts := 10
	stringList := make([]string, numStrings)
	for i := 0; i < numStrings; i++ {
		stringList[i] = fmt.Sprintf("string%d", i)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ch := make(chan string, parts)
		var wg sync.WaitGroup

		for part := 0; part < parts; part++ {
			start := part * numStrings / parts
			end := (part + 1) * numStrings / parts
			wg.Add(1)

			go func() {
				defer wg.Done()
				var partResult []string
				for i := start; i < end; i++ {
					partResult = append(partResult, stringList[i])
				}
				ch <- strings.Join(partResult, " ")
			}()
		}

		var result string
		for part := 0; part < parts; part++ {
			result += <-ch + " "
		}
		result = strings.TrimRight(result, " ")
	}
}
