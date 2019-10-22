package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(10)
		got := wallet.Balance()
		want := Bitcoin(20)
		if got != want {
			t.Errorf("wallet balance got %s but want %s", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("wallet balance got %s but want %s", got, want)
		}
	})

}
