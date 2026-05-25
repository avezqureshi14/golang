package ratelimiter

import (
	"fmt"
	"time"
)
type RateLimiter struct {
	tokens chan struct{}
}
func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, rate),
	}

	// fill initial tokens
	for i := 0; i < rate; i++ {
		rl.tokens <- struct{}{}
	}

	// refiller goroutine
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
				// added token
			default:
				// bucket full, drop token
			}
		}
	}()

	return rl
}
func (rl *RateLimiter) Allow() {
	<-rl.tokens
}
func Ratelimiter() {
	rl := NewRateLimiter(3, time.Second)

	for i := 1; i <= 10; i++ {
		go func(id int) {
			rl.Allow()
			fmt.Println("Request processed:", id, time.Now().Format("15:04:05"))
		}(i)
	}

	time.Sleep(5 * time.Second)
}