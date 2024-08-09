package messagepassing

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Run10() {
	t, _ := strconv.Atoi(os.Args[1])
	message := sendMsgAfter(3 * time.Second)
	timeoutDuration := time.Duration(t) * time.Second
	fmt.Printf("Waiting for message %v seconds\n", t)
	select {
	case msg := <-message:
		fmt.Println("message received:", msg)
	case tNow := <-time.After(timeoutDuration):
		fmt.Println("time out. waited until", tNow.Format("15:04:05"))
	}
}
