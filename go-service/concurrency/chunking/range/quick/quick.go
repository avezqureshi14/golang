package main

import (
	"fmt"
	"sync"
	"time"
)

var MAX_INT = 10
var CONCURRENCY = 5

func printFunc(num int) {
	fmt.Println("Print", num)
}

// 🔥 reusable chunk function
func getChunk(id, base, total, workers int) (start, end int) {
	chunkSize := total / workers

	start = base + id*chunkSize
	end = start + chunkSize

	// last worker takes remainder
	if id == workers-1 {
		end = base + total
	}
	return
}

func main() {

	var wg sync.WaitGroup

	base := 0
	total := MAX_INT - base

	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			startTime := time.Now()

			startRange, endRange := getChunk(id, base, total, CONCURRENCY)

			for v := startRange; v < endRange; v++ {
				printFunc(v)
			}

			fmt.Printf(
				"worker %d [%d,%d) completed in %s\n",
				id, startRange, endRange, time.Since(startTime),
			)
		}(i)
	}

	wg.Wait()
}