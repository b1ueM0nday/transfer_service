package client

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) Deposit(amount uint64) error {
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

func (c *Client) Withdraw(amount uint64) error {
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
func (c *Client) Transfer(receiver string, amount uint64) error {
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
func (c *Client) GetBalance(accountAddress *string) (uint64, error) {

	if accountAddress != nil {

		account, err := c.contract.Accounts(nil, common.HexToAddress(*accountAddress))
		if err != nil {
			return 0, err
		}
		return account.Balance, nil
	} else {

		account, err := c.contract.Accounts(nil, c.owner.address)
		if err != nil {
			return 0, err
		}
		return account.Balance, nil
	}
}

func (c *Client) deposit(amount uint64, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if amount == 0 {
		return nil, fmt.Errorf("amount can't be nil")
	}
	return c.contract.Deposit(txOpts, amount)
}

func (c *Client) withdraw(amount uint64, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if amount == 0 {
		return nil, fmt.Errorf("amount can't be nil")
	}
	return c.contract.Withdraw(txOpts, amount)
}
func (c *Client) transfer(receiver string, amount uint64, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if receiver == "" {
		return nil, errors.New("no receiver")
	}
	receiverAddress := common.HexToAddress(receiver)
	if amount == 0 {
		return nil, fmt.Errorf("amount can't be nil")
	}
	return c.contract.Transfer(txOpts, receiverAddress, amount)
}
