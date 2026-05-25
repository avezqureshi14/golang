package structclosure

import "fmt"

func Routine() {
	for i := 0; i < 3; i++ {
		go func(x int) {
			fmt.Println("Processing: ",x)
		}(i)
	}
}
