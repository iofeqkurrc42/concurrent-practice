package deadlock

import "sync"

type Arbitrator struct {
	// store account with their availabilty status, either freee or in use
	accountIsUse map[string]bool
	// condition varibale to be user to suspend goroutine if account are not available
	cond *sync.Cond
}

func NewArbitrator() *Arbitrator {
	return &Arbitrator{
		accountIsUse: make(map[string]bool),
		cond:         sync.NewCond(&sync.Mutex{}),
	}
}

func (a *Arbitrator) LockAccounts(ids ...string) {
	a.cond.L.Lock()
	for allAvailable := false; !allAvailable; {
		allAvailable = true
		for _, id := range ids {
			if !a.accountIsUse[id] {
				allAvailable = false
				a.cond.Wait()
			}
		}
	}
	for _, id := range ids {
		a.accountIsUse[id] = true
	}
	a.cond.L.Unlock()
}

func (a *Arbitrator) UnlockAccounts(ids ...string) {
	a.cond.L.Lock()
	for _, id := range ids {
		a.accountIsUse[id] = false
	}
	a.cond.Broadcast()
	a.cond.L.Unlock()
}
