package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	checkBlance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("wallet balance got %s but want %s", got, want)
		}
	}
	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(10)
		checkBlance(t, wallet, Bitcoin(20))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		checkBlance(t, wallet, Bitcoin(10))
	})

}
