package generator

import (
	"fmt"
	"math/rand"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():

			}
		}
	}()
	return stream
}

func Generator() {
	done := make(chan int)
	defer close(done)

	randomNumFetcher := func() int { return rand.Intn(5000000) }
	for rando := range repeatFunc(done,randomNumFetcher){
		fmt.Println(rando)
	}
}