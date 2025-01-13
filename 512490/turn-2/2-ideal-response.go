package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	numStrings = 10
)

func main() {
	strList := generateStrings(numStrings)

	start := time.Now()
	sequentialJoin(strList)
	elapsed := time.Since(start)
	fmt.Printf("Sequential join took: %s\n", elapsed)

	start = time.Now()
	concurrentJoin(strList, len(strList))
	elapsed = time.Since(start)
	fmt.Printf("Concurrent join took: %s\n", elapsed)

	start = time.Now()
	concurrentJoinWithChannels(strList)
	elapsed = time.Since(start)
	fmt.Printf("Concurrent join with channels took: %s\n", elapsed)
}

func generateStrings(count int) []string {
	strList := make([]string, count)
	for i := 0; i < count; i++ {
		strList[i] = randomString(rand.Intn(20))
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

func sequentialJoin(strList []string) string {
	if len(strList) == 0 {
		return ""
	}

	result := strList[0]
	for i := 1; i < len(strList); i++ {
		result = fmt.Sprintf("%s %s", result, strList[i])
	}

	return result
}

func concurrentJoin(strList []string, batchSize int) string {
	if len(strList) == 0 {
		return ""
	}
	if len(strList) == 1 {
		return strList[0]
	}
	var result string
	var wg sync.WaitGroup

	totalConcurrency := len(strList) / batchSize

	wg.Add(totalConcurrency)

	for i := 0; i < totalConcurrency; i++ {
		start := batchSize * i
		end := batchSize * (i + 1)

		go func(start, end int) {
			defer wg.Done()
			part := strings.Join(strList[start:end], " ")
			result += part
		}(start, end)
	}
	wg.Wait()

	return result
}

func concurrentJoinWithChannels(strList []string) string {
	if len(strList) == 0 {
		return ""
	}

	if len(strList) == 1 {
		return strList[0]
	}

	var result string
	strChan := make(chan string, len(strList))

	var wg sync.WaitGroup
	wg.Add(len(strList))

	go func() {
		for _, str := range strList {
			defer wg.Done()
			strChan <- str
		}
	}()

	go func() {
		wg.Wait()
		close(strChan)
	}()

	for str := range strChan {
		result += str + " "
	}

	return result[:len(strList)-1]
}
