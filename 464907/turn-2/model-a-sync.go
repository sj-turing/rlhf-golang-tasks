package main

import (
	"fmt"
	"time"
)

func syncTask() (int, error) {
	// Simulate a long-running task with potential error
	time.Sleep(2 * time.Second)
	if err := someSimulatedError(); err != nil {
		return 0, err
	}
	fmt.Println("Sync task completed without error.")
	return 42, nil
}

func someSimulatedError() error {
	if random := randInt(0, 10); random < 5 {
		return fmt.Errorf("Simulated error: %d", random)
	}
	return nil
}

func randInt(min, max int) int {
	return min + int(rand.Intn(max-min+1))
}

func main() {
	fmt.Println("Starting sync task...")

	result, err := syncTask()
	if err != nil {
		fmt.Println("Error from sync task:", err)
	} else {
		fmt.Println("Sync task result:", result)
		fmt.Println("Main thread terminated without error.")
	}
}
