package main

/*

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	myMap := make(map[int]int)
	var wg sync.WaitGroup
	numGoroutines := 1000
	numOps := 10000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numOps; j++ {
				myMap[j] = j
			}
			wg.Done()
		}()
	}

	start := time.Now()
	wg.Wait()
	end := time.Now()

	fmt.Printf("Execution time: %v\n", end.Sub(start))
	fmt.Printf("Map size: %d\n", len(myMap))
}
*/
