package account

import (
	"sync"
)

type Account struct {
	balance int64
	opened bool
	mutex sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	var acc Account
	acc.balance = amount
	acc.opened = true
	acc.mutex = sync.Mutex{}
	return &acc
}

func (a *Account) Balance() (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.opened {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.opened {
		return 0, false
	}
	if a.balance + amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.opened {
		return 0, false
	}
	sum := a.balance
	a.balance = 0
	a.opened = false
	return sum, true
}
