package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity       int
	tokens         int
	refillRate     int
	lastRefillTime time.Time
	mutex          sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:       capacity,
		tokens:         capacity,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTime).Seconds()

	tokensToAdd := int(elapsed * float64(tb.refillRate))
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefillTime = now
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// 🔥 Global limiter (for now)
var limiter = NewTokenBucket(5, 2)

func rateLimitedHandler(w http.ResponseWriter, r *http.Request) {
	if !limiter.Allow() {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintln(w, "429 - Too Many Requests")
		return
	}

	fmt.Fprintln(w, "Request Successful")
}

func main() {
	http.HandleFunc("/api", rateLimitedHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
