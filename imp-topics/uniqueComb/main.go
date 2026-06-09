package main

import "fmt"

//////////////////////
// 1. ENCAPSULATION  : Encapsulation is the bundling of data and methods that operate on that data into a single unit, while restricting direct access to object's components
//////////////////////

type Person struct {
	name string // unexported (private)
	age  int
}

// Getter
func (p *Person) GetName() string {
	return p.name
}

// Setter
func (p *Person) SetName(name string) {
	p.name = name
}

//////////////////////
// 2. ABSTRACTION   : Exposing all the implementation and exposing only the essential features of an object
//////////////////////

type Animal interface {
	Speak() string
	Move() string
}

//////////////////////
// 3. POLYMORPHISM  : It is the ability of different objects to respond to the same call in different ways
//////////////////////

type Dog struct {
	Person // embedding (composition)
	Breed  string
}

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) Move() string {
	return "Runs"
}

type Cat struct {
	Person
	Color string
}

func (c Cat) Speak() string {
	return "Meow"
}

func (c Cat) Move() string {
	return "Walks silently"
}

//////////////////////
// 4. COMPOSITION   : composition is a desing principle where a type is built by combining other types intead of inheriting from them
//////////////////////

type Engine struct {
	HorsePower int
}

func (e Engine) Start() {
	fmt.Println("Engine started with HP:", e.HorsePower)
}

type Car struct {
	Engine // embedding instead of inheritance
	Brand  string
}

//////////////////////
// MAIN FUNCTION    //
//////////////////////

func main() {

	// Encapsulation
	p := Person{name: "Avez", age: 22}
	fmt.Println("Name:", p.GetName())

	p.SetName("Khan")
	fmt.Println("Updated Name:", p.GetName())

	// Polymorphism using interface
	var a Animal

	dog := Dog{
		Person: Person{name: "Tommy", age: 5},
		Breed:  "Labrador",
	}

	cat := Cat{
		Person: Person{name: "Kitty", age: 3},
		Color:  "White",
	}

	a = dog
	fmt.Println("Dog:", a.Speak(), "-", a.Move())

	a = cat
	fmt.Println("Cat:", a.Speak(), "-", a.Move())

	// Composition
	car := Car{
		Engine: Engine{HorsePower: 150},
		Brand:  "Toyota",
	}

	car.Start() // inherited behavior via embedding
	fmt.Println("Car brand:", car.Brand)
}
