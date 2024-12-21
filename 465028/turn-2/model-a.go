package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// WorkerPool is a structure to manage a pool of workers
type WorkerPool struct {
	workers        int
	wg             sync.WaitGroup
	taskChan       chan func()
	shutdownSignal chan struct{}
}

// NewWorkerPool initializes a new worker pool with a given number of workers
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:        workers,
		taskChan:       make(chan func()),
		shutdownSignal: make(chan struct{}),
	}
}

// Start starts the worker pool
func (pool *WorkerPool) Start() {
	for i := 0; i < pool.workers; i++ {
		go pool.worker()
	}
}

// worker is the main function that each worker goroutine executes
func (pool *WorkerPool) worker() {
	for {
		select {
		case task, ok := <-pool.taskChan:
			if !ok {
				return // shutdown signal received
			}
			pool.wg.Add(1)
			task()
			pool.wg.Done()
		case <-pool.shutdownSignal:
			return // shutdown signal received
		}
	}
}

// Submit adds a task to the pool
func (pool *WorkerPool) Submit(task func()) {
	select {
	case pool.taskChan <- task:
	default:
		fmt.Println("[ERROR] Worker pool is full, unable to submit task")
	}
}

// Shutdown shuts down the worker pool
func (pool *WorkerPool) Shutdown() {
	close(pool.taskChan)
	pool.wg.Wait()
	close(pool.shutdownSignal)
}

// AsyncTask executes a task asynchronously in the worker pool
func AsyncTask(workerPool *WorkerPool, taskID int, taskName string) {
	startTime := time.Now()
	fmt.Printf("[DEBUG] %s-%d: Task started at %s\n", taskName, taskID, startTime)

	// Simulate some work
	time.Sleep(time.Duration(int64(taskID)%2) * time.Second)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	workerPool.Submit(func() {
		defer atomic.AddInt64(&completedTasks, 1)
		fmt.Printf("[INFO] %s-%d: Task completed in %s\n", taskName, taskID, executionTime)
	})
}

var completedTasks int64

func main() {
	var wg sync.WaitGroup
	numTasks := 20
	taskName := "SampleTask"

	// Initialize worker pool with 5 workers
	workerPool := NewWorkerPool(5)
	workerPool.Start()
	defer workerPool.Shutdown()

	fmt.Println("[INFO] Starting tasks in worker pool...")

	// Initialize wait group for main function
	wg.Add(numTasks)

	// Launch tasks concurrently
	for taskID := 1; taskID <= numTasks; taskID++ {
		go func(taskID int) {
			defer wg.Done()
			AsyncTask(workerPool, taskID, taskName)
		}(taskID)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Printf("[INFO] Total completed tasks: %d\n", atomic.LoadInt64(&completedTasks))
	fmt.Println("[INFO] All tasks have completed in the worker pool.")
}
