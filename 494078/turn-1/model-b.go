package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// define a task
type task struct {
	id   int
	data string
}

// define a worker function
func worker(taskCh <-chan task, resultCh chan<- task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskCh {
		fmt.Println("Processing task: ", task.id, task.data)
		// Simulate work with random sleep
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		// Send the result back to the main goroutine
		resultCh <- task
	}
}

// Simulate a concurrent database operation
func processDatabaseTask(task task) {
	fmt.Println("Database task: ", task.id)
	// Simulate work with random sleep
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func main() {
	numTasks := 20
	numWorkers := 5
	// Initialize the wait group
	var wg sync.WaitGroup
	wg.Add(numTasks)
	// Create channels for task and result
	taskCh := make(chan task, numTasks)
	resultCh := make(chan task, numTasks)
	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(taskCh, resultCh, &wg)
	}
	// Generate tasks and send them to the workers
	for i := 0; i < numTasks; i++ {
		task := task{id: i, data: fmt.Sprintf("Task data %d", i)}
		taskCh <- task
	}
	// Close the task channel to signal the workers that there are no more tasks
	close(taskCh)
	// Wait for all tasks to be completed
	wg.Wait()
	// Close the result channel to signal that there are no more results
	close(resultCh)
	// Process the results
	for result := range resultCh {
		processDatabaseTask(result)
	}
	fmt.Println("All tasks completed")
}
