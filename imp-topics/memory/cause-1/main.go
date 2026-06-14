package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

/*
This is cause 1 of memory allocation : when we are creating too much
1. make() inside a for loop
2. append() causing repeated allocation
3. string concatenation in using a foor loop
*/

func badMake(n int) [][]int {
	var result [][]int
	for i := 0; i < n; i++ {
		temp := make([]int,0) // allocation happening in every iteration
		temp = append(temp, i)
		result = append(result, temp)
	}
	return result
}

func badAppend(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i) // reallocation + copy every time
	}
	return arr
}

func badString(n int) string{
	s := ""
	for i := 0 ; i < n ; i++{
		s += "a" // new string every time
	}
	return s 
}

func main() {
	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	for {
		_ = badAppend(10_000_000)
	}
}
