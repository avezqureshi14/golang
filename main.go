package main

import (
	"fmt"
	"time"
)

var counter = 0

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // ❌ race condition happens here
	}
}

func main() {
	go increment()
	go increment()

	time.Sleep(time.Second)

	fmt.Println("Final Counter:", counter)
}
