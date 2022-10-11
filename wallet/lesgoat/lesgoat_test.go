package wallet_lesgoat

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Satoshi(10))

		got := wallet.Balance()
		want := Satoshi(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Satoshi(20)}

		wallet.Withdraw(Satoshi(10))

		got := wallet.Balance()

		want := Satoshi(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
