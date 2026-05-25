// https://share.google/aimode/AeITQ78hFGRlJXlGG

/*
go does not supports inheritance similar to OOPLanguage like cpp,java but it acheives similar behaviour using composition primary through embedings

using embedings we can include struct inside a struct and it can get access to all the methods automatically , now this make us to resued the code and enable insheritance like behaviour wihtout tight coupking
*/

package inheritance

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Some sound")
}

type Dog struct{
	Animal
}

type Pup struct {
	Dog
}


func (d Dog) Bark(){
	fmt.Println("Woof!")
}

func Inheritance() {
	mypup := Pup{
		Dog{
			Animal{
				Name: "Don",
			},
		},
	}

	mypup.Speak()
}