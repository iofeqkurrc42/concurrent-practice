package messagepassing

import (
	"fmt"
	"time"
)

func writeEvery(msg string, seconds time.Duration) <-chan string {
	msChannl := make(chan string)
	go func() {
		for {
			time.Sleep(seconds)
			msChannl <- msg
		}
	}()
	return msChannl
}

func Run7() {
	mFromA := writeEvery("tick", 1*time.Second)
	mFromB := writeEvery("tock", 3*time.Second)

	for {
		select {
		case msg1 := <-mFromA:
			fmt.Println(msg1)
		case msg2 := <-mFromB:
			fmt.Println(msg2)
		}
	}
}
