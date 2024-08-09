package messagepassing

import (
	"fmt"
	"time"
)

func onlyReceiver(message <-chan int) {
	for {
		msg := <-message
		fmt.Println(time.Now().Format("15:04:05"), "received:", msg)
		time.Sleep(1 * time.Second)
	}
}

func Run4() {
	msgChannel := make(chan int)
	go onlyReceiver(msgChannel)
	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "sending:", i)
		msgChannel <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChannel)
	time.Sleep(3 * time.Second)
}

func closseFlag(message <-chan int) {
	for {
		msg, more := <-message
		fmt.Println(time.Now().Format("15:04:05"), "received:", msg, more)
		time.Sleep(1 * time.Second)
		if !more {
			return
		}
	}
}
