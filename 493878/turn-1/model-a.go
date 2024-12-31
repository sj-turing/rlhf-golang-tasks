package main

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Function func(interface{}) error
	Args     interface{}
	Callback func(error)
	Priority int
	Status   string // e.g., "pending", "running", "complete"
}

var taskQueue []*Task
var queueLock sync.RWMutex

func generateTaskID() int {
	return int(rand.Int63())
}

func SubmitTask(function func(interface{}) error, args interface{}, callback func(error), priority int) {
	task := &Task{
		ID:       generateTaskID(),
		Function: function,
		Args:     args,
		Callback: callback,
		Priority: priority,
		Status:   "pending",
	}
	queueLock.Lock()
	taskQueue = append(taskQueue, task)
	sort.Slice(taskQueue, func(i, j int) bool {
		return taskQueue[i].Priority < taskQueue[j].Priority
	})
	queueLock.Unlock()
}

const numWorkers = 10 // Adjust based on needs

func startWorkers() {
	for i := 0; i < numWorkers; i++ {
		go worker()
	}
}

func worker() {
	for {
		task := &Task{}
		queueLock.RLock()
		if len(taskQueue) > 0 {
			task = taskQueue[0]
			taskQueue = taskQueue[1:]
		}
		queueLock.RUnlock()

		if task == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		task.Status = "running"
		err := task.Function(task.Args)
		if err != nil {
			// Implement retry logic if needed
			task.Status = "failed"
			task.Callback(err)
		} else {
			task.Status = "complete"
			task.Callback(nil)
		}
	}
}
