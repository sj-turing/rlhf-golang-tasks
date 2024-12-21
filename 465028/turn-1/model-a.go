package main

import (
	"fmt"
	"sync"
	"time"
)

func asyncTask(wg *sync.WaitGroup, taskID int, taskName string) {
	defer wg.Done()
	startTime := time.Now()

	fmt.Printf("[DEBUG] %s-%d: Task started at %s\n", taskName, taskID, startTime)

	// Simulate some work
	time.Sleep(time.Duration(int64(taskID)%2) * time.Second)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("[INFO] %s-%d: Task completed in %s\n", taskName, taskID, executionTime)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 10
	taskName := "SampleTask"

	fmt.Println("[INFO] Starting tasks...")

	// Initialize wait group
	wg.Add(numTasks)

	// Launch tasks concurrently
	for taskID := 1; taskID <= numTasks; taskID++ {
		go asyncTask(&wg, taskID, taskName)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("[INFO] All tasks have completed.")
}
