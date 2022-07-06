package transfer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type Client struct {
	eCli     *ethclient.Client
	sender   common.Address
	instance *Transfer
	defOpts  *bind.TransactOpts
}

func New() *Client {
	return &Client{}
}
func (c *Client) Deploy(address, pk string) (err error) {
	c.eCli, err = ethclient.Dial(address) //json-rpc
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	c.sender = crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.eCli.PendingNonceAt(context.Background(), c.sender)
	if err != nil {
		return err
	}

	gasPrice, err := c.eCli.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	chId, err := c.eCli.ChainID(context.Background())
	if err != nil {
		return err
	}
	c.defOpts, err = bind.NewKeyedTransactorWithChainID(privateKey, chId)
	if err != nil {
		return err
	}
	c.defOpts.From = c.sender
	c.defOpts.Nonce = big.NewInt(int64(nonce))
	c.defOpts.Value = big.NewInt(0)      // in wei
	c.defOpts.GasLimit = uint64(3000000) // in units
	c.defOpts.GasPrice = gasPrice

	_, _, c.instance, err = DeployTransfer(c.defOpts, c.eCli)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Deposit(amount *big.Int) (err error) {
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}
	err = c.printBalance(nil, fmt.Sprintf("Before depositing in amount of %d", amount))
	if err != nil {
		return err
	}
	_, err = c.instance.Deposit(c.defOpts, amount)
	if err != nil {
		return err
	}

	return c.printBalance(nil, fmt.Sprintf("After depositing in amount of %d", amount))
}
func (c *Client) Withdraw(amount *big.Int) (err error) {
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	err = c.printBalance(nil, fmt.Sprintf("Before withdrawing in amount of %d", amount))
	if err != nil {
		return err
	}
	_, err = c.instance.Withdraw(c.defOpts, amount)
	if err != nil {
		return err
	}

	return c.printBalance(nil, fmt.Sprintf("After withdrawing in amount of %d", amount))
}
func (c *Client) Transfer(receiver string, amount *big.Int) (err error) {
	receiverAddress := common.HexToAddress(receiver)
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	err = c.printBalance(&receiverAddress, fmt.Sprintf("Before transfering in amount of %d", amount))
	if err != nil {
		return err
	}
	_, err = c.instance.Transfer(c.defOpts, receiverAddress, amount)
	if err != nil {
		return err
	}

	return c.printBalance(&receiverAddress, fmt.Sprintf("After transfering in amount of %d", amount))
}

func (c *Client) GetBalance(accountAddress *string) (*big.Int, error) {
	if accountAddress != nil {
		return c.instance.Balances(nil, common.HexToAddress(*accountAddress))
	} else {
		return c.instance.Balances(nil, c.sender)
	}
}
func (c *Client) printBalance(adr *common.Address, operation string) (err error) {
	time.Sleep(time.Millisecond * 500)
	senderBalance, err := c.instance.Balances(nil, c.sender)
	if err != nil {
		return err
	}

	if adr != nil {
		recBalance, err := c.instance.Balances(nil, *adr)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("%s: \t sender balance: %d \t receiver balance: %d", operation, senderBalance, recBalance))
	} else {
		fmt.Println(fmt.Sprintf("%s: \t sender balance: %d \t", operation, senderBalance))
	}
	return nil
}

func (c *Client) getNonce() (*big.Int, error) {
	nonce, err := c.eCli.PendingNonceAt(context.Background(), c.sender)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), err
}
