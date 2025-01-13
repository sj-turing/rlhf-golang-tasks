package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	numStrings  = 100000
	concurrency = 100
)

func main() {
	strList := generateStrings(numStrings)

	start := time.Now()
	sequentialJoin(strList)
	elapsed := time.Since(start)
	fmt.Printf("Sequential join took: %s\n", elapsed)

	start = time.Now()
	concurrentJoin(strList, concurrency)
	elapsed = time.Since(start)
	fmt.Printf("Concurrent join took: %s\n", elapsed)

	start = time.Now()
	concurrentJoinWithChannels(strList, concurrency)
	elapsed = time.Since(start)
	fmt.Printf("Concurrent join with channels took: %s\n", elapsed)
}

func generateStrings(count int) []string {
	strList := make([]string, count)
	for i := 0; i < count; i++ {
		strList[i] = randomString(20)
	}
	return strList
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func sequentialJoin(strList []string) {
	joinedString := strings.Join(strList, " ")
	fmt.Println(len(joinedString))
}

func concurrentJoin(strList []string, concurrency int) {
	var result string
	var wg sync.WaitGroup
	wg.Add(concurrency)

	strPerGoroutine := len(strList) / concurrency
	for i := 0; i < concurrency; i++ {
		start := i * strPerGoroutine
		end := (i + 1) * strPerGoroutine
		if end > len(strList) {
			end = len(strList)
		}
		go func(start, end int) {
			defer wg.Done()
			part := strings.Join(strList[start:end], " ")
			result += part
		}(start, end)
	}

	wg.Wait()
	fmt.Println(len(result))
}

func concurrentJoinWithChannels(strList []string, concurrency int) {
	var result string
	var wg sync.WaitGroup
	wg.Add(concurrency)

	strPerGoroutine := len(strList) / concurrency
	ch := make(chan string)

	for i := 0; i < concurrency; i++ {
		start := i * strPerGoroutine
		end := (i + 1) * strPerGoroutine
		if end > len(strList) {
			end = len(strList)
		}
		go func(start, end int) {
			defer wg.Done()
			part := strings.Join(strList[start:end], " ")
			ch <- part
		}(start, end)
	}

	go func() {
		for part := range ch {
			result += part
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(len(result))
}
