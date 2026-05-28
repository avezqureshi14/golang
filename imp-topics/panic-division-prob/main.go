package main

import "fmt"

func SafeDivide(a, b int) int {
	fmt.Println("Defer inside safe divide always run")
	if b == 0 {
		panic("division by zero detected")
	}
	return a / b
}

func ExcecuteDivison(a, b int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic ", r)
		}
	}()

	fmt.Println("Result :", SafeDivide(a, b))
}

func NestedDefer() {
	fmt.Println("Nested defer demo")
	defer fmt.Println("defer 1")
	defer func() {
		fmt.Println("defer 2 start")
		defer fmt.Println("defer 2.1")
		defer fmt.Println("defer 2.2")
		fmt.Println("defer 2 end")
	}()
	defer fmt.Println("defer 4")
	fmt.Println("function body excecuting")
}

func main() {
	ExcecuteDivison(1, 0)
	NestedDefer()
}
