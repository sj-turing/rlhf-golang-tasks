package main

import (
  "fmt"
  "strings"
)

// Define a callback type for processing an integer
type IntegerProcessor func(int) int

// Define another callback type for processing a string
type StringProcessor func(string) string

// Compose callbacks of type IntegerProcessor
func composeIntegerProcessors(processors ...IntegerProcessor) IntegerProcessor {
	return func(x int) int {
		for _, processor := range processors {
			x = processor(x)
		}
		return x
	}
}

// Simple integer processors
func addOne(x int) int {
	return x + 1
}

func multiplyByTwo(x int) int {
	return x * 2
}

// Compose callbacks of type StringProcessor
func composeStringProcessors(processors ...StringProcessor) StringProcessor {
	return func(s string) string {
		for _, processor := range processors {
			s = processor(s)
		}
		return s
	}
}

// Simple string processors
func uppercase(s string) string {
    return strings.ToUpper(s)
}

func reverse(s string) string {
    return string(s[::-1])
}


func main() {

	// Create a composed processor
	composedProcessor := composeIntegerProcessors(addOne, multiplyByTwo)

	// Use the composed processor
	result := composedProcessor(3) // Output: 8
	fmt.Println("Result:", result)


  // Create a composed string processor
  composedStringProcessor := composeStringProcessors(uppercase, reverse)

  // Use the composed string processor
  stringResult := composedStringProcessor("hello") // Output: OLLEH
  fmt.Println("String Result:", stringResult)

}

