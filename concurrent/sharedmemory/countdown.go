package sharedmemory

import (
	"fmt"
	"time"
)

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1
	}
}

func Run() {
	count := 5
	go countdown(&count)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}
}
