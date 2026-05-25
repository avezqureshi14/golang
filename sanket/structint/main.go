package main

import "fmt"

type CarRepoInterface interface {
	Create(number int) string
}

type Car struct {
	name string
}

var _ CarRepoInterface = (*Car)(nil)

func (c *Car) Create(number int) string {
	return fmt.Sprintf("Car %s created with number %d ", c.name, number)
}

func main() {
	var repo CarRepoInterface = &Car{name: "bmw"}
	result := repo.Create(101)
	fmt.Println(result)
}
