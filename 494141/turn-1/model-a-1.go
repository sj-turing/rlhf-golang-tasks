package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m map[string]int
	var wg sync.WaitGroup

	m = make(map[string]int)
	wg.Add(2)

	go incrementMap(&m, &wg)
	go incrementMap(&m, &wg)

	wg.Wait()
	fmt.Println(m)
}

func incrementMap(m *map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		(*m)[fmt.Sprintf("key%d", i%100)]++
	}
}
