package main

import (
	"fmt"
	"sync"
)

func (a *Account) depositWithoutRaceCondition(
	amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	// Decrement the wait group
	defer wg.Done()

	lock.Lock()
	b := a.balance
	a.balance = b + amount
	lock.Unlock()
}

func RunRaceConditionFixExample() {
	var wg sync.WaitGroup
	// A lock in the program to restrict the code lines between Lock and Unlock methods
	var lock sync.Mutex

	// Code with a race condition
	// go build --race main.go
	a := Account{
		balance: 100,
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go a.depositWithoutRaceCondition(i*100, &wg, &lock)
	}

	wg.Wait()
	fmt.Println(a.getBalance())
}
