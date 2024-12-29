package main

import (
	"sync"
	"testing"
	"time"
)

type Notification struct {
	ID      int
	Message string
	Time    time.Time
}

var (
	queue []Notification
	mu    = sync.Mutex{}
)

func producer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		notification := Notification{
			ID:      i + 1,
			Message: fmt.Sprintf("Notification %d at %v", i+1, time.Now()),
			Time:    time.Now(),
		}
		mu.Lock()
		queue = append(queue, notification)
		mu.Unlock()
	}
}

func consumer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if len(queue) > 0 {
			mu.Lock()
			notification := queue[0]
			queue = append(queue[:0], queue[1:]...)
			mu.Unlock()
			fmt.Println(notification)
		} else {
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func TestPerformance(t *testing.T) {
	t.Run("Producer", func(b *testing.B) {
		b.ResetTimer()
		producer(b)
	})

	t.Run("Consumer", func(b *testing.B) {
		b.ResetTimer()
		consumer(b)
	})
}
