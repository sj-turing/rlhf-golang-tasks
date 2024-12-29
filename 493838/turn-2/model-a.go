package main

import (
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	ID       int
	Message string
	Time     time.Time
}

var (
	queue []Notification
	mu    = sync.Mutex{}
)

func producer(id chan int, channel chan<- Notification) {
	defer func() {
		close(id)
	}()
	for i := 1; ; i++ {
		select {
		case <<-id:
			notification := Notification{
				ID:       i,
				Message: fmt.Sprintf("New notification %d at %v", i, time.Now()),
				Time:    time.Now(),
			}
			mu.Lock()
			queue = append(queue, notification)
			mu.Unlock()
			channel <- notification
			time.Sleep(time.Millisecond * 100)
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func consumer(id chan int, channel <-chan Notification) {
	defer func() {
		close(channel)
	}()
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
		case <-id:
			fmt.Println("Consumer is done processing all notifications")
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func findIndex(s []Notification, n Notification) int {
	for i, v := range s {
		if v.ID == n.ID {
			return i
		}
	}
	return -1
}

func main() {
	idChannel := make(chan int, 1)
	channel := make(chan Notification)

	go producer(idChannel, channel)
	go consumer(idChannel, channel)

	// Signal to stop after a period
	time.Sleep(5 * time.Second)
	close(idChannel)
}
