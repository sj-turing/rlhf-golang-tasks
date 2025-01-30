package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate HTTP request duration
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("Fetched data from:", url)
}
func main() {
	urls := []string{"url1", "url2", "url3", "url4", "url5"}
	var wg sync.WaitGroup

	// Start fetching data from all URLs using goroutines
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg)
	}

	// Wait for all data fetching to complete
	wg.Wait()

	fmt.Println("Data fetching completed.")
}
