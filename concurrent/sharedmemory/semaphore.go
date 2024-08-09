package sharedmemory

import "sync"

type Semaphore struct {
	cond     *sync.Cond
	permints int
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		cond:     sync.NewCond(&sync.Mutex{}),
		permints: n,
	}
}

func (rw *Semaphore) Acquire() {
	rw.cond.L.Lock()
	for rw.permints <= 0 {
		rw.cond.Wait()
	}
	rw.permints--
	rw.cond.L.Unlock()
}

func (rw *Semaphore) Release() {
	rw.cond.L.Lock()

	rw.permints++
	rw.cond.Signal()

	rw.cond.L.Unlock()
}
