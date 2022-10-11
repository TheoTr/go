package wallet_lesgoat

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Satoshi) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Satoshi(10))
		assertBalance(t, wallet, Satoshi(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Satoshi(20)}
		wallet.Withdraw(Satoshi(10))
		assertBalance(t, wallet, Satoshi(10))
	})
	t.Run("Outoflimit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Satoshi(21_000_000_0000_0001))
		assertBalance(t, wallet, Satoshi(0))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Satoshi(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Satoshi(100))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
