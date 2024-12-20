package main

import (
	"fmt"
	"time"
)

func asyncTask(done chan bool) {
	// Simulate a long-running task
	time.Sleep(2 * time.Second)
	fmt.Println("Async task completed.")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Starting async task...")
	go asyncTask(done)
	// Do other tasks while waiting for the async task to complete
	fmt.Println("Doing other tasks...")
	<-done
	fmt.Println("Main thread terminated.")
}
