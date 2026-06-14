package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id     int
	status string
}

// 1. so for work pool pattern we see it from most
// producer -> consumer-1 consumer-2

func generate(ctx context.Context) <-chan Job {
	out := make(chan Job)
	// 2. the channel should be closed by the one who knows when values will stop entering the channel
	go func() {
		defer close(out)
		for i := 0; i <= 100; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- Job{id: i, status: "COMPLETED"}:
			}
		}
	}()
	return out
}

// 5. now lets come towards making our workers
func worker(ctx context.Context, wg *sync.WaitGroup, results chan Job, input <-chan Job, errChn chan interface{}) {
	defer wg.Done()

	// 11. now lets say worker fails or some error occurs  and we want that failure of this worker must trigger some action it can be shutting down of entire system or something else
	defer func() {
		err := recover()
		if err != nil {
			errChn <- err
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-input:
			// 7. over here this not ok means producer have finished and closed the channel so no jobs will be coming through it
			if !ok {
				return
			}

			// 13. now another strategy of backpressure can be of having timeout in the processing stuffs which worker is doing 
			// so that those processing dosen't blocks the workers and our downstream systems are non blocking when our producer is producing more data  , so in this case we can use context with timeout and similar stuffs
			// process some fake work and push the data in results channel
			time.Sleep(5 * time.Millisecond)
			results <- Job{id: val.id, status: "COMPLETED"}

			if val.id == 4 {
				panic("worker failed")
			}
		}
	}
}

func main() {

	// 10. now to make sure our workers stops producing data taken by the consumer results , when consumer stop early or fails
	// so for this we will be introducing context
	ctx, cancel := context.WithCancel(context.Background())

	// 3. so over here generate function will return a channel which will be used as an input for the workers
	input := generate(ctx)

	// 4. lets make the channel for storing results which will be accumulated from worker pool
	// 12. now lets say if our producer are producing more than what our consumers can consume in this case we will add buffer to 
	// our results channel 
	results := make(chan Job,5)

	// 6. so over here we will be spinning multiple workers , so we need to keep track of all worker before closing the results
	// channel and also for blocking the main function , for this reason we will be introducing waitgroups for our worker go routines
	var wg sync.WaitGroup

	// 12. so to propogate the errors throughout the system when any failure occurs we will be using a error channel which will
	// be signaled when any error occurs in the system
	errCh := make(chan interface{})

	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, results, input,errCh)
	}

	// 8. Now we need to make sure that the results channel is closed after all go routines have finished putting that their data in it , and we saw above that channel will be closed by someone who knows when go routines are going to get completed , so now in our case wg.Wait() knows when all go routines will be finised and soon after than we will close the channel
	go func() {
		wg.Wait()
		close(results)
	}()

	for {
		select {
		case <-errCh:
			// now over here we will listen to errors and will trigger system wide shutdown
			cancel()
		// 9. now whatever results we have got we will print them out
		case val, ok := <-results:
			if !ok {
				return
			}
			if val.id == 3 {
				// 11. so during such time when we need to do early exit, we need to stop the entire chain of workers and producers
				cancel()
			}
			fmt.Println(val.id, val.status)
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