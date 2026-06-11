package main

import (
	"context"
	"fmt"
	"sync"
)

type Job struct {
	id     int
	status string
}

func generate() <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for i := 0; i < 99; i++ {
			out <- Job{id: i, status: "PENDING"}
		}
	}()
	return out
}

func producer(ctx context.Context, wg *sync.WaitGroup, input <-chan Job, results chan Job) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-input:
			if !ok {
				return
			}
			results <- Job{id: val.id, status: "COMPLETED"}

		}
	}

}

func main() {
	input := generate()
	results := make(chan Job)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go producer(ctx, &wg, input, results)
	wg.Add(1)
	go producer(ctx, &wg, input, results)
	go func() {
		wg.Wait()
		close(results)
	}()
	for val := range results {
		fmt.Printf("Job Id %d status is %s \n", val.id, val.status)
	}

	for {
		select {
		case val, ok := <-results:
			if !ok {
				return
			}
			fmt.Println(val.id, val.status)

			// simulate early exit scenario
			if val.id == 3 {
				cancel()
			}

		// also listening from consumer side
		case <-ctx.Done():
			fmt.Println("Context cancellation ", ctx.Err())
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
