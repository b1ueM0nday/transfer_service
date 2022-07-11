package contracts

import (
	"context"
	"crypto/ecdsa"
	balance_op "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func CreateFakeContract() (*Client, error) {
	var err error
	c := NewClient(context.Background())
	c.owner, _ = crypto.GenerateKey()
	c.chainId = big.NewInt(1337)
	txOpts, err := bind.NewKeyedTransactorWithChainID(c.owner, c.chainId)
	if err != nil {
		return nil, err
	}

	txOpts.GasLimit = uint64(3000000)
	alloc := make(core.GenesisAlloc)
	alloc[txOpts.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}
	blockchain = backends.NewSimulatedBackend(alloc, txOpts.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		return nil, err
	}
	txOpts.GasPrice = gasPrice

	_, _, c.contract, err = balance_op.DeployBalanceOp(
		txOpts,
		blockchain,
	)

	if err != nil {
		return nil, err
	}

	blockchain.Commit()
	return c, nil
}

func TestContract_Deposit(t *testing.T) {
	fakeCli, err := CreateFakeContract()

	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)
	amount := big.NewInt(100500)

	err = fakeCli.Deposit(amount, options)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	afterBalance, err := fakeCli.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	if afterBalance.Cmp(amount) != 0 {
		t.Fatalf("balance and deposit sum are different")
	}
}

func TestContract_Withdraw(t *testing.T) {
	fakeCli, err := CreateFakeContract()

	options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.owner, fakeCli.chainId)

	amount := big.NewInt(100500)
	err = fakeCli.Deposit(amount, options)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	chargeAmount := big.NewInt(500)
	err = fakeCli.Withdraw(chargeAmount, options)
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

	err = fakeCli.Deposit(amount, options)
	if err != nil {
		t.Fatalf("failed to send deposit")
	}
	blockchain.Commit()
	strAddress := receiver.String()

	err = fakeCli.Transfer(strAddress, amount, options)
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
	fakeCli, err := CreateFakeContract()

	b, err := fakeCli.GetBalance(nil)
	if err != nil {
		t.Fatalf("failed to get balance")
	}
	if b == nil {
		t.Fatalf("balance is not correct")
	}
	if b.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("balance is not correct")
	}
}
