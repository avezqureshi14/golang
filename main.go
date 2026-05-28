package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TokenBuckets struct {
	capacity       int
	tokens         int
	refillRate     int
	lastRefillTime time.Time
	mutex          sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBuckets {
	return &TokenBuckets{
		capacity:       capacity,
		tokens:         capacity,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (t *TokenBuckets) refill() {
	elapsed := time.Since(t.lastRefillTime).Seconds()
	tokensToAdd := t.refillRate * int(elapsed)
	if tokensToAdd > 0 {
		t.tokens = min(t.capacity, t.tokens+tokensToAdd)
		t.lastRefillTime = time.Now()
	}
}

func (t *TokenBuckets) Allow() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.tokens > 0 {
		t.tokens--
		return true
	}
	return false
}

var limiter = NewTokenBucket(5, 2)

func rateLimitedHandler(w http.ResponseWriter, r *http.Request) {
	if !limiter.Allow() {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintln(w, "429 - too many requests")
	}

	fmt.Println(w, "Request successful")
}

func main() {

}
