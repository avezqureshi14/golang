package structclosure

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() int {
	c.count++
	return c.count
}


func Struct() {
	c := &Counter{}

	fmt.Println(c.Increment()) // 1
	fmt.Println(c.Increment()) // 2
	
}