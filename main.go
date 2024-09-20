package main

import (
	"cocurrent/concurrent/deadlock"
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
	// messagepassing.Run24(100)
	//for i := 0; i < 5; i++ {
	//  go func() {
	//	  fmt.Println(i)
	//	}()
	//}
	//time.Sleep(1 * time.Second)
	// patterns.Run6()
	deadlock.Run()
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
