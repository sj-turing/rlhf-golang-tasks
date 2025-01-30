package main

import (
	"fmt"
	"sync"
	"time"
)

// fetchURL simulates fetching data from a URL with a delay
func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Fetching URL: %s\n", url)
	// Simulate network delay
	time.Sleep(2 * time.Second)
	fmt.Printf("Completed fetching URL: %s\n", url)
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"http://example.com/a",
		"http://example.com/b",
		"http://example.com/c",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	// Wait for all fetch operations to complete
	wg.Wait()
	fmt.Println("All URLs fetched")
}
