package main

import (
	"sync"
	"testing"
)

func BenchmarkSlicesNotificationSystem(b *testing.B) {
	numNotifications := b.N
	var wg sync.WaitGroup
	wg.Add(2)
	go generateNotifications(&wg)
	go consumeNotifications(&wg)
	wg.Wait()
}
