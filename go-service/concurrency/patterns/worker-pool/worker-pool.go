package workerpool

import (
	"fmt"
	"sync"
)

type Job struct {
	id int
}

type Result struct {
	jobID int
	sum   int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// simulate work
		sum := job.id * 2

		fmt.Printf("worker %d processed job %d\n", id, job.id)

		results <- Result{
			jobID: job.id,
			sum:   sum,
		}
	}
}

func WorkerPool() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// start workers (fan-out)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{id: i}
	}
	close(jobs)

	// close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// consume results (fan-in)
	for res := range results {
		fmt.Printf("result: job %d -> %d\n", res.jobID, res.sum)
	}
}

func main() {
	WorkerPool()
}