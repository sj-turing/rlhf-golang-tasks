package main

import (
	"fmt"
	"sync"
	"time"
)

func processConcurrently(value int, wg *sync.WaitGroup) {
	fmt.Println("Processing value:", value)
	time.Sleep(100 * time.Millisecond)
	wg.Done()
}

func main() {
	data := make([]int, 100)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	var wg sync.WaitGroup
	for _, value := range data {
		wg.Add(1)
		go processConcurrently(value, &wg)
	}

	wg.Wait()
}
