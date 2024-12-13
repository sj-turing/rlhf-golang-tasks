package main

/*

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// FetchURL fetches the content from the given URL and sends the result to the results channel
func FetchURL(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // Notify that this goroutine is done
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Failed to fetch %s: %v", url, err)
		return
	}
	defer resp.Body.Close() // Ensure that the response body is closed

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Failed to read response from %s: %v", url, err)
		return
	}

	results <- fmt.Sprintf("Fetched %s: %d bytes", url, len(body))
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.reddit.com",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls)) // Buffered channel to hold results

	// Start goroutines for each URL
	for _, url := range urls {
		wg.Add(1) // Increment wait group counter
		go FetchURL(url, &wg, results)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(results) // Close the channel once all goroutines are done
	}()

	// Process results as they come in
	for result := range results {
		fmt.Println(result)
	}
}

*/
