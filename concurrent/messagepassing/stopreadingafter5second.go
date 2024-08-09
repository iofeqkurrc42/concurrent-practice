package messagepassing

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers() chan int {
	output := make(chan int)
	go func() {
		for {
			output <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func Run16() {
	timeoutDuration := time.Duration(5) * time.Second
	timeOut := time.After(timeoutDuration)
	fmt.Printf("Waiting for message %v seconds\n", 5)

	numbers := generateNumbers()
	for {
		select {
		case generatedNumber := <-numbers:
			fmt.Println("number received:", generatedNumber)
		case tNow := <-timeOut:
			fmt.Println("time out. waited until", tNow.Format("15:04:05"))
			return
		}
	}
}
