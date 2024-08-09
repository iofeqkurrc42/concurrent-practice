package messagepassing

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	msgChannel := make(chan string)
	go badReceiver(msgChannel)
	fmt.Println("sending hello")
	msgChannel <- "hello"
	fmt.Println("sending there")
	msgChannel <- "there"
	fmt.Println("sending stop")
	msgChannel <- "stop"
}

func Run1() {
	msgChannel := make(chan string)
	go badSender(msgChannel)
	fmt.Println("reading message from channel")
	msg := <-msgChannel
	fmt.Println("received", msg)
}

func Run2() {
	msgChannel := make(chan int, 3)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go slowReceiver(msgChannel, &wGroup)
	for i := 1; i <= 6; i++ {
		size := len(msgChannel)
		fmt.Printf("%s sending: %d. buffer size: %d\n", time.Now().Format("15:04:05"), i, size)
		msgChannel <- i
	}
	msgChannel <- -1
	wGroup.Wait()
	fmt.Println("channel size: ", len(msgChannel))
}

func Run6() {
	c := NewChannel[int](3)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go slowReceiver1(c, &wGroup)
	for i := 1; i <= 6; i++ {
		fmt.Printf("%s sending: %d\n", time.Now().Format("15:04:05"), i)
		c.Send(i)
	}
	c.Send(-1)
	wGroup.Wait()
}

func slowReceiver1(message *Channel[int], wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = message.Receive()
		fmt.Println("received:", msg)
	}
	wGroup.Done()
}

func badSender(message chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("sender slept for 5 seconds")
}

func slowReceiver(message chan int, wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = <-message
		fmt.Println("received:", msg)
	}
	wGroup.Done()
}

func badReceiver(message chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("receiver slept for 5 seconds")
}

func receiver(message chan string) {
	msg := ""
	for msg != "stop" {
		msg = <-message
		fmt.Println("receiverd:", msg)
	}
}
