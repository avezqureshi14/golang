package practice

import (
	"context"
	"fmt"
	"sync"
)

func generator(wg *sync.WaitGroup, ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
		defer close(out)
		for _, val := range nums {
			select {
			case <-ctx.Done():
			case out <- val:
			}
		}
	}()
	return out
}

func worker(wg *sync.WaitGroup, ctx context.Context, in <-chan int, num int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
		case val, ok := <-in:
			if !ok {
				return
			}
			fmt.Println("data  ", val, " from worker number ", num+1)
		}
	}
}

func FanOut() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	ch := generator(&wg, ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// spawn four workers to receive this above channel
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(&wg, ctx, ch, i)
	}
	wg.Wait()
}
