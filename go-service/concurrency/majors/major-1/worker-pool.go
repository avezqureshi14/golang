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

// PRODUCER
func generate(ctx context.Context, nums ...int) <-chan Job {
	jobs := make(chan Job)

	go func() {
		defer close(jobs)

		for _, num := range nums {
			select {
			case <-ctx.Done():
				return
			case jobs <- Job{ID: strconv.Itoa(num), Status: "Pending"}:
			}
		}
	}()

	return jobs
}

// WORKER
func dowork(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Job, errCh chan<- error) {
	defer wg.Done()

	// If worker fails or producer crashes : failure propogation
	// panic safety (fail-fast)
	defer func() {
		if r := recover(); r != nil {
			errCh <- fmt.Errorf("worker panic: %v", r)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case val, ok := <-jobs:
			if !ok {
				return
			}

			// simulate work
			time.Sleep(100 * time.Millisecond)

			select {
			// this ctx done  will help us when consumer stop early , so when consumer stop it will cancel form the context1e
			case <-ctx.Done():
				return
			case results <- Job{ID: val.ID, Status: "Done"}:
			}
		}
	}
}

func WorkerPool() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	errCh := make(chan error, 1) //keep equal to the number of workers or use error group context

	jobs := generate(ctx, 1, 2, 3, 4, 5)

	// buffered to avoid blocking workers
	results := make(chan Job, 10)

	// start workers
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go dowork(ctx, &wg, jobs, results, errCh)
	}

	// close results AFTER workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// consumer loop
	for {
		select {
		case err := <-errCh:
			fmt.Println("Error:", err)
			cancel() // fail-fast
			return

		case res, ok := <-results:
			if !ok {
				return
			}
			fmt.Println(res.ID, res.Status)

			// If consumer stops early
			// simulate early consumer stop (optional test)
			// if res.ID == "3" {
			// 	cancel()
			// 	return
			// }

		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
			return
		}
	}
}
