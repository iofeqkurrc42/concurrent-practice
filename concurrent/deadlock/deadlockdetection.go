package deadlock

import (
	"fmt"
	"sync"
	"time"
)

func lockBoth(lock1, lock2 *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		lock1.Lock()
		lock2.Lock()
		lock1.Unlock()
		lock2.Unlock()
	}
	wg.Done()
}

func Run2() {
	lockA, lockB := sync.Mutex{}, sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go lockBoth(&lockA, &lockB, &wg)
	go lockBoth(&lockB, &lockA, &wg)
	wg.Wait()
	fmt.Println("Done")
}

func Run3() {
	lockA, lockB := sync.Mutex{}, sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go lockBoth(&lockA, &lockB, &wg)
	go lockBoth(&lockB, &lockA, &wg)
	go func() {
		wg.Wait()
		fmt.Println("Done waiting on waitgroup")
	}()
	time.Sleep(30 * time.Second)
	fmt.Println("Done")
}
