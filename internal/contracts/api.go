package contracts

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func (c *Client) Deposit(amount *big.Int, txOpts *bind.TransactOpts) (err error) {
	if txOpts != nil && amount != nil {
		_, err = c.contract.Deposit(txOpts, amount)
	}
	return err

}
func (c *Client) Withdraw(amount *big.Int, txOpts *bind.TransactOpts) (err error) {
	if txOpts != nil && amount != nil {
		_, err = c.contract.Withdraw(txOpts, amount)
	}
	return err
}
func (c *Client) Transfer(receiver string, amount *big.Int, txOpts *bind.TransactOpts) (err error) {
	if receiver == "" {
		return errors.New("no receiver")
	}
	receiverAddress := common.HexToAddress(receiver)
	if txOpts != nil && amount != nil {
		_, err = c.contract.Transfer(txOpts, receiverAddress, amount)
	}
	return err
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
