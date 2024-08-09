package messagepassing

import (
	"fmt"
	"time"
)

func directReceiver(message <-chan int) {
	for {
		msg := <-message
		fmt.Println(time.Now().Format("15:04:05"), "Received", msg)
	}
}

func directSender(message chan<- int) {
	for i := 0; ; i++ {
		message <- i
		fmt.Println(time.Now().Format("15:04:05"), "Sending", i)
	}
}

func Run3() {
	msgChannel := make(chan int)
	go directReceiver(msgChannel)
	go directSender(msgChannel)
	time.Sleep(5 * time.Second)
}
