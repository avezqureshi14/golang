package factory

import "fmt"

// over here we are able to create factory functions

func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func Factory() {
	double := multiplier(2)
	triple := multiplier(3)
	dbVal := double(10)
	trVal := triple(10)
	fmt.Println("Double is ",dbVal)
	fmt.Println("Triple is ",trVal)
}