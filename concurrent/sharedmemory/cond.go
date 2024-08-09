package sharedmemory

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func Run4() {
	money := 100
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go stingy1(&money, cond)
	go spendy1(&money, cond)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("moeny in bank account:", money)
	mutex.Unlock()
}

func stingy1(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		cond.Signal()
		cond.L.Unlock()
	}
	fmt.Println("stingly done")
}

func spendy1(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		cond.L.Lock()
		for *money < 50 {
			cond.Wait()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("money is negative")
			os.Exit(1)
		}
		cond.L.Unlock()
	}
	fmt.Println("spendy done")
}
