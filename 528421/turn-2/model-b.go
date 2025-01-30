package main

import (
	"fmt"
	"sync"
	"time"
)

type URLFetchJob struct {
	URL  string
	Err  error
	Data string // You can store the fetched data here if needed
}

func worker(id int, jobs <-chan URLFetchJob, results chan<- URLFetchJob, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d is fetching %s\n", id, job.URL)
		// Simulate network delay
		time.Sleep(2 * time.Second)

		// Add some error handling for demonstration purposes
		if job.URL == "http://example.com/c" {
			job.Err = fmt.Errorf("Simulated error fetching %s", job.URL)
		}
		results <- job
	}
}

func main() {
	const numWorkers = 3
	var wg sync.WaitGroup

	jobs := make(chan URLFetchJob, 100)
	results := make(chan URLFetchJob, 100)

	// Create worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	urls := []string{
		"http://example.com/a",
		"http://example.com/b",
		"http://example.com/c",
		"http://example.com/d",
		"http://example.com/e",
		// Add more URLs as needed
	}

	// Send jobs to the workers
	go func() {
		for _, url := range urls {
			jobs <- URLFetchJob{URL: url}
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf("Error fetching %s: %v\n", result.URL, result.Err)
		} else {
			fmt.Println("Completed fetching:", result.URL)
		}
	}

	fmt.Println("All URLs fetched")
}
