package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// task represents a unit of work to be done asynchronously.
type task struct {
	id   int
	name string
}

// these will be used to see the completion status
var c chan bool
var count = 0

func work(wg *sync.WaitGroup, t task) {
	defer wg.Done()
	// Simulate work being done
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	// Now notify the main that the task is complete
	c <- true

	log.Printf("Task %d (%s) completed", t.id, t.name)
}

func main() {
	c = make(chan bool, 3)
	tasks := []task{
		{1, "Task A"},
		{2, "Task B"},
		{3, "Task C"},
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go work(&wg, task)
	}

	for i := 0; i < len(tasks); i++ {
		<-c
	}
	fmt.Println("All tasks completed")
	// Wait for all workers to finish
	wg.Wait()

	// Log completion of all tasks
	log.Println("All tasks have been processed")
}
