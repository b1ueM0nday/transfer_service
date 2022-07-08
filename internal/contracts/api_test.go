package contracts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestContract_Deposit(t *testing.T) {
	err := Deploy_FakeContract()
	if err != nil {
		t.Fatalf("Unable to deploy contract")
	}
	balance, err := FakeContract.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}

	amount := big.NewInt(100500)

	nonce, err := blockchain.PendingNonceAt(FakeContract.ctx, FakeContract.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}
	FakeContract.defOpts.Nonce = big.NewInt(int64(nonce))
	err = FakeContract.Deposit(amount)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	afterBalance, err := FakeContract.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	diff := big.NewInt(0)
	diff = diff.Sub(afterBalance, balance)
	if diff.Cmp(amount) != 0 {
		t.Fatalf("balance not changed")
	}
}

func TestContract_Withdraw(t *testing.T) {
	err := Deploy_FakeContract()
	if err != nil {
		t.Fatalf("Unable to deploy contract")
	}
	amount := big.NewInt(100500)

	nonce, err := blockchain.PendingNonceAt(FakeContract.ctx, FakeContract.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}
	FakeContract.defOpts.Nonce = big.NewInt(int64(nonce))
	err = FakeContract.Deposit(amount)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	balance, err := FakeContract.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	nonce, err = blockchain.PendingNonceAt(FakeContract.ctx, FakeContract.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}

	FakeContract.defOpts.Nonce = big.NewInt(int64(nonce))
	err = FakeContract.Withdraw(amount)
	if err != nil {
		t.Fatalf("failed to send withdraw")
	}
	blockchain.Commit()
	afterBalance, err := FakeContract.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}

	diff := big.NewInt(0)
	diff = diff.Sub(balance, afterBalance)
	if diff.Cmp(amount) != 0 {
		t.Fatalf("balance not changed")
	}
}

func TestContract_Transfer(t *testing.T) {
	err := Deploy_FakeContract()
	if err != nil {
		t.Fatalf("Unable to deploy contract")
	}

	receiverKey, _ := crypto.GenerateKey()
<<<<<<< Updated upstream

	receiverPublicKey := receiverKey.Public()
	publicKeyECDSA, ok := receiverPublicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Errorf("error casting public key to ECDSA")
	}

	receiver := crypto.PubkeyToAddress(*publicKeyECDSA)
=======
	receiverPublicKey := receiverKey.Public()
	publicKeyECDSA, _ := receiverPublicKey.(*ecdsa.PublicKey)
	receiver := crypto.PubkeyToAddress(*publicKeyECDSA)
	strAddress := receiver.String()

>>>>>>> Stashed changes
	amount := big.NewInt(100500)

	nonce, err := blockchain.PendingNonceAt(FakeContract.ctx, FakeContract.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}
	FakeContract.defOpts.Nonce = big.NewInt(int64(nonce))
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	err = FakeContract.Deposit(amount)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
<<<<<<< Updated upstream
	strAddress := receiver.String()
=======

>>>>>>> Stashed changes
	balance, err := FakeContract.GetBalance(&strAddress)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	nonce, err = blockchain.PendingNonceAt(FakeContract.ctx, FakeContract.defOpts.From)
	if err != nil {
		t.Fatalf("failed to get nonce")
	}
<<<<<<< Updated upstream

	FakeContract.defOpts.Nonce = big.NewInt(int64(nonce))
	err = FakeContract.Transfer(strAddress, amount)
	if err != nil {
		t.Fatalf("failed to transfer withdraw")
	}
	blockchain.Commit()
=======
	FakeContract.defOpts.Nonce = big.NewInt(int64(nonce))

	err = FakeContract.Transfer(strAddress, amount)
	if err != nil {
		t.Fatalf("failed to transfer")
	}
	blockchain.Commit()

>>>>>>> Stashed changes
	afterBalance, err := FakeContract.GetBalance(&strAddress)
	if err != nil {
		t.Fatalf("failed to get balance")
	}

	diff := big.NewInt(0)
	diff = diff.Sub(afterBalance, balance)
	if diff.Cmp(amount) != 0 {
		t.Fatalf("balance not changed")
	}
}
