package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	// Simulate HTTP request
	time.Sleep(time.Duration(randomDuration()) * time.Second)

	// Prepare a sample response for the example
	response := fmt.Sprintf("Fetched %s", url)
	results <- response
}

func randomDuration() int {
	return int(rand.Intn(5) + 1) // Random duration between 1 and 5 seconds
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string)

	urls := []string{
		"https://example.com/1",
		"https://example.com/2",
		"https://example.com/3",
		"https://example.com/4",
		"https://example.com/5",
	}

	// Fan-Out: Start goroutines for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg, results)
	}

	// Fan-In: Collect results
	fmt.Println("Fetching URLs concurrently...")
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All URLs fetched.")
}
