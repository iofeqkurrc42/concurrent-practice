package deadlock

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type BankAccount struct {
	id      string
	balance int
	mutex   sync.Mutex
}

func NewBankAccount(id string) *BankAccount {
	return &BankAccount{
		id:      id,
		balance: 100,
		mutex:   sync.Mutex{},
	}
}

func (src *BankAccount) Transfer(to *BankAccount, amount, exId int) {
	fmt.Printf("%d Locking %s's account\n", exId, src.id)
	src.mutex.Lock()
	fmt.Printf("%d Locking %s's account\n", exId, to.id)
	to.mutex.Lock()
	src.balance -= amount
	to.balance += amount
	to.mutex.Unlock()
	src.mutex.Unlock()
	fmt.Printf("%d Unlocked %s and %s\n", exId, src.id, to.id)
}

func (src *BankAccount) Transfer2(to *BankAccount, amount int, tellerId int, arb *Arbitrator) {
	fmt.Printf("$d Locking %s and %s\n", tellerId, src.id, to.id)
	arb.LockAccounts(src.id, to.id)
	src.balance -= amount
	to.balance += amount
	arb.UnlockAccounts(src.id, to.id)
	fmt.Printf("%d Unlocked %s and %s\n", tellerId, src.id, to.id)
}

func (src *BankAccount) Transfer3(to *BankAccount, amount int, tellerId int) {
	accounts := []*BankAccount{src, to}
	sort.Slice(accounts, func(a, b int) bool {
		return accounts[a].id < accounts[b].id
	})
	fmt.Printf("%d Locking %s's account\n", tellerId, accounts[0].id)
	accounts[0].mutex.Lock()
	fmt.Printf("%d Locking %s's account\n", tellerId, accounts[1].id)
	accounts[1].mutex.Lock()
	src.balance -= amount
	to.balance += amount
	to.mutex.Unlock()
	src.mutex.Unlock()
	fmt.Printf("%d Unlocked %s and %s\n", tellerId, src.id, to.id)
}

func Run1() {
	accounts := []BankAccount{
		*NewBankAccount("Sam"),
		*NewBankAccount("Paul"),
		*NewBankAccount("Akmy"),
		*NewBankAccount("Mia"),
	}
	total := len(accounts)
	for i := 0; i < 4; i++ {
		go func(eId int) {
			for j := 1; j < 1000; j++ {
				from, to := rand.Intn(total), rand.Intn(total)
				for from == to {
					to = rand.Intn(total)
				}
				accounts[from].Transfer(&accounts[to], 10, eId)
			}
			fmt.Println(eId, "COMPLETE")
		}(i)
		time.Sleep(60 * time.Second)
	}
}

func Run4() {
	accounts := []BankAccount{
		*NewBankAccount("Sam"),
		*NewBankAccount("Paul"),
		*NewBankAccount("Akmy"),
		*NewBankAccount("Mia"),
	}
	total := len(accounts)
	arb := NewArbitrator()
	for i := 0; i < 4; i++ {
		go func(tellerId int) {
			for i := 1; i < 1000; i++ {
				from, to := rand.Intn(total), rand.Intn(total)
				for from == to {
					to = rand.Intn(total)
				}
				accounts[from].Transfer2(&accounts[to], 10, tellerId, arb)
			}
			fmt.Println(tellerId, "COMPLETE")
		}(i)
	}
	time.Sleep(60 * time.Second)
}

func Run5() {
	accounts := []BankAccount{
		*NewBankAccount("Sam"),
		*NewBankAccount("Paul"),
		*NewBankAccount("Akmy"),
		*NewBankAccount("Mia"),
	}
	total := len(accounts)
	for i := 0; i < 4; i++ {
		go func(tellerId int) {
			for i := 1; i < 10; i++ {
				from, to := rand.Intn(total), rand.Intn(total)
				for from == to {
					to = rand.Intn(total)
				}
				accounts[from].Transfer3(&accounts[to], 10, tellerId)
			}
			fmt.Println(tellerId, "COMPLETE")
		}(i)
	}
	time.Sleep(60 * time.Second)
}
