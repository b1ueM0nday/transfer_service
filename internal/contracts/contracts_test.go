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
	tc         = NewContract(context.Background())
)

func TestPrepare_EmptyConfig(t *testing.T) {

	badConfig := Config{
		IP:             "",
		HttpPort:       "",
		WsPort:         "",
		AddressPath:    "",
		PrivateKeyPath: "",
	}
	err := tc.Prepare(&badConfig)
	if err == nil {
		t.Errorf("bad config didn't trigger an error")
	}
}

func TestContract_Deploy(t *testing.T) {
	var err error
	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	tc.defOpts, err = bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("Failed to create NewKeyedTransactorWithChainID : %v", err)
	}

	tc.defOpts.GasLimit = uint64(3000000)
	alloc := make(core.GenesisAlloc)
	alloc[tc.defOpts.From] = core.GenesisAccount{Balance: big.NewInt(100000000000000000)}
	blockchain = backends.NewSimulatedBackend(alloc, tc.defOpts.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		t.Fatalf("Failed to get SuggestGasPrice : %v", err)
	}
	tc.defOpts.GasPrice = gasPrice
	//Deploy contract
	_, _, tc.contract, err = balance_op.DeployBalanceOp(
		tc.defOpts,
		blockchain,
	)

	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	blockchain.Commit()

}
