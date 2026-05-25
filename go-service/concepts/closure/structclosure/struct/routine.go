package structclosure

import "fmt"

type Worker struct {
	val int
}

func (w Worker) Run() {
	fmt.Println("Processing...", w.val)
}

func Routine() {
	for i := 0; i < 3; i++ {
		w := Worker{val:i}
		go w.Run()
	}
}
