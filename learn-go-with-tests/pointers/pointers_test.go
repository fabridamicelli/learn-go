package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, exp Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("Expected no error, got %q", got)

		}

	}

	assertError := func(t testing.TB, got error, exp error) {
		t.Helper()
		if got == nil {
			t.Error("Expected error, got none")
		}
		if got.Error() != exp.Error() {
			t.Errorf("Expected %q, got %q", exp, got)

		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(Bitcoin(23))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)

	})
}
