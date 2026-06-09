package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// doRequest performs HTTP call with retry + backoff
func doRequest(ctx context.Context, url string, maxRetries int) (*http.Response, error) {
	baseDelay := 100 * time.Millisecond

	client := &http.Client{}

	for attempt := 0; attempt <= maxRetries; attempt++ {

		// Create request with context timeout
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)

		// ✅ Success case
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}

		// ❌ If last attempt, return error
		if attempt == maxRetries {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("failed with status: %d", resp.StatusCode)
		}

		// ⏳ Exponential backoff with jitter
		backoff := baseDelay * time.Duration(1<<attempt) // 2^attempt
		jitter := time.Duration(rand.Int63n(int64(backoff / 2)))
		sleep := backoff + jitter

		fmt.Printf("retry %d: sleeping %v\n", attempt+1, sleep)

		select {
		case <-time.After(sleep):
			// continue retry
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, fmt.Errorf("unreachable")
}

func main() {
	// ⏱️ Global timeout for entire operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := doRequest(ctx, "https://httpstat.us/500", 20)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("success:", resp.Status)
}
