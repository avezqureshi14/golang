package pipeline

import (
	"fmt"
	"strconv"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func multiply(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func toString(in <-chan int) <-chan string {
	out := make(chan string)

	go func() {
		for v := range in {
			out <- strconv.Itoa(v)
		}
		close(out)
	}()

	return out
}

func PipeLine() {
	// stage 1
	input := generator(1, 2, 3, 4, 5)

	// stage 2
	multiplied := multiply(input)

	// stage 3
	final := toString(multiplied)

	// consume output
	for v := range final {
		fmt.Println(v)
	}
}