package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Callback is a function that will be called after a delay
type Callback func()

// ExecuteAfterDelay delays the execution of the provided callback by the specified duration.
// The wg is used to wait for this callback to complete.
func ExecuteAfterDelay(duration time.Duration, callback Callback, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		time.Sleep(duration)
		callback()
	}()
}

func callback1() {
	log.Println("Callback1 executed")
}

func callback2() {
	log.Println("Callback2 executed")
}

func main() {
	var wg sync.WaitGroup

	log.Println("Starting application...")

	// Add 1 to the wait group for each callback
	wg.Add(1)
	ExecuteAfterDelay(time.Second, callback1, &wg)

	wg.Add(1)
	ExecuteAfterDelay(time.Second*2, callback2, &wg)

	// Wait for all callbacks to complete
	wg.Wait()

	log.Println("Application completed.")
}
