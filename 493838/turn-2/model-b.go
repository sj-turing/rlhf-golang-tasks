package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Notification struct {
	Message string
	Time    time.Time
}

var (
	queues      = make(map[int][]Notification)
	mu          = sync.Mutex{}
	numQueues   = 100
	queueSize   = 1000
	notifPerSec = 1000
)

func init() {
	// Initialize queues with some data
	for i := 0; i < numQueues; i++ {
		queues[i] = make([]Notification, queueSize)
	}
}

func generateNotifications(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second / time.Duration(notifPerSec))
	defer ticker.Stop()
	for range ticker.C {
		randQueueID := rand.Intn(numQueues)
		notification := Notification{
			Message: fmt.Sprintf("Notification for queue %d at %v", randQueueID, time.Now()),
			Time:    time.Now(),
		}
		mu.Lock()
		queues[randQueueID] = append(queues[randQueueID], notification)
		mu.Unlock()
	}
}

func consumeNotifications(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(1 * time.Millisecond)
		for queueID, queue := range queues {
			mu.Lock()
			if len(queue) > 0 {
				fmt.Println("Consuming:", queue[0])
				queues[queueID] = queue[1:]
			}
			mu.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go generateNotifications(&wg)
	go consumeNotifications(&wg)

	wg.Wait()
}
