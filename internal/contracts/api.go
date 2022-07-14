package contracts

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func (c *Client) deposit(amount *big.Int, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {

	if amount.Sign() < 0 {
		return nil, fmt.Errorf("amount can't be negative")
	}
	return c.contract.Deposit(txOpts, amount)
}
func (c *Client) Deposit(amount *big.Int, txOpts *bind.TransactOpts) error {
	tx, err := c.deposit(amount, txOpts)
	c.log.HandleTransaction(tx)
	return err
}

func (c *Client) withdraw(amount *big.Int, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if amount.Sign() < 0 {
		return nil, fmt.Errorf("amount can't be negative")
	}
	return c.contract.Withdraw(txOpts, amount)
}
func (c *Client) Withdraw(amount *big.Int, txOpts *bind.TransactOpts) error {
	tx, err := c.withdraw(amount, txOpts)
	c.log.HandleTransaction(tx)
	return err
}
func (c *Client) Transfer(receiver string, amount *big.Int, txOpts *bind.TransactOpts) error {
	tx, err := c.transfer(receiver, amount, txOpts)
	c.log.HandleTransaction(tx)
	return err
}
func (c *Client) transfer(receiver string, amount *big.Int, txOpts *bind.TransactOpts) (tx *types.Transaction, err error) {
	if receiver == "" {
		return nil, errors.New("no receiver")
	}
	receiverAddress := common.HexToAddress(receiver)
	if amount.Sign() < 0 {
		return nil, fmt.Errorf("amount can't be negative")
	}
	return c.contract.Transfer(txOpts, receiverAddress, amount)
}

func (c *Client) GetBalance(accountAddress *string) (*big.Int, error) {
	if accountAddress != nil {
		return c.contract.Balances(nil, common.HexToAddress(*accountAddress))
	} else {
		owner, err := c.contract.Owner(nil)
		if err != nil {
			return nil, err
		}
		return c.contract.Balances(nil, owner)
	}
}
func (c *Client) printBalance(adr *common.Address, operation string) error {
	owner, err := c.contract.Owner(nil)
	if err != nil {
		return err
	}
	senderBalance, err := c.contract.Balances(nil, owner)
	if err != nil {
		return err
	}
	if adr != nil {
		recBalance, err := c.contract.Balances(nil, *adr)
		if err != nil {
			return err
		}
		log.Println(fmt.Sprintf("%s: \t sender balance: %d \t receiver balance: %d", operation, senderBalance, recBalance))
	} else {
		log.Println(fmt.Sprintf("%s: \t sender balance: %d \t", operation, senderBalance))
	}
	return nil
}
