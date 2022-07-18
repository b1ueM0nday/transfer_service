package client

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (c *Client) Deposit(amount *big.Int) error {
	txOpts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.deposit(amount, txOpts)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return err
}

func (c *Client) Withdraw(amount *big.Int) error {
	txOpts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.withdraw(amount, txOpts)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return err
}
func (c *Client) Transfer(receiver string, amount *big.Int) error {
	txOpts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.transfer(receiver, amount, txOpts)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return err
}
func (c *Client) GetBalance(accountAddress *string) (*big.Int, error) {
	if accountAddress != nil {
		return c.contract.Balances(nil, common.HexToAddress(*accountAddress))
	} else {
		return c.contract.Balances(nil, c.owner.address)
	}
}

func (c *Client) deposit(amount *big.Int, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if amount == nil {
		return nil, fmt.Errorf("amount can't be nil")
	}
	if amount.Sign() < 0 {
		return nil, fmt.Errorf("amount can't be negative")
	}

	return c.contract.Deposit(txOpts, amount)
}

func (c *Client) withdraw(amount *big.Int, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if amount == nil {
		return nil, fmt.Errorf("amount can't be nil")
	}
	if amount.Sign() < 0 {
		return nil, fmt.Errorf("amount can't be negative")
	}
	return c.contract.Withdraw(txOpts, amount)
}
func (c *Client) transfer(receiver string, amount *big.Int, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if receiver == "" {
		return nil, errors.New("no receiver")
	}
	receiverAddress := common.HexToAddress(receiver)
	if amount == nil {
		return nil, fmt.Errorf("amount can't be nil")
	}
	if amount.Sign() < 0 {
		return nil, fmt.Errorf("amount can't be negative")
	}
	return c.contract.Transfer(txOpts, receiverAddress, amount)
}
