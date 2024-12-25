package main

import (
	"fmt"
	"log"
	"math"
)

// Simple function that accepts a callback
func processNumber(num float64, callback func(float64)) {
	result := num * num
	fmt.Println("Processing number:", num)
	callback(result)
}

// Example callback function
func logResult(result float64) {
	log.Println("Callback Result:", result)
}

func main() {
	fmt.Println("Starting program")
	processNumber(5.0, logResult)
	fmt.Println("Program finished")
}
