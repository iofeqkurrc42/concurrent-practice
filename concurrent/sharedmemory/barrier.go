package sharedmemory

import "sync"

type Barrier struct {
	cond      *sync.Cond
	size      int
	waitCount int
}

func NewBarrier(size int) *Barrier {
	return &Barrier{
		cond:      sync.NewCond(&sync.Mutex{}),
		size:      size,
		waitCount: 0,
	}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount++
	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.cond.L.Unlock()
}
