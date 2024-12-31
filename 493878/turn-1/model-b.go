package main

import (
	"container/heap"
	"fmt"
	"log"
	"sync"
	"time"
)

type Task struct {
	Id       int
	Priority int
	Payload  interface{}
	Callback func() error
}

type CallbackQueue struct {
	tasks []*Task
	lock  sync.Mutex
}

func NewCallbackQueue() *CallbackQueue {
	return &CallbackQueue{tasks: make([]*Task, 0)}
}

// enqueueTask adds a task to the callback queue based on its priority
func (q *CallbackQueue) EnqueueTask(task *Task) {
	q.lock.Lock()
	defer q.lock.Unlock()
	// Heap implementation for priority queue
	heap.Push(q, task)
}

// dequeueTask removes and returns the highest priority task from the queue
func (q *CallbackQueue) DequeueTask() *Task {
	q.lock.Lock()
	defer q.lock.Unlock()
	// Heap implementation for priority queue
	return heap.Pop(q).(*Task)
}

// GetQueueSize returns the current number of tasks in the queue
func (q *CallbackQueue) GetQueueSize() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.tasks)
}

func taskWorker(workerID int, queue *CallbackQueue) {
	for {
		task := queue.DequeueTask()
		if task == nil {
			time.Sleep(time.Second * 1) // Small sleep to avoid busy waiting
			continue
		}

		if err := task.Callback(); err != nil {
			log.Printf("Worker %d: Error executing task %d: %v", workerID, task.Id, err)
		}

		log.Printf("Worker %d: Completed task %d", workerID, task.Id)
	}
}

type TaskManager struct {
	queue      *CallbackQueue
	workers    []chan bool
	numWorkers int
}

func NewTaskManager(numWorkers int) *TaskManager {
	tm := &TaskManager{
		queue:      NewCallbackQueue(),
		workers:    make([]chan bool, numWorkers),
		numWorkers: numWorkers,
	}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		tm.workers[i] = make(chan bool)
		go taskWorker(i, tm.queue)
	}

	return tm
}

// AddTask adds a task to the task manager
func (tm *TaskManager) AddTask(task *Task) {
	tm.queue.EnqueueTask(task)
}

// StopWorkers stops all worker goroutines
func (tm *TaskManager) StopWorkers() {
	for _, worker := range tm.workers {
		close(worker)
	}
}

func main() {
	// Create a task manager with 4 worker goroutines
	taskManager := NewTaskManager(4)
	defer taskManager.StopWorkers()

	// Add tasks with different priorities
	taskManager.AddTask(&Task{
		Id:       1,
		Priority: 1,
		Payload:  "High-priority task",
		Callback: func() error {
			fmt.Println("Executing high-priority task.")
			return nil
		},
	})

	// Add more tasks...

	// Block the main thread until all workers are stopped
	select {}
}
