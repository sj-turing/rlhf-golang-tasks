package main

import (
	"sync"
	"testing"
	"time"
)

// testCounter will be global but it is initialized using variable with initializers
var testCounter = func() *int {
	var counter int
	return &counter
}()

// this function increments the testCounter variable by n times
func incrementCounter(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		*testCounter++
	}
}

// We will use the table-driven test cases here
func TestIncrementCounter(t *testing.T) {
	cases := []struct {
		name      string
		increment int
		expected  int
	}{
		{
			name:      "Increment by 100",
			increment: 100,
			expected:  100,
		},
		{
			name:      "Increment by 1000 in 5 goroutines",
			increment: 1000,
			expected:  5000,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			*testCounter = 0 // Reset the counter for each test case
			var wg sync.WaitGroup
			wg.Add(tc.increment / 100) // We will use 5 goroutines to increment
			for i := 0; i < tc.increment; i += 100 {
				go incrementCounter(100, &wg)
			}
			wg.Wait()
			// Assertion to verify the correctness of the counter
			if *testCounter != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, *testCounter)
			}
		})
	}
}

func TestGarbageCollectionImpact(t *testing.T) {
	const iterations = 100000
	*testCounter = 0 // Reset the counter
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			*testCounter++
		}
	}()
	// Simulate work
	time.Sleep(1 * time.Second)
	wg.Wait()
	// Assertion to verify the correctness of the counter after iterations
	if *testCounter != iterations {
		t.Errorf("Expected: %d, Got: %d", iterations, *testCounter)
	}
}
