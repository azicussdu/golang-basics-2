package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type BankAccout struct {
	Name    string
	Balance float64
}

func (b *BankAccout) ProcessPayment(amount float64) error {
	if b.Balance <= 0 {
		return fmt.Errorf("No balance %w\n", ErrInsufficientFunds)
	} else if amount > b.Balance {
		return fmt.Errorf("Balance < amount %w\n", ErrInsufficientFunds)
	}

	b.Balance -= amount
	return nil
}

//func main() {
//	ba := BankAccout{Name: "Alibek Kaidasyn", Balance: 500}
//	err := ba.ProcessPayment(1000)
//	if err != nil {
//		fmt.Println(err.Error())
//		innerErr := errors.Unwrap(err)
//		if innerErr != nil {
//			fmt.Println(innerErr.Error())
//		}
//	}
//}
