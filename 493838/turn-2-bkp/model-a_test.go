package main

import (
	"testing"
	"time"
)

func BenchmarkAddNotifications(b *testing.B) {
	for i := 0; i < b.N; i++ {
		notification := Notification{
			Message: fmt.Sprintf("Benchmark notification %d at %v", i, time.Now()),
			Time:    time.Now(),
		}
		mu.Lock()
		queue = append(queue, notification)
		mu.Unlock()
	}
}
