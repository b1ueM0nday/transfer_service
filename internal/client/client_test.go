package client

import (
	"context"
	balance_op "github.com/b1uem0nday/transfer_service/internal/client/balance_operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

type TestClient struct {
	cli          *Client
	contractAddr common.Address

	blockchain *backends.SimulatedBackend
}

func CreateFakeContract() (*TestClient, error) {
	var err error
	c := NewClient(nil, context.Background())
	c.owner.pk, _ = crypto.GenerateKey()
	c.chainId = big.NewInt(1337)
	txOpts, err := bind.NewKeyedTransactorWithChainID(c.owner.pk, c.chainId)
	if err != nil {
		return nil, err
	}

	txOpts.GasLimit = uint64(3000000)
	alloc := make(core.GenesisAlloc)
	alloc[txOpts.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, txOpts.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		return nil, err
	}
	txOpts.GasPrice = gasPrice
	var addr common.Address
	addr, _, c.contract, err = balance_op.DeployBalanceOp(
		txOpts,
		blockchain,
	)

	if err != nil {
		return nil, err
	}

	blockchain.Commit()
	c.owner.address, _ = c.contract.Owner(nil)
	return &TestClient{
		cli:          c,
		contractAddr: addr,
		blockchain:   blockchain,
	}, nil
}
func TestPrepare_EmptyConfig(t *testing.T) {
	emptyConfig := Config{
		IP:             "",
		HttpPort:       "",
		WsPort:         "",
		AddressPath:    "",
		PrivateKeyPath: "",
	}
	c := NewClient(nil, context.Background())
	err := c.Prepare(&emptyConfig)
	if err == nil {
		t.Errorf("bad config didn't trigger an error")
	}
}
