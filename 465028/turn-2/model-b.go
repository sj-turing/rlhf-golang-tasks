package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Define a task struct that holds task details and a channel for results
type Task struct {
	ID         int
	Name       string
	ResultChan chan string
}

func worker(wg *sync.WaitGroup, taskChan <-chan Task) {
	defer wg.Done()

	for task := range taskChan {
		startTime := time.Now()
		log.Printf("[DEBUG] Worker %d: Starting task %d (%s)", wg.Waiters(), task.ID, task.Name)

		// Simulate work with random duration
		duration := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(duration)

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		result := fmt.Sprintf("Task %d completed in %s", task.ID, executionTime)

		// Send the result back to the main thread
		task.ResultChan <- result
		log.Printf("[INFO] Worker %d: Task %d (%s) completed in %s", wg.Waiters(), task.ID, task.Name, executionTime)
	}
}

func main() {
	numWorkers := 5
	taskChan := make(chan Task)
	var wg sync.WaitGroup

	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.Println("[INFO] Starting worker pool...")

	// Start worker goroutines
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(&wg, taskChan)
	}

	// Generate and dispatch tasks
	numTasks := 20
	for taskID := 1; taskID <= numTasks; taskID++ {
		task := Task{
			ID:         taskID,
			Name:       fmt.Sprintf("Task-%d", taskID),
			ResultChan: make(chan string),
		}
		taskChan <- task
	}

	// Close the task channel to signal workers to exit
	close(taskChan)

	// Wait for all tasks to complete
	wg.Wait()

	log.Println("[INFO] All tasks have completed.")

	// Collect and log results
	for i := 1; i <= numTasks; i++ {
		task := Task{ID: i}
		result := <-task.ResultChan
		log.Printf("[INFO] Result for Task %d: %s", task.ID, result)
	}
}
