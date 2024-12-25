package main

import (
	"fmt"
	"log"
	"time"
)

// Callback is a function that will be called after a delay
type Callback func()

// ExecuteAfterDelay delays the execution of the provided callback by the specified duration
func ExecuteAfterDelay(duration time.Duration, callback Callback) {
	go func() {
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
	log.Println("Starting application...")

	ExecuteAfterDelay(time.Second, callback1)
	ExecuteAfterDelay(time.Second*2, callback2)

	// The program will exit before the callbacks have completed
	// To prevent this, we can add a wait group
	// For this example, we'll simply sleep for a while.
	time.Sleep(time.Second * 3)

	log.Println("Application completed.")
}
