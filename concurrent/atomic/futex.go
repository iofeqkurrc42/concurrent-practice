package atomic

import "sync/atomic"

// Futex is short for fast userspace mutex.
// A futex is a wait queue primitive that we can access from user space.
// assume  we  have  two  system  calls  named  futex_wait(address,value) and futex_wake(address,count)
type FutexLock int32

// The futex_wake(addr, count) wakes up suspended executions (threads and processes)
// that are waiting on the address specified.
func futex_wakeup(address *int32, count int) {
	// system call
}

// If the value at the memory address is equal to the specified parameter value,
// the execution of the caller is suspended and placed at the back of a queue
func futex_wait(address *int32, value int32) {
	// system call
}

func (f *FutexLock) Lock() {
	if !atomic.CompareAndSwapInt32((*int32)(f), 0, 1) {
		for atomic.SwapInt32((*int32)(f), 2) != 0 {
			futex_wait((*int32)(f), 2)
		}
	}
}

func (f *FutexLock) Unlock() {
	oldValue := atomic.SwapInt32((*int32)(f), 0)
	if oldValue == 2 {
		futex_wakeup((*int32)(f), 1)
	}
}
