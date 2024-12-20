package main

import (
	"fmt"
	"sync"
)

func main() {
	var sharedCounter int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go increment(&sharedCounter, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Final shared counter:", sharedCounter)
}

func increment(sharedCounter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		mu.Lock() // Lock the shared variable
		*sharedCounter++
		mu.Unlock() // Unlock the shared variable
	}
}
