package errogroup

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
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

// WORKER (now returns error instead of using errCh)
func dowork(ctx context.Context, jobs <-chan Job, results chan<- Job) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case val, ok := <-jobs:
			if !ok {
				return nil
			}

			// simulate work
			time.Sleep(100 * time.Millisecond)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case results <- Job{ID: val.ID, Status: "Done"}:
			}
		}
	}
}

func WorkerPool() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	jobs := generate(ctx, 1, 2, 3, 4, 5)

	results := make(chan Job, 10)

	// start workers
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		g.Go(func() error {
			return dowork(ctx, jobs, results)
		})
	}

	// close results when workers finish
	go func() {
		_ = g.Wait()
		close(results)
	}()

	// consumer
	for {
		select {
		case res, ok := <-results:
			if !ok {
				return
			}
			fmt.Println(res.ID, res.Status)

		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
			return
		}
	}

	// optional final error check
	// if err := g.Wait(); err != nil {
	// 	fmt.Println("worker error:", err)
	// }
}