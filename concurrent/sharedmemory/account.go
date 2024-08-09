package sharedmemory

import (
	"fmt"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		*money++
		mutex.Unlock()
		// runtime.Gosched()
	}
	fmt.Println("Stingy done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		*money--
		mutex.Unlock()
		// runtime.Gosched()
	}
	fmt.Println("Spendy done")
}

func Run2() {
	money := 100
	mutex := sync.Mutex{}
	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	println("money in bank account:", money)
	mutex.Unlock()
}
