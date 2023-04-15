package main

import (
	"fmt"
	"sync"
)

func (a *Account) depositWithoutRaceConditionRWMutex(
	amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	// Decrement the wait group
	defer wg.Done()

	lock.Lock()
	b := a.balance
	a.balance = b + amount
	lock.Unlock()
}

func (a *Account) getBalanceWithoutRaceConditions(
	lock *sync.RWMutex) int {
	lock.RLock()
	b := a.balance
	lock.RUnlock()
	return b
}

func RunRaceConditionFixWithRWExample() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	// Code with a race condition
	// go build --race main.go
	a := Account{
		balance: 100,
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go a.depositWithoutRaceConditionRWMutex(i*100, &wg, &lock)
	}

	wg.Wait()
	fmt.Println(a.getBalanceWithoutRaceConditions(&lock))
}
