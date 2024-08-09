package sharedmemory

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents, "Match event"+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Append match event")
	}
}

func clientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.RLock()
		allEvents := copyAllEvents(mEvents)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "evnets copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	allEvents = append(allEvents, *matchEvents...)
	return allEvents
}

func Run3() {
	mutex := sync.RWMutex{}
	matchEvents := make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "match event")
	}
	go matchRecorder(&matchEvents, &mutex)
	start := time.Now()
	for j := 0; j < 5000; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)
}
