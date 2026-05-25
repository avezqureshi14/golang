package routines

import (
	"fmt"
	"time"
)

func workerRoutine(num string) {
	fmt.Println(num)
}

func GoRoutines() {
	go workerRoutine("1")
	go workerRoutine("2")
	go workerRoutine("3")

	time.Sleep(time.Second * 2)

	fmt.Println("hi")
}