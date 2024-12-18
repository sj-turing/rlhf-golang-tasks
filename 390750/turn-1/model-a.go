package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		// Simulating a time-consuming task
		result := n * n
		fmt.Printf("Worker %d processed job: %d to result: %d\n", id, n, result)
		results <- result
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start 3 worker goroutines (fan-out)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs (e.g., numbers to be processed)
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs) // Close jobs channel to signal no more jobs
	}()

	// Close results channel after all workers are done processing
	go func() {
		wg.Wait()
		close(results) // Close results channel when done
	}()

	// Collect results (fan-in)
	fmt.Println("Results:")
	for result := range results {
		fmt.Println(result)
	}
}
