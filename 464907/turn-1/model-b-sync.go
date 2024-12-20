package main

import (
	"fmt"
	"time"
)

func syncTask() int {
	// Simulate a long-running task
	time.Sleep(2 * time.Second)
	fmt.Println("Sync task completed.")
	return 42
}

func main() {
	fmt.Println("Starting sync task...")
	result := syncTask()
	fmt.Println("Sync task result:", result)
	fmt.Println("Main thread terminated.")
}
