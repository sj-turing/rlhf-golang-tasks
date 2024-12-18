package main

/*

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// FetchData fetches data from a given URL
func FetchData(ctx context.Context, url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Printf("Fetching %s cancelled\n", url)
		return
	default:
	}

	fmt.Printf("Fetching %s...\n", url)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate network latency

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", url, err)
		return
	}

	results <- string(data)
}

// ProcessData processes the fetched data
func ProcessData(results <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range results {
		fmt.Printf("Processing data:\n%s\n", data)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Simulate processing time
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	results := make(chan string)

	urls := []string{
		"https://example.com/data1",
		"https://example.com/data2",
		"https://example.com/data3",
		"https://example.com/data4",
	}

	// Start fetching goroutines
	for _, url := range urls {
		wg.Add(1)
		go FetchData(ctx, url, &wg, results)
	}

	// Start processing goroutine
	wg.Add(1)
	go ProcessData(results, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed.")
}
*/
