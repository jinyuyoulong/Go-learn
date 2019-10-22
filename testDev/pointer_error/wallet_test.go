package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposite(10)
	fmt.Println("addr ", &wallet.balance)
	got := wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("wallet balance got %d but want %d", got, want)
	}
}
