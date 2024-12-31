package main

import (
	"sync"
	"testing"
)

func TestMaps(t *testing.T) {
	t.Run("Map", func(t *testing.T) {
		var m map[string]int = make(map[string]int)
		testAccess(t, &m)
	})
	t.Run("SyncMap", func(t *testing.T) {
		var m sync.Map
		testAccess(t, &m)
	})
}

func testAccess(t *testing.T, m interface{}) {
	numThreads := 10
	duration := time.Second
	incrementAmount := 1000000

	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(m interface{}) {
			defer wg.Done()
			for j := 0; j < incrementAmount; j++ {
				key := fmt.Sprintf("key%d", j%100)
				if _, ok := m.(map[string]int); ok {
					mm, _ := m.(map[string]int)
					for !ok {
						select {
						case <-time.After(time.Millisecond):
							mm, ok = m.(map[string]int)
						}
					}
					var val int
					if _, ok := mm[key]; ok {
						val = mm[key]
					} else {
						mm[key] = 0
					}
					val++
					mm[key] = val
				} else {
					var val int
					err := m.(sync.Map).LoadOrStore(key, &val)
					if err != nil {
						panic(err)
					}
					val++
					m.(sync.Map).Store(key, val)
				}
			}
		}(m)
	}

	wg.Wait()
}
