package sharedmemory

import "sync"

type FiexableWaitGrp struct {
	cond      *sync.Cond
	groupSize int
}

func NewFiexableWaitGrp() *FiexableWaitGrp {
	return &FiexableWaitGrp{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *FiexableWaitGrp) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *FiexableWaitGrp) Wait() {
	wg.cond.L.Lock()
	if wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *FiexableWaitGrp) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}
