package main

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Priority int
	Payload  interface{}
	Callback func(interface{}, error) error // Callback with results and error
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Task)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	*pq = old[0 : n-1]
	return old[n-1]
}

type TaskQueue struct {
	queue *PriorityQueue
	lock  sync.Mutex
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{queue: &PriorityQueue{}}
}

func (q *TaskQueue) Enqueue(task *Task) {
	q.lock.Lock()
	heap.Push(q.queue, task)
	q.lock.Unlock()
}

func (q *TaskQueue) Dequeue() *Task {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.Len() == 0 {
		return nil
	}
	return heap.Pop(q.queue).(*Task)
}

func (q *TaskQueue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(*q.queue)
}

type TaskManager struct {
	queue       *TaskQueue
	workerPool  *sync.WaitGroup
	shutdownCtx context.Context
	shutdownFn  context.CancelFunc
}

func NewTaskManager(numWorkers int) *TaskManager {
	queue := NewTaskQueue()
	tm := &TaskManager{queue: queue, workerPool: &sync.WaitGroup{}, shutdownCtx: context.Background(), shutdownFn: nil}

	for i := 0; i < numWorkers; i++ {
		go tm.startWorker(i)
	}

	return tm
}

func (tm *TaskManager) startWorker(id int) {
	for {
		task := tm.queue.Dequeue()
		if task == nil {
			if tm.shutdownFn != nil && tm.shutdownCtx.Err() != nil {
				return
			}
			time.Sleep(time.Millisecond * 100) // Avoid busy waiting
			continue
		}

		err := tm.processTask(task)
		if err != nil {
			log.Printf("Worker %d: Error processing task %d: %v", id, task.ID, err)
		}
	}
}

func (tm *TaskManager) processTask(task *Task) error {
	defer tm.workerPool.Done()

	log.Printf("Worker starting task %d", task.ID)

	result, err := tm.executeTask(task.Payload)

	log.Printf("Worker completed task %d", task.ID)

	return task.Callback(result, err)
}

func (tm *TaskManager) executeTask(payload interface{}) (interface{}, error) {
	// Placeholder for task execution logic
	time.Sleep(time.Second) // Simulate work
	return "Task completed successfully", nil
}

func (tm *TaskManager) Shutdown() {
	if tm.shutdownFn != nil {
		tm.shutdownFn()
	}
	tm.workerPool.Wait()
}

func main() {
	taskManager := NewTaskManager(4)

	defer taskManager.Shutdown()

	// Example tasks
	for i := 0; i < 10; i++ {
		task := &Task{
			ID:       i,
			Priority: rand.Intn(10),
			Payload:  fmt.Sprintf("Task %d payload", i),
			Callback: func(result interface{}, err error) error {
				log.Printf("Task callback: Result: %v, Error: %v", result, err)
				return nil
			},
		}
		taskManager.queue.Enqueue(task)
	}

	select {} // Block main thread to keep the application running
}
