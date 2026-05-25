package donechannel

import (
	"fmt"
	"time"
)

func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

func Donechannel() {
	done := make(chan bool)
	go dowork(done)

	time.Sleep(time.Second * 3)
	close(done)
}
