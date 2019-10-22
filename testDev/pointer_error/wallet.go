package main

type Wallet struct {
	balance int
}

func (w Wallet) Deposite(amount int) {
	w.balance = amount
}
func (w Wallet) Balance() int {
	return w.balance
}
