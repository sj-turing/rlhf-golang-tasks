package main

import (
	"fmt"
	"sync"
	"time"
)

type rateLimiter struct {
	limit int
	burst int
	count int
	last  time.Time
	mu    sync.Mutex
}

func (rl *rateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if rl.count >= rl.limit && time.Since(rl.last) < time.Second {
		return false
	}
	if rl.count > 0 && time.Since(rl.last) >= time.Second {
		rl.count = 0
	}
	rl.count++
	rl.last = time.Now()
	return true
}
func main() {
	limiter := &rateLimiter{limit: 5, burst: 5}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 200)
		allowed := limiter.Allow()
		fmt.Println("Request", i, "allowed:", allowed)
	}
}
