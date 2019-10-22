package main

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposite(amount int) {
	fmt.Println("addr ", &w.balance)
	w.balance = amount
}
func (w Wallet) Balance() int {
	return w.balance
}
