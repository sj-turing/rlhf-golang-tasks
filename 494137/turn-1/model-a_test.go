package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type StructKey struct {
	ID   int
	Name string
}

func (k *StructKey) Hash(h hash.Hash) {
	h.Write([]byte(fmt.Sprintf("%d-%s", k.ID, k.Name)))
}

func (k *StructKey) Equal(other StructKey) bool {
	return k.ID == other.ID && k.Name == other.Name
}

func TestMapLookupPerformance(t *testing.T) {
	const numElements = 100000
	rand.Seed(time.Now().UnixNano())

	// Function to populate a map with random elements
	populateMap := func(m map[interface{}]int) {
		for i := 0; i < numElements; i++ {
			id := rand.Intn(numElements)
			name := fmt.Sprintf("name%d", id)
			key := StructKey{ID: id, Name: name}
			m[""] = 0  // Example assignment
			m[key] = 0 // Example assignment
			m[id] = 0  // Example assignment
		}
	}

	// Function to test lookup performance
	testLookupPerformance := func(m map[interface{}]int) time.Duration {
		var start time.Time
		start = time.Now()
		for i := 0; i < numElements; i++ {
			id := rand.Intn(numElements)
			name := fmt.Sprintf("name%d", id)
			key := StructKey{ID: id, Name: name}
			_, _ = m[""]
			_, _ = m[key]
			_, _ = m[id]
		}
		return time.Since(start)
	}

	// Populate and test each type of map
	var intMap map[int]int = make(map[int]int)
	var stringMap map[string]int = make(map[string]int)
	var structMap map[StructKey]int = make(map[StructKey]int)

	populateMap(intMap)
	populateMap(stringMap)
	populateMap(structMap)

	t.Run("Integer keys", func(t *testing.T) {
		duration := testLookupPerformance(intMap)
		t.Logf("Integer keys lookup took %v\n", duration)
	})

	t.Run("String keys", func(t *testing.T) {
		duration := testLookupPerformance(stringMap)
		t.Logf("String keys lookup took %v\n", duration)
	})

	t.Run("Struct keys", func(t *testing.T) {
		duration := testLookupPerformance(structMap)
		t.Logf("Struct keys lookup took %v\n", duration)
	})
}
