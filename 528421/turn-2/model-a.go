package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function simulating HTTP request processing
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the WaitGroup counter is decremented when the worker exits
	for url := range jobs {
		fmt.Printf("Worker %d started fetching %s\n", id, url)
		// Simulate a network delay for the HTTP request
		time.Sleep(2 * time.Second)
		result := fmt.Sprintf("Worker %d finished fetching %s", id, url)
		fmt.Println(result)
		results <- result
	}
}

func main() {
	// Configuration for worker pool and jobs
	const numWorkers = 3
	var wg sync.WaitGroup

	// Channels to hold jobs and results
	jobs := make(chan string, 10)
	results := make(chan string, 10)

	// Launch workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// List of URLs to simulate fetching
	urls := []string{
		"http://example.com/a",
		"http://example.com/b",
		"http://example.com/c",
		"http://example.com/d",
		"http://example.com/e",
	}

	// Submit jobs to the jobs channel
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs) // Close the channel once all jobs are submitted
	}()

	// Collect results from the results channel
	go func() {
		wg.Wait()      // Wait for all workers to finish processing
		close(results) // Close the results channel
	}()

	// Print out all collected results
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All URLs have been processed")
}
