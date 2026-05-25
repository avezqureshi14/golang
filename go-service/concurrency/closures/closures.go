package closures

import (
	"fmt"
	"time"
)

func outerFn() {
	for i := 0; i < 3; i++ {
		go func() {
					time.Sleep(10 * time.Millisecond)

			fmt.Println(i)
		}()
	}
}

func Closures() {
	outerFn()
	time.Sleep(time.Second*2)
}
