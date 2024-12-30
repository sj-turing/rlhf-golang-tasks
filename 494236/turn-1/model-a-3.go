package main

import (
	"fmt"
	"time"
)

var cache = make(map[int]string)

func fetchFromCache(key int) (string, bool) {
	return cache[key], cache[key] != ""
}

func storeInCache(key int, value string) {
	cache[key] = value
}

func processWithCaching(value int) {
	cachedValue, inCache := fetchFromCache(value)
	if inCache {
		fmt.Println("Cached: Processing value:", value, "with cache:", cachedValue)
	} else {
		fmt.Println("Not cached: Processing value:", value)
		time.Sleep(100 * time.Millisecond)
		storeInCache(value, fmt.Sprintf("Enriched value %d", value))
	}
}

func main() {
	data := make([]int, 100)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	for _, value := range data {
		processWithCaching(value)
	}
}
