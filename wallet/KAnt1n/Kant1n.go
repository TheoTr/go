package wallet_KAnt1n

import "fmt"

type Satoshi int

func (s Satoshi) Validate() bool {
	return s <= 21_000_000_0000_0000
}

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Satoshi
}

func (w *Wallet) Deposit(amount Satoshi) {
	if !amount.Validate() {
		return
	}
	w.balance += amount
}

func (w Wallet) Balance() Satoshi {
	return w.balance
}

func (b Satoshi) String() string {
	return fmt.Sprintf("%d SATOSHI", b)
}

func (w *Wallet) Withdraw(amount Satoshi) {
	w.balance -= amount
}
