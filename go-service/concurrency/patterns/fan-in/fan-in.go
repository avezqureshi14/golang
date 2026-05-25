package fanin

import (
	"fmt"
	"time"
)

func producer1() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(300 * time.Millisecond)
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func producer2() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 100; i <= 105; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func fanIn(ch1 <-chan int, ch2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		close(out)

		for {
			select {
			case v, ok := <-ch1:
				if ok {
					out <- v
				} else {
					ch1 = nil
				}

			case v, ok := <-ch2:
				if ok {
					out <- v
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				return
			}
		}
	}()

	return out
}

func FanIn() {
	ch1 := producer1()
	ch2 := producer2()

	merged := fanIn(ch1, ch2)

	for v := range merged {
		fmt.Println(v)
	}
}
