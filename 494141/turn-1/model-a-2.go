package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	wg.Add(2)

	go incrementConcurrentMap(&m, &wg)
	go incrementConcurrentMap(&m, &wg)

	wg.Wait()
	fmt.Println(m)
}

func incrementConcurrentMap(m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("key%d", i%100)
		var val int
		err := m.LoadOrStore(key, &val)
		if err != nil {
			panic(err)
		}
		// Increment value
		val++
		m.Store(key, val)
	}
}
