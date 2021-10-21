package pointer_test

import (
	"fmt"
	"testing"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Printf("address of balance in Deposit  is %v \n", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {

	return w.balance
}
func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
