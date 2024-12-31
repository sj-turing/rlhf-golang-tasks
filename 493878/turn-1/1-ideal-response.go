package main

import (
	"container/heap"
	"fmt"
	"log"
	"math/rand"
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
	return &CallbackQueue{tasks: make([]*Task, 0), lock: sync.Mutex{}}
}

// enqueueTask adds a task to the callback queue based on its priority
func (q *CallbackQueue) Push(item any) {
	// Heap implementation for priority queue
	q.lock.Lock()
	defer q.lock.Unlock()

	task, _ := item.(*Task)

	q.tasks = append(q.tasks, task)
}

// dequeueTask removes and returns the highest priority task from the queue
func (q *CallbackQueue) Pop() any {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.tasks) <= 0 {
		return nil
	}

	firstTask := q.tasks[0]
	q.tasks = q.tasks[1:]

	return firstTask
}

// GetQueueSize returns the current number of tasks in the queue
func (q *CallbackQueue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.tasks)
}

func (q *CallbackQueue) Less(i, j int) bool {
	if q == nil {
		return false
	}

	return (*q.tasks[i]).Priority < (*q.tasks[j]).Priority
}

// Swap swaps the elements with indexes i and j.
func (q *CallbackQueue) Swap(i, j int) {
	if len(q.tasks) <= 1 {
		return
	}
	q.tasks[i], q.tasks[j] = q.tasks[j], q.tasks[i]
}

func taskWorker(workerID int, queue *CallbackQueue) {
	for {
		item := heap.Pop(queue)
		if item == nil {
			time.Sleep(time.Second * 1) // Small sleep to avoid busy waiting
			continue
		}

		task, _ := item.(*Task)

		if err := task.Callback(); err != nil {
			log.Printf("Worker %d: Error executing task %d: %v", workerID, task.Id, err)
		}

		log.Printf("Worker %d: Completed task %d", workerID, task.Id)
	}
}

type TaskManager struct {
	queue      *CallbackQueue
	workers    int
	numWorkers int
}

func NewTaskManager(numWorkers int) *TaskManager {
	queue := NewCallbackQueue()
	heap.Init(queue)
	tm := &TaskManager{
		queue:      queue,
		workers:    numWorkers,
		numWorkers: numWorkers,
	}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go taskWorker(i, tm.queue)
	}

	return tm
}

// AddTask adds a task to the task manager
func (tm *TaskManager) AddTask(task *Task) {
	heap.Push(tm.queue, task)
}

func main() {
	// Create a task manager with 4 worker goroutines
	taskManager := NewTaskManager(4)

	for i := 0; i < 10; i++ {

		// Add tasks with different priorities
		taskManager.AddTask(&Task{
			Id:       i,
			Priority: rand.Intn(10),
			Payload:  fmt.Sprintf("High-priority task: %d", i),
			Callback: func() error {
				fmt.Printf("Executing high-priority task #%d\n", i)
				return nil
			},
		})

	}

	// Add more tasks...

	// Block the main thread until all workers are stopped
	select {}
}
