package main

import (
	"fmt"
)

func main() {
	// Synchronous function that takes a callback
	syncFunction(func() {
		fmt.Println("This is a synchronous callback.")
	})
	fmt.Println("Main function continues after callback.")
}

func syncFunction(callback func()) {
	// Simulate some work
	fmt.Println("Doing some synchronous work...")
	callback() // Call the callback
}
