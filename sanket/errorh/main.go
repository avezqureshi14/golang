package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Shutting systems ", r)
		}
	}()
	// program dosen't crashes because panic is handled by recover it gracefully shutsdown
	panic("Something is wrong !!!")
}
