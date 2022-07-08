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
	blockchain   *backends.SimulatedBackend
	FakeContract = NewContract(context.Background())
)

func TestPrepare_EmptyConfig(t *testing.T) {

	badConfig := Config{
		IP:             "",
		HttpPort:       "",
		WsPort:         "",
		AddressPath:    "",
		PrivateKeyPath: "",
	}
	err := FakeContract.Prepare(&badConfig)
	if err == nil {
		t.Errorf("bad config didn't trigger an error")
	}
}

func Deploy_FakeContract() error {
	if FakeContract.contract != nil {
		return nil
	}
	var err error
	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	FakeContract.defOpts, err = bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return err
	}

	FakeContract.defOpts.GasLimit = uint64(3000000)
	alloc := make(core.GenesisAlloc)
	alloc[FakeContract.defOpts.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}
	blockchain = backends.NewSimulatedBackend(alloc, FakeContract.defOpts.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		return err
	}
	FakeContract.defOpts.GasPrice = gasPrice
	//Deploy contract
	_, _, FakeContract.contract, err = balance_op.DeployBalanceOp(
		FakeContract.defOpts,
		blockchain,
	)

	if err != nil {
		return err
	}

	blockchain.Commit()
	return nil
}
