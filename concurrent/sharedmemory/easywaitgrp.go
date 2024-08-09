package sharedmemory

type WaitGrp struct {
	sema *Semaphore
}

func NewWaitGrp(size int) *WaitGrp {
	return &WaitGrp{sema: NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
	wg.sema.Release()
}
