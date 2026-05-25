package forselect

import "fmt"

func Forselect() {
	chars := []string{"a", "b", "c"}
	channel := make(chan string, 3)
	for i := 0; i < 3; i++ {
		channel <- chars[i]
	}
	close(channel)
	//we can travese a channel until it is been closed  
	for msg := range channel {
		fmt.Println("message from channel", msg)
	}
}
