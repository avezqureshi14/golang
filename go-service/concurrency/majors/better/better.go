package workerpool

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	ID     string
	Status string
}

// =========================
// PRODUCER (C, S, E)
// =========================
func generate(ctx context.Context, nums ...int) <-chan Job {
	jobs := make(chan Job)

	go func() {
		defer close(jobs) // Ownership: producer closes

		for _, num := range nums {
			select {
			case <-ctx.Done(): // Early Exit: stop if system cancels
				return
			case jobs <- Job{ID: strconv.Itoa(num), Status: "Pending"}:
			}
		}
	}()

	return jobs
}

// =========================
// WORKER (S, E, B)
// =========================
func dowork(
	ctx context.Context,
	wg *sync.WaitGroup,
	jobs <-chan Job,
	results chan<- Job,
	errCh chan<- error,
) {
	defer wg.Done()

	// Panic safety (Early Exit: fail-fast without deadlock)
	defer func() {
		if r := recover(); r != nil {
			select {
			case errCh <- fmt.Errorf("worker panic: %v", r):
			default: // Backpressure: don't block if error already sent
			}
		}
	}()

	for {
		select {
		case <-ctx.Done(): // Stop condition
			return

		case val, ok := <-jobs:
			if !ok {
				return // producer finished
			}

			// simulate work
			time.Sleep(100 * time.Millisecond)

			job := Job{ID: val.ID, Status: "Done"}

			// =========================
			// Early Exit + Backpressure
			// =========================
			select {
			case results <- job:
				// Normal flow

			case <-ctx.Done():
				// Early Exit: if consumer stopped or system cancelled
				return

			case <-time.After(200 * time.Millisecond):
				// Backpressure strategy:
				// If downstream is slow → fail fast instead of blocking forever
				select {
				case errCh <- fmt.Errorf("worker timeout sending result"):
				default:
				}
				return
			}
		}
	}
}

// =========================
// ORCHESTRATOR (C, O, E, B)
// =========================
func WorkerPool() {
	var wg sync.WaitGroup

	// Global lifecycle controller
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Error channel (signal channel, not stream)
	errCh := make(chan error, 1)

	jobs := generate(ctx, 1, 2, 3, 4, 5)

	// =========================
	// Backpressure Strategy:
	// Buffered channel allows burst,
	// but still blocks eventually → controlled pressure
	// =========================
	results := make(chan Job, 10)

	// Create workers
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go dowork(ctx, &wg, jobs, results, errCh)
	}

	// Close results AFTER all workers exit
	go func() {
		wg.Wait()
		close(results) // Ownership: orchestrator closes
	}()

	// =========================
	// CONSUMER LOOP (E, B)
	// =========================
	for {
		select {

		// Fail-fast path
		case err := <-errCh:
			fmt.Println("Error:", err)
			cancel() // Early Exit: propagate shutdown

			// Drain results to avoid goroutine leaks
			for range results {
			}
			return

		// Normal consumption
		case res, ok := <-results:
			if !ok {
				return // all done
			}
			fmt.Println(res.ID, res.Status)

			// Simulate early exit scenario
			// if res.ID == "3" {
			// 	cancel()
			// }

		// Global timeout / cancellation
		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())

			// Drain results to unblock workers
			for range results {
			}
			return
		}
	}
}


/*
1. Consumer exits early so for that we added context + select , now on top our producer will be listening ctx.Done() channel and when on failure our consumer will trigger cancel all producer will stop

2. We use recover inside workers to catch unexpected panics during job processing and convert them into errors. This error is propagated to the orchestrator (main()) , which cancels the context and stops the entire pipeline.

3. Same manner we have to use panic recover inside a producer and propogate errors through error channel so that orchestrator cancels it and stops the entire pipeline


4. Now when our application is taking more than expected time to complete in this case we will context.WithTimeout to stop the workers across the chain , now context.WithTimeout will lead to system wide shutdown

5. Now we have backpressure strategy , now this is used when producer are producing more data than consumer can cosume it , now in this case
	i.  first thing which we can do is using buffered channels
	ii. Another strategy is adding a timeout at the worker level. If downstream is slow, we can fail fast instead of blocking while sending results. Additionally, we can apply a timeout to individual job processing using context, so that a single slow job doesn’t block the worker indefinitely. This timeout is applied to the job execution, not to terminating the worker itself


1. Start minimal worker pool
2. Add context cancellation
3. Add error handling
4. Talk about backpressure
5. Talk about timeouts
6. Mention edge cases
*/