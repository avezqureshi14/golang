package main

import "fmt"

func printEven(evenCh, oddCh chan bool) {
	for i := 0; i <= 8; i += 2 {
		<-evenCh
		fmt.Println(i)
		oddCh <- true
	}
}

func printOdd(evenCh, oddCh chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- true
	}
}

func main() {
	evenCh := make(chan bool)
	oddCh := make(chan bool)

	go printEven(evenCh, oddCh)
	go printOdd(evenCh, oddCh)

	evenCh <- true // start execution

	select {}
}