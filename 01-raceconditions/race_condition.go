package main

import (
	"fmt"
	"sync"
)

// A shared resource
type Account struct {
	balance int
}

// Async func who modify a Shared resource
func (a *Account) depositRaceCondition(amount int, wg *sync.WaitGroup) {
	// Decrement the wait group
	defer wg.Done()

	b := a.balance
	a.balance = b + amount
}

func (a *Account) getBalance() int {
	b := a.balance
	return b
}

func RunRaceConditionExample() {
	var wg sync.WaitGroup

	// Code with a race condition
	// go build --race main.go
	a := Account{
		balance: 100,
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go a.depositRaceCondition(i*100, &wg)
	}

	wg.Wait()
	fmt.Println(a.getBalance())
}
