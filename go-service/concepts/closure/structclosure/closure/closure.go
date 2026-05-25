package structclosure

import "fmt"

func NewCounter() (func() int) {
	count := 0
	return func() int {
		count++
		return count
	}
}
func Closure() {
	inc := NewCounter()

	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
}