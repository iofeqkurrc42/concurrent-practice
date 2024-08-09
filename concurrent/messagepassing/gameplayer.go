package messagepassing

import (
	"fmt"
	"math/rand"
	"time"
)

func player() chan string {
	output := make(chan string)
	count := rand.Intn(100)
	move := []string{"up", "down", "left", "right"}

	go func() {
		defer close(output)
		for i := 0; i < count; i++ {
			output <- move[rand.Intn(4)]
			d := time.Duration(rand.Intn(200))
			time.Sleep(d * time.Millisecond)
		}
	}()

	return output
}

func Run17() {
	player1 := player()
	player2 := player()
	player3 := player()
	player4 := player()
	count := 4

	for {
		select {
		case p1, ok := <-player1:
			if ok {
				fmt.Println("player 1", p1)
			} else {
				player1 = nil
				count -= 1
				fmt.Println("player 1 left the game. Remaining", count)
			}
		case p2, ok := <-player2:
			if ok {
				fmt.Println("player 2", p2)
			} else {
				player2 = nil
				count -= 1
				fmt.Println("player 2 left the game. Remaining", count)
			}
		case p3, ok := <-player3:
			if ok {
				fmt.Println("player 3", p3)
			} else {
				player3 = nil
				count -= 1
				fmt.Println("player 3 left the game. Remaining", count)
			}
		case p4, ok := <-player4:
			if ok {
				fmt.Println("player 4", p4)
			} else {
				player4 = nil
				count -= 1
				fmt.Println("player 4 left the game. Remaining", count)
			}
		}
		if count == 0 {
			fmt.Println("game finished")
			return
		}
	}
}
