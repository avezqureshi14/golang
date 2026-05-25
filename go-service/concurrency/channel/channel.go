package channel

import "fmt"

func Channel() {
	mychannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		mychannel <- "data"
	}()
	select {
	case msgFromMyChannel := <-mychannel:
		fmt.Println("message from my channel ",msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println("msg from another channel ",msgFromAnotherChannel)
	}
 
}