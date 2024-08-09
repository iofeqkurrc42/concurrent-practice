package messagepassing

import (
	"fmt"
	"time"
)

func sendMsgAfter(seconds time.Duration) <-chan string {
	message := make(chan string)
	go func() {
		time.Sleep(seconds)
		message <- "Hello"
	}()
	return message
}

func Run8() {
	message := sendMsgAfter(3 * time.Second)
	for {
		select {
		case msg := <-message:
			fmt.Println("message received", msg)
			return
		default:
			fmt.Println("No message waiting")
			time.Sleep(1 * time.Second)
		}
	}
}
