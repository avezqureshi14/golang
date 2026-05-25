package practice

import (
	"context"
	"fmt"
)

type Pair struct {
	a int
	b int
}

func generator(ctx context.Context) <-chan Pair {
	out := make(chan Pair)
	// Fixed: Added 'go' so this doesn't block the main thread
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				select {
				case <-ctx.Done():
					return
				case out <- Pair{i, j}:
				}
			}
		}
	}()
	return out
}

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
					// SEND ONLY ONCE HERE: When 'in' is closed, send the final result
					select {
					case mp <- storage:
					case <-ctx.Done():
					}
					return
				}
				sum := val.a + val.b
				storage[sum] = append(storage[sum], val)
				// Removed 'mp <- storage' from here to prevent repeating
			}
		}
	}()
	return mp
}

func Practice() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numbers := generator(ctx)
	resultsChan := twins(ctx, numbers)

	// This loop will now only run ONCE because 'twins' only sends one map
	for finalMap := range resultsChan {
		for sum, pairs := range finalMap {
			fmt.Printf("%d -> ", sum)
			for _, p := range pairs {
				fmt.Printf("{ %d, %d } ", p.a, p.b)
			}
			fmt.Println()
		}
	}
}
