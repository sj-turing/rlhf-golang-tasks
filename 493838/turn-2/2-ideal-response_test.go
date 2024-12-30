// filename: main_test.go

package main

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkNotificationSystem(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
}
