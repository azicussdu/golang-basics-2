package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance float64
	mu      sync.Mutex
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	ba.balance += amount
	fmt.Printf("deposited: %.2f, balance: %.2f\n", amount, ba.balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	if ba.balance >= amount {
		ba.balance -= amount
		fmt.Printf("withdrawed: %.2f, balance: %.2f\n", amount, ba.balance)
	} else {
		fmt.Println("Couldn't withdraw")
	}
}

func (ba *BankAccount) MyBalance() {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	fmt.Printf("My balance: %.2f\n", ba.balance)
}

//func main() {
//	bac := BankAccount{balance: 1000}
//
//	var wg sync.WaitGroup
//	wg.Add(4)
//
//	go func() {
//		defer wg.Done()
//		bac.Deposit(500)
//	}()
//
//	go func() {
//		defer wg.Done()
//		bac.Withdraw(2000)
//	}()
//
//	go func() {
//		defer wg.Done()
//		bac.Deposit(1000)
//	}()
//
//	go func() {
//		defer wg.Done()
//		bac.Withdraw(2000)
//	}()
//
//	wg.Wait()
//
//	bac.MyBalance()
//}
