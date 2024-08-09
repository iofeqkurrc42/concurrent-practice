package messagepassing

import (
	"fmt"
	"math/rand"
	"time"
)

func generateAmounts(n int) <-chan int {
	amounts := make(chan int)
	go func() {
		defer close(amounts)
		for i := 0; i < n; i++ {
			amounts <- rand.Intn(100) + 1
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return amounts
}

func Run13() {
	sales := generateAmounts(50)
	expenses := generateAmounts(40)

	endOfDayAmount := 0
	for sales != nil || expenses != nil {
		select {
		case sale, ok := <-sales:
			if ok {
				fmt.Println("sale of:", sale)
				endOfDayAmount += sale
			} else {
				sales = nil
			}
		case expense, ok := <-expenses:
			if ok {
				fmt.Println("expense of:", expense)
				endOfDayAmount -= expense
			} else {
				expenses = nil
			}
		}
	}
	fmt.Println("End of day profit and loss:", endOfDayAmount)
}
