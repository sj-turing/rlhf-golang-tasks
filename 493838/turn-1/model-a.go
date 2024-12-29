package main

import (
	"fmt"
	"runtime/debug"
	"sync"
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
	isRunning = true
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
	gcStats := debug.GCStats{}
	debug.ReadGCStats(&gcStats)

	wg.Add(2)
	go producer()
	go consumer()
	time.Sleep(10 * time.Second)
	isRunning = false

	fmt.Println(gcStats)
	wg.Wait()
}
