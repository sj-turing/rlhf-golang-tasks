package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Notification struct {
	Message string
	Time    time.Time
}

var (
	queue     []Notification
	mu        sync.Mutex
	channel   = make(chan Notification)
	wg        sync.WaitGroup
	isRunning       = true
	size      int32 = 0
)

func producer() {
	defer wg.Done()
	for isRunning {
		notification := Notification{
			Message: fmt.Sprintf("New notification at %v", time.Now()),
			Time:    time.Now(),
		}
		mu.Lock()
		queue = append(queue, notification)
		mu.Unlock()
		channel <- notification
		time.Sleep(time.Second)
	}
}

func consumer() {
	defer wg.Done()
	for isRunning {
		select {
		case notification := <-channel:
			fmt.Println(notification)
			// Process the notification
			atomic.AddInt32(&size, -1)
			mu.Lock()
			if index := findIndex(queue, notification); index != -1 {
				queue = append(queue[:index], queue[index+1:]...)
			}
			mu.Unlock()
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func findIndex(s []Notification, n Notification) int {
	for i, v := range s {
		if v.Time.Equal(n.Time) {
			return i
		}
	}
	return -1
}

func main() {
	wg.Add(2)
	go producer()
	go consumer()
	time.Sleep(10 * time.Second)
	isRunning = false
	wg.Wait()
	fmt.Println("Queue size:", atomic.LoadInt32(&size))
}
