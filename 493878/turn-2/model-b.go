package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"container/heap"
)

// Task represents a single unit of work to be processed asynchronously
type Task struct {
	ID       int
	Priority int
	Payload  interface{}
	Callback func(ctx context.Context, payload interface{}) error
	Retries   int
}

// CallbackQueue is a priority queue using a heap data structure
type CallbackQueue struct {
	tasks []*Task
	lock  sync.Mutex
}

func NewCallbackQueue() *CallbackQueue {
	return &CallbackQueue{tasks: make([]*Task, 0)}
}

// Implement heap.Interface to use the heap package for priority queue
func (q *CallbackQueue) Len() int           { return len(q.tasks) }
func (q *CallbackQueue) Less(i, j int) bool { return q.tasks[i].Priority < q.tasks[j].Priority }
func (q *CallbackQueue) Swap(i, j int)      { q.tasks[i], q.tasks[j] = q.tasks[j], q.tasks[i] }
func (q *CallbackQueue) Push(x interface{}) {
	q.tasks = append(q.tasks, x.(*Task))
}
func (q *CallbackQueue) Pop() interface{} {
	old := q.tasks
	n := len(old)
	x := old[n-1]
	q.tasks = old[0 : n-1]
	return x
}

// TaskManager manages the overall task processing
type TaskManager struct {
	queue      *CallbackQueue
	workers    int
	wg         sync.WaitGroup
	cancelFunc context.CancelFunc
}

// NewTaskManager creates a new TaskManager with the specified number of workers
func NewTaskManager(numWorkers int) *TaskManager {
	ctx, cancel := context.WithCancel(context.Background())
	tm := &TaskManager{
		queue:      NewCallbackQueue(),
		workers:    numWorkers,
		cancelFunc: cancel,
	}

	// Start worker goroutines
	tm.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go tm.taskWorker(ctx, i)
	}

	return tm
}

// AddTask adds a new task to the queue with a given priority
func (tm *TaskManager) AddTask(task *Task) {
	heap.Push(tm.queue, task)
}

// Stop gracefully shuts down the TaskManager, waiting for all workers to finish processing tasks
func (tm *TaskManager) Stop() {
	tm.cancelFunc()
	tm.wg.Wait()
}

func (tm *TaskManager) taskWorker(ctx context.Context, workerID int) {
	defer tm.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		item := heap.Pop(tm.queue)
		if item == nil {
			time.Sleep(time.Second) // Small sleep to avoid busy waiting
			continue
		}

		task := item.(*Task)

		if err := tm.executeTask(ctx, task); err != nil {
			fmt.Printf("Worker %d: Error executing task %d: %v\n", workerID, task.ID, err)
			// Retry task if it has remaining retries
			if task.Retries > 0 {
				task.Retries--
				heap.Push(tm.queue, task)
