package main

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type Counter2 struct {
	count int
}

func (c *Counter2) increment() int {
	c.count++
	return c.count
}

func main() {

}
