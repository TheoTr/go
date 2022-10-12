package wallet_aegla

import (
	"errors"
	"fmt"
)

type Satoshi uint64

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

func (s Satoshi) String() string {
	bitcoins := s / 1_000_000_00
	return fmt.Sprintf("%d BTC", bitcoins)
}

func (w *Wallet) Withdraw(amount Satoshi) error {
	if amount > w.balance {
		return errors.New("NOT ENOUGH MONEY")
	}

	w.balance -= amount
	return nil
}