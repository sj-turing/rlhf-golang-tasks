package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	numStrings  = 100000
	concurrency = 100
)

func main() {
	// Sequential concatenation
	start := time.Now()
	sequentialConcatenation()
	elapsed := time.Since(start)
	fmt.Printf("Sequential concatenation took: %s\n", elapsed)

	// Concurrent concatenation
	start = time.Now()
	concurrentConcatenation()
	elapsed = time.Since(start)
	fmt.Printf("Concurrent concatenation took: %s\n", elapsed)

	// Concurrent concatenation with channels
	start = time.Now()
	concurrentConcatenationWithChannels()
	elapsed = time.Since(start)
	fmt.Printf("Concurrent concatenation with channels took: %s\n", elapsed)
}

func sequentialConcatenation() {
	var result string
	for i := 0; i < numStrings; i++ {
		result += fmt.Sprintf("%d", i)
	}
	fmt.Println(len(result))
}

func concurrentConcatenation() {
	var result string
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				result += fmt.Sprintf("%d", i)
			}
		}(i*numStrings/concurrency, (i+1)*numStrings/concurrency)
	}

	wg.Wait()
	fmt.Println(len(result))
}

func concurrentConcatenationWithChannels() {
	var result string
	var wg sync.WaitGroup
	wg.Add(concurrency)

	ch := make(chan string)

	for i := 0; i < concurrency; i++ {
		go func(start, end int) {
			defer wg.Done()
			var part string
			for i := start; i < end; i++ {
				part += fmt.Sprintf("%d", i)
			}
			ch <- part
		}(i*numStrings/concurrency, (i+1)*numStrings/concurrency)
	}

	go func() {
		defer close(ch)
		for part := range ch {
			result += part
		}
	}()

	wg.Wait()
	fmt.Println(len(result))
}
