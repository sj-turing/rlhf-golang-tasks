package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1 sync.Mutex
	var mu2 sync.Mutex

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mu1.Lock()
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Goroutine 1 locks mu1")
		mu2.Lock()
		fmt.Println("Goroutine 1 locks mu2")
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		mu2.Lock()
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Goroutine 2 locks mu2")
		mu1.Lock()
		fmt.Println("Goroutine 2 locks mu1")
	}()

	wg.Wait()
}
