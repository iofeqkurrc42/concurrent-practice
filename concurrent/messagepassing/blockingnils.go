package messagepassing

import "fmt"

func Run12() {
	var ch chan string = nil

	ch <- "message"
	fmt.Println("This is never printed")
}
