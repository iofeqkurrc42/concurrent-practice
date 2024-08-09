package messagepassing

import (
	"fmt"
	"math/rand"
	"time"
)

func generateTemp() chan int {
	output := make(chan int)
	go func() {
		temp := 50
		for {
			output <- temp
			temp += rand.Intn(3) + 1
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func outputTemp(input chan int) {
	go func() {
		for {
			fmt.Println("current tmep", <-input)
			time.Sleep(2 * time.Second)
		}
	}()
}

func Run15() {
	var last int

	input := make(chan int)
	outputTemp(input)
	output := generateTemp()

	for {
		select {
		case tmp := <-output:
			last = tmp
		case input <- last:
		}
	}
}
