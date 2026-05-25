package main

import (
	"context"
	"fmt"
	"sync"
)

type Pair struct {
	a int
	b int
}

func getChunk(id, base, total, workers int) (start, end int) {
	chunkSize := total / workers
	start = base + id*chunkSize
	end = start + chunkSize

	if id == workers-1 {
		end = base + total
	}
	return
}

func worker(ctx context.Context, id int, out chan<- Pair, wg *sync.WaitGroup, n int, workers int) {
	defer wg.Done()

	start, end := getChunk(id, 0, n, workers)

	for i := start; i < end; i++ {
		for j := 0; j < n; j++ {
			select {
			case <-ctx.Done():
				return
			case out <- Pair{i, j}:
			}
		}
	}
}

// unchanged (your logic is good)
func twins(ctx context.Context, in <-chan Pair) <-chan map[int][]Pair {
	mp := make(chan map[int][]Pair)
	go func() {
		defer close(mp)
		storage := make(map[int][]Pair)

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					select {
					case mp <- storage:
					case <-ctx.Done():
					}
					return
				}
				sum := val.a + val.b
				storage[sum] = append(storage[sum], val)
			}
		}
	}()
	return mp
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n := 10
	workers := 4

	out := make(chan Pair)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ctx, i, out, &wg, n, workers)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	resultsChan := twins(ctx, out)

	for finalMap := range resultsChan {
		for sum, pairs := range finalMap {
			fmt.Printf("%d -> ", sum)
			for _, p := range pairs {
				fmt.Printf("{%d,%d} ", p.a, p.b)
			}
			fmt.Println()
		}
	}
}