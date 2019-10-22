package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposite(10)
	got := wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("wallet balance got %d but want %d", got, want)
	}
}
