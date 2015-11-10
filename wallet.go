package wallet

import "errors"

type Cash int
type Wallet struct {
	balance Cash
}

func NewWallet(balance Cash) *Wallet {
	return &Wallet{balance}
}

func (w *Wallet) Balance() Cash {
	return w.balance
}

func (w *Wallet) Deposit(x Cash) error {
	if x <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	w.balance += x
	return nil
}

func (w *Wallet) Withdraw(x Cash) (error, Cash) {
	if x <= 0 {
		return errors.New("Withdrawal amount must be positive"), 0
	}

	if w.balance >= x {
		w.balance -= x
		return nil, x
	} else {
		return errors.New("Not enough cash in your wallet"), 0
	}
}

func (w *Wallet) Transfer(amount Cash, receiver *Wallet) error {
	err, cash := w.Withdraw(amount)
	if err != nil {
		return err
	}

	receiver.Deposit(cash)
	return nil
}
