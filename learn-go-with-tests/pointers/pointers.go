package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance

}
