package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(10)
		checkBlance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		checkBlance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(30)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(40)
		checkBlance(t, wallet, startBalance)
		assertError(t, err, InsufficientFoundsError)
	})
}

func checkBlance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("wallet balance got %s but want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
}
