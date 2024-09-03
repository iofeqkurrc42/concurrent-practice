package main

import (
	"cocurrent/concurrent/messagepassing"
	"cocurrent/concurrent/sharedmemory"
	"fmt"
	"time"
)

func main() {
	// wg := concurrent.NewWaitGrp(4)
	// for i := 1; i <= 4; i++ {
	// go doWork(i, wg)
	// }
	// wg.Wait()
	// fmt.Println("All done")
	// barrier := concurrent.NewBarrier(2)
	// go workAndWait("red", 4, barrier)
	// go workAndWait("blue", 4, barrier)

	// time.Sleep(100 * time.Second)
	// shared_memory.Run7()
	messagepassing.Run18()
}

func doWork(id int, wg *sharedmemory.WaitGrp) {
	fmt.Println(id, "done working")
	wg.Done()
}

func workAndWait(name string, timeToWork int, barrier *sharedmemory.Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is ruuning")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}
