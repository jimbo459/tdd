package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Error("Expected error to Have Ocured")
		}
		if got.Error() != want.Error() {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Error("Expected Error not to have occured")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(20.0))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5.0))
	})

	t.Run("Insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: 10}

		err := wallet.Withdraw(Bitcoin(15.0))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, errInsuficientFunds)
	})
}
