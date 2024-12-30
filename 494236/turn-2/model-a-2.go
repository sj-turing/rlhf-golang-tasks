package main

import (
	"fmt"
	"sync"
	"time"
)

func processStep1(data chan int, output chan int) {
	for value := range data {
		output <- value * 2
		fmt.Println("Step 1 processed value:", value)
		time.Sleep(100 * time.Millisecond)
	}
	close(output)
}

func processStep2(data chan int, output chan int, wg *sync.WaitGroup) {
	for value := range data {
		output <- value + 1
		fmt.Println("Step 2 processed value:", value)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
	close(output)
}

func main() {
	data := make(chan int, 100)
	step1Output := make(chan int)
	step2Output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go processStep1(data, step1Output)
	go processStep2(step1Output, step2Output, &wg)

	for i := 0; i < 100; i++ {
		data <- i
	}
	close(data)

	wg.Wait()

	for result := range step2Output {
		fmt.Println("Final result:", result)
	}
}
