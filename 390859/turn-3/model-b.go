package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int // Maximum number of tokens in the bucket
	tokens     int // Current number of tokens
	refillRate int // Tokens to refill per tick
	ticker     *time.Ticker
	mu         *sync.Mutex
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	ctx, cancelFunc := context.WithCancel(context.Background())
	ticker := time.NewTicker(time.Second / time.Duration(refillRate))

	tb := &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		ticker:     ticker,
		mu:         &sync.Mutex{},
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}

	go tb.refillTokens()

	return tb
}

func (tb *TokenBucket) refillTokens() {
	defer tb.cancelFunc()
	for range tb.ticker.C {
		tb.mu.Lock()
		defer tb.mu.Unlock()
		tb.tokens = min(tb.tokens+tb.refillRate, tb.capacity)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (tb *TokenBucket) Consume(count int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	if tb.tokens < count {
		return false
	}

	tb.tokens -= count
	return true
}

func rateLimitedHandler(tb *TokenBucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !tb.Consume(1) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Request successful!")
	}
}

func main() {
	tokenBucket := NewTokenBucket(10, 10) // 10 requests per second

	http.HandleFunc("/api", rateLimitedHandler(tokenBucket))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
