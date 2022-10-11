package wallet_lesgoat

import "fmt"

type Satoshi int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Satoshi
}

func (w *Wallet) Deposit(amount Satoshi) {
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
