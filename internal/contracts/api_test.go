package contracts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestContract_Deposit(t *testing.T) {
	fakeCli, err := CreateFakeContract()

	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)
	amount := big.NewInt(100500)
	tx, err := fakeCli.deposit(amount, options)

	blockchain.Commit()
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	if tx == nil {
		t.Fatalf("deposit returns empty transaction")
	}
	afterBalance, err := fakeCli.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	if afterBalance.Cmp(amount) != 0 {
		t.Fatalf("balance and deposit sum are different")
	}
}

func TestClient_Deposit_NegativeAmount(t *testing.T) {
	fakeCli, _ := CreateFakeContract()
	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)
	amount := big.NewInt(-1000)

	_, err := fakeCli.deposit(amount, options)

	blockchain.Commit()
	if err == nil {
		t.Fatalf("sended negative deposit")
	}
}

func TestClient_Withdraw_NegativeAmount(t *testing.T) {
	fakeCli, _ := CreateFakeContract()
	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)
	amount := big.NewInt(-1000)

	_, err := fakeCli.withdraw(amount, options)

	blockchain.Commit()

	if err == nil {
		t.Fatalf("sended negative withdraw")
	}
}

func TestClient_Withdraw_InsufficientFunds(t *testing.T) {
	fakeCli, _ := CreateFakeContract()
	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)
	amount := big.NewInt(500)

	fakeCli.deposit(amount, options)

	blockchain.Commit()

	_, err := fakeCli.withdraw(amount.Neg(amount), options)

	blockchain.Commit()
	if err == nil {
		t.Fatal("error \"InsufficientFunds\" was not triggered on withdraw")
	}
}

func TestContract_Withdraw(t *testing.T) {
	fakeCli, err := CreateFakeContract()

	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)

	amount := big.NewInt(100500)
	_, err = fakeCli.deposit(amount, options)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	chargeAmount := big.NewInt(500)
	_, err = fakeCli.withdraw(chargeAmount, options)
	if err != nil {
		t.Fatalf("failed to send withdraw")
	}
	blockchain.Commit()
	afterBalance, err := fakeCli.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}

	if afterBalance.Cmp(big.NewInt(100000)) != 0 {
		t.Fatalf("balance was not charged correctly")
	}
}

func TestContract_Transfer(t *testing.T) {

	fakeCli, err := CreateFakeContract()

	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)

	receiverKey, _ := crypto.GenerateKey()

	receiverPublicKey := receiverKey.Public()
	publicKeyECDSA, ok := receiverPublicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Errorf("error casting public key to ECDSA")
	}

	receiver := crypto.PubkeyToAddress(*publicKeyECDSA)
	amount := big.NewInt(100500)

	_, err = fakeCli.deposit(amount, options)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	strAddress := receiver.String()

	_, err = fakeCli.transfer(strAddress, amount, options)
	if err != nil {
		t.Fatalf("failed to transfer withdraw")
	}
	blockchain.Commit()
	afterBalance, err := fakeCli.GetBalance(&strAddress)
	if err != nil {
		t.Fatalf("failed to get balance")
	}

	if afterBalance.Cmp(amount) != 0 {
		t.Fatalf("balance and deposit sum are different")
	}
}

func TestClient_GetBalance(t *testing.T) {
	fakeCli, _ := CreateFakeContract()

	b, err := fakeCli.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	if b == nil {
		t.Fatalf("balance is nil")
	}
	if b.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("balance is not zero on empty account")
	}
}
