package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// ComputePrimes runs a heavy CPU loop finding prime numbers , if there was GC work clearing memory heavy data structures it would make concurrency struggle u won't see same faster timing even with convurrency
func ComputePrimes(n int) {
	count := 0
	for i := 2; count < n; i++ {
		isPrime := true
		// Check if i is divisible by any number up to its square root
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
}

// runSingleThreaded executes tasks sequentially on a single core
func runSingleThreaded(tasks int) time.Duration {
	runtime.GOMAXPROCS(1) // Force the Go runtime to use exactly 1 CPU core
	start := time.Now()

	for i := 0; i < tasks; i++ {
		ComputePrimes(5000)
	}

	return time.Since(start)
}

// runConcurrent executes tasks in parallel using all available cores
func runConcurrent(tasks int, cores int) time.Duration {
	runtime.GOMAXPROCS(cores) // Tune the runtime to use all available cores
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < tasks; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ComputePrimes(5000)
		}()
	}
	wg.Wait()

	return time.Since(start)
}

func main() {
	availableCores := runtime.NumCPU()
	tasks := availableCores * 2 // Ensure enough tasks to saturate all cores

	fmt.Printf("--- System Info ---\n")
	fmt.Printf("Available CPU Cores: %d\n\n", availableCores)

	fmt.Printf("--- Running Benchmarks (Total Tasks: %d) ---\n", tasks)

	// 1. Benchmark Single-Threaded (1 Core)
	fmt.Print("Running single-threaded... ")
	durationST := runSingleThreaded(tasks)
	fmt.Printf("Done in %v\n", durationST)

	// 2. Benchmark Concurrent (Tuned GOMAXPROCS)
	fmt.Printf("Running concurrent (GOMAXPROCS=%d)... ", availableCores)
	durationCC := runConcurrent(tasks, availableCores)
	fmt.Printf("Done in %v\n\n", durationCC)

	// 3. Performance Analysis Output
	speedup := float64(durationST) / float64(durationCC)
	fmt.Printf("--- Results Analysis ---\n")
	fmt.Printf("Concurrent execution is %.2fx faster than single-threaded.\n", speedup)
}
