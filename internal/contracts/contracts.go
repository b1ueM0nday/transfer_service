package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	balance_op "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
)

type Worker struct {
	owner           common.Address
	contractAddress common.Address
	contract        *balance_op.BalanceOp
	ethCli          *ethclient.Client
	defOpts         *bind.TransactOpts
}

func New() *Worker {
	return &Worker{}
}

func (c *Worker) Prepare(filepath, address, pkpath string) (err error) {
	pk, err := ioutil.ReadFile(pkpath)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(string(pk))
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	c.owner = crypto.PubkeyToAddress(*publicKeyECDSA)
	c.ethCli, err = ethclient.Dial(address) //json-rpc
	if err != nil {
		return err
	}
	chId, err := c.ethCli.ChainID(context.Background())
	if err != nil {
		return err
	}
	c.defOpts, err = bind.NewKeyedTransactorWithChainID(privateKey, chId)
	gasPrice, err := c.ethCli.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	c.defOpts.Value = big.NewInt(0)      // in wei
	c.defOpts.GasLimit = uint64(3000000) // in units
	c.defOpts.GasPrice = gasPrice
	if b, err := ioutil.ReadFile(filepath); b == nil || err != nil {
		c.contractAddress, err = c.deploy(filepath)
		if err != nil {
			return err
		}
	} else {

		c.contractAddress = common.BytesToAddress(b)
		log.Println("using existent contract", c.contractAddress)
	}

	c.contract, err = balance_op.NewBalanceOp(c.contractAddress, c.ethCli)
	if err != nil {
		return err
	}

	return nil
}

func (c *Worker) deploy(path string) (address common.Address, err error) {
	if nonce, err := c.getNonce(); err != nil {
		return common.Address{}, err
	} else {
		c.defOpts.Nonce = nonce
	}

	addr, _, _, err := balance_op.DeployBalanceOp(c.defOpts, c.ethCli)
	if err != nil {
		return common.Address{}, err
	}
	log.Println("contract deployed")
	if err = ioutil.WriteFile(path, addr.Bytes(), 0777); err != nil {
		log.Println("contract deployed, but address wasn't saved")
	}
	return addr, nil
}

func (c *Worker) getNonce() (*big.Int, error) {
	nonce, err := c.ethCli.PendingNonceAt(context.Background(), c.owner)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), err
}
