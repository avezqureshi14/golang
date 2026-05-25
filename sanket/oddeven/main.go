package main

import (
	"fmt"
	"sync"
)

func printEven(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		<-even
		fmt.Println(i)
		if i != 100 {
			odd <- true
		}
	}
}

func printOdd(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 99; i += 2 {
		<-odd
		fmt.Println(i)
		even <- true
	}
}

func main() {
	odd := make(chan bool)
	even := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go printEven(even, odd, &wg)
	go printOdd(even, odd, &wg)

	odd <- true

	wg.Wait()
}
