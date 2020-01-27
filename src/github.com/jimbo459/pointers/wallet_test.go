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

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10.0))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5.0))
	})

	t.Run("Insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: 10}

		err := wallet.Withdraw(Bitcoin(15))

		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Errorf("Expected Error for insufficent funds")
		}

	})
}
