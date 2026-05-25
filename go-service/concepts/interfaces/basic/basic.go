package basic

type speaker interface {
	Speak() string
}

type Dog struct{}

// Dog implements Speaker just by having this method
func (d Dog) Speak() string {
	return "Woof!"
}

func Basic() {

}