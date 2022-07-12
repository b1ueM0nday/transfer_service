package contracts

import (
	"context"
	balance_op "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

var (
	blockchain *backends.SimulatedBackend
)

func CreateFakeContract() (*Client, error) {
	var err error
	c := NewClient(nil, context.Background())
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
func TestPrepare_EmptyConfig(t *testing.T) {
	badConfig := Config{
		IP:             "",
		HttpPort:       "",
		WsPort:         "",
		AddressPath:    "",
		PrivateKeyPath: "",
	}
	c := NewClient(nil, context.Background())
	err := c.Prepare(&badConfig)
	if err == nil {
		t.Errorf("bad config didn't trigger an error")
	}
}
func TestDeploy(t *testing.T) {
	c, err := CreateFakeContract()
	if err != nil || c == nil {
		t.Fatalf("failed contract deployment")
	}
}
