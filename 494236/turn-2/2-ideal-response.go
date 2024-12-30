package main

import (
	"fmt"
	"time"
)

// processStep1 acts like performing some operations concurrently
func processStep1(data <-chan int, output chan<- int) {
	for value := range data {
		output <- value * 2
		fmt.Println("Step 1 processed value:", value)
		time.Sleep(60 * time.Millisecond)
	}
	close(output)
}

// processStep1 acts like performing some operations concurrently
func processStep2(data <-chan int, output chan<- int) {
	for value := range data {
		output <- value + 1
		fmt.Println("Step 2 processed value:", value)
		time.Sleep(50 * time.Millisecond)
	}
	close(output)
}

func main() {
	data := make(chan int, 100)
	step1Output := make(chan int)
	step2Output := make(chan int)

	go processStep1(data, step1Output)
	go processStep2(step1Output, step2Output)

	// dumpping data into data channel
	for i := 0; i < 100; i++ {
		data <- i
	}
	close(data) // closing data channel

	// after processing with 1st and 2nd step
	// finialized data will the there in the step2Output channel
	for result := range step2Output {
		fmt.Println("Final result:", result)
	}

}
