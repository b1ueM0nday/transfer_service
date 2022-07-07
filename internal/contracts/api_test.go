package contracts

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestContract_Deposit(t *testing.T) {
	var err error
	TestContract_Deploy(t)
	balance, err := tc.GetBalance(nil)
	if err != nil {
		fmt.Println(err)
		t.Fatalf("failed to get balance")
	}
	r := rand.NewSource(time.Now().Unix())
	amount := big.NewInt(r.Int63())

	nonce, err := blockchain.PendingNonceAt(tc.ctx, tc.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}
	tc.defOpts.Nonce = big.NewInt(int64(nonce))
	err = tc.Deposit(amount)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	afterBalance, err := tc.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	diff := big.NewInt(0)
	diff = diff.Sub(afterBalance, balance)
	if diff.Cmp(amount) != 0 {
		t.Fatalf("balance not changed")
	}
	fmt.Println(afterBalance, balance)
}

func TestContract_Withdraw(t *testing.T) {

	var err error

	TestContract_Deploy(t)
	r := rand.NewSource(time.Now().Unix())
	amount := big.NewInt(r.Int63())

	nonce, err := blockchain.PendingNonceAt(tc.ctx, tc.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}
	tc.defOpts.Nonce = big.NewInt(int64(nonce))
	err = tc.Deposit(amount)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	balance, err := tc.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	nonce, err = blockchain.PendingNonceAt(tc.ctx, tc.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}

	tc.defOpts.Nonce = big.NewInt(int64(nonce))
	err = tc.Withdraw(amount)
	if err != nil {
		t.Fatalf("failed to send withdraw")
	}
	blockchain.Commit()
	afterBalance, err := tc.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}

	diff := big.NewInt(0)
	diff = diff.Sub(balance, afterBalance)
	if diff.Cmp(amount) != 0 {
		t.Fatalf("balance not changed")
	}
	fmt.Println(afterBalance, balance)
}

func TestContract_Transfer(t *testing.T) {

}
