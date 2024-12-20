package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: waiting for Goroutine 2")
		wg.Wait()
		fmt.Println("Goroutine 1: Done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: waiting for Goroutine 1")
		wg.Wait()
		fmt.Println("Goroutine 2: Done")
	}()

	wg.Wait() // This will deadlock
}
