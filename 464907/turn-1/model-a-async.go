package main

import (
	"fmt"
	"time"
)

func main() {
	// Asynchronous function that takes a callback
	asyncFunction(func() {
		fmt.Println("This is an asynchronous callback.")
	}, 2*time.Second)

	fmt.Println("Main function continues without waiting for callback.")
}

func asyncFunction(callback func(), delay time.Duration) {
	go func() {
		// Simulate some work with a delay
		time.Sleep(delay)
		fmt.Println("Doing some asynchronous work...")
		callback() // Call the callback
	}()
}
