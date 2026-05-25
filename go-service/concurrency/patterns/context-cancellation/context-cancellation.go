package contextcancellation

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func toString(ctx context.Context, in <-chan int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return

			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case out <- strconv.Itoa(v):
				}
			}
		}
	}()

	return out
}

func multiply(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return

			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case out <- v * 2:
				}
			}
		}
	}()

	return out
}

func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	go func() {
		

		for _, n := range nums {
			select {
			case <-ctx.Done():
				defer close(out)
				return
			case out <- n:
			}
		}
	}()

	return out
}

func Contextcancellation() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	input := generator(ctx, 1, 2, 3, 4, 5, 6, 7)

	multiplied := multiply(ctx, input)

	final := toString(ctx, multiplied)

	// so over here we aren't finding need to use waitgroup becuase main is blocked indirectly for recieving the input
	for v := range final {
		fmt.Println("output:", v)
	}

	fmt.Println("pipeline finished or cancelled")
}