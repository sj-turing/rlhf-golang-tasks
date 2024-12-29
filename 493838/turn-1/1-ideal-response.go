package main

import (
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	Message string
	Time    time.Time
}

var (
	queue []Notification
	mu    = sync.Mutex{}
)

func producer(wg *sync.WaitGroup, channel chan<- Notification, isRunningChan <-chan bool) {
	defer wg.Done()
forloop:
	for {
		select {
		case _, ok := <-isRunningChan:
			if !ok {
				break forloop
			}
		default:
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
}

func consumer(wg *sync.WaitGroup, channel <-chan Notification, isRunningChan <-chan bool) {
	defer wg.Done()
forloop:
	for {
		select {
		case notification := <-channel:
			fmt.Println(notification)
			// Process the notification
			mu.Lock()
			if index := findIndex(queue, notification); index != -1 {
				queue = append(queue[:index], queue[index+1:]...)
			}
			mu.Unlock()
		case _, ok := <-isRunningChan:
			if !ok {
				break forloop
			}
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
	isRunningChan := make(chan bool, 1)
	channel := make(chan Notification)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go producer(&wg, channel, isRunningChan)
	go consumer(&wg, channel, isRunningChan)
	time.Sleep(10 * time.Second)

	close(isRunningChan)
	wg.Wait()
}
