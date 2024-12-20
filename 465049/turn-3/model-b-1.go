package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var count int
	var wg sync.WaitGroup

	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}
