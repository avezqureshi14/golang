package closure

import "fmt"

// ok so first benefit of this closure is encapsulation data hiding , we don't have to create a global variable var count and shared them

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Closure() {
	count := counter()
	fmt.Println("Count  ", count())
	fmt.Println("Count  ", count())
	fmt.Println("Count  ", count())
}
