package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func NewBankAccount(owner string, balance float64) BankAccount {
	if balance < 0 {
		balance = 0
	}
	bankAcc := BankAccount{Owner: owner, Balance: balance}
	return bankAcc
}

func (b BankAccount) Inform() {
	fmt.Printf("Name: %s, balance: %.2f\n", b.Owner, b.Balance)
}

func (b *BankAccount) Deposit(amount float64) {
	if amount < 0 {
		amount = 0
	}

	b.Balance = b.Balance + amount
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount < 0 || amount > b.Balance {
		return false
	}

	b.Balance = b.Balance - amount
	return true
}

//func main() {
//	ba := NewBankAccount("Beka", 10000)
//	ba.Inform()
//
//	ba.Deposit(5000)
//	ba.Inform()
//
//	sheshtim := ba.Withdraw(7000)
//	if sheshtim == false {
//		fmt.Println("Sheshe almadyk")
//	}
//	ba.Inform()
//}
