package main

import (
	"fmt"
	"sync"
)

const totalGoroutines = 5
const incrementValue = 100000

func incrementCounter(counter *int, wg *sync.WaitGroup) {
	for i := 0; i < incrementValue; i++ {
		*counter++
	}
	wg.Done()
}
func main() {
	var counter int
	var wg sync.WaitGroup
	// Create a WaitGroup to wait for all goroutines to finish
	wg.Add(totalGoroutines)
	for i := 0; i < totalGoroutines; i++ {
		go incrementCounter(&counter, &wg)
	}
	// Wait for all goroutines to complete their execution
	wg.Wait()
	// Print the final value of the counter
	fmt.Println("Final value of counter:", counter)
}
