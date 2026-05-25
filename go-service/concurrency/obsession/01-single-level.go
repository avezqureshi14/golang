/*
1. Single producer → single consumer
Problem:
Generate numbers 1–100 in a goroutine and consume them.
Constraints:
* producer must close channel
* consumer must stop safely
* no WaitGroup allowed
*/

package obsession

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 100; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}
func Level1() {
	ctx, cancel := context.WithTimeout(context.Background(),time.Millisecond*5)
	defer cancel()

	ch := producer(ctx)
	for i := range ch {
		fmt.Println(i)
	}
}

/*
1. Who starts go routines
over here out go routine is started by producer
2. stop condition
in our above code stop condition for our go routine is finite loop completion
3. who sends on channel : producer sends on the channel
4. who closes the cahnnel : producer closes the channel
5. who consumes from the channel : main routine consumes from the channel
6. if consumer stops early i m using context+select to stop producer safely
7. if producer crashes we can use error group to propogate failuer and stop everything, but don't know how to use it over here 
8. if timeout happens we are using auto shutdown so no one runs infinitely
*/
