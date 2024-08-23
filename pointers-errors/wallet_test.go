package pointers_errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if !errors.Is(got, want) {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(15))

		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(65)}

		err := wallet.Withdraw(Bitcoin(30))
		if err != nil {
			t.Fatal("received an error but didn't expect one")
		}

		assertBalance(t, wallet, Bitcoin(35))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(50)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(80))

		assertError(t, err, ErrorInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
