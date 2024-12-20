package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// wg.Wait() will block indefinitely here if not called correctly
	start := time.Now()
	wg.Wait() // As no goroutines have been added, this blocks forever

	fmt.Printf("Program finished in %v\n", time.Since(start))
}
