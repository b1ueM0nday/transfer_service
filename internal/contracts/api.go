package contracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

func (c *Worker) Deposit(amount *big.Int) (err error) {
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}
	err = c.printBalance(nil, fmt.Sprintf("Before depositing in amount of %d", amount))
	if err != nil {
		return err
	}
	_, err = c.contract.Deposit(c.defOpts, amount)
	if err != nil {
		return err
	}

	return c.printBalance(nil, fmt.Sprintf("After depositing in amount of %d", amount))
}
func (c *Worker) Withdraw(amount *big.Int) (err error) {
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	err = c.printBalance(nil, fmt.Sprintf("Before withdrawing in amount of %d", amount))
	if err != nil {
		return err
	}
	_, err = c.contract.Withdraw(c.defOpts, amount)
	if err != nil {
		return err
	}

	return c.printBalance(nil, fmt.Sprintf("After withdrawing in amount of %d", amount))
}
func (c *Worker) Transfer(receiver string, amount *big.Int) (err error) {
	receiverAddress := common.HexToAddress(receiver)
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	err = c.printBalance(&receiverAddress, fmt.Sprintf("Before transfering in amount of %d", amount))
	if err != nil {
		return err
	}
	_, err = c.contract.Transfer(c.defOpts, receiverAddress, amount)
	if err != nil {
		return err
	}

	return c.printBalance(&receiverAddress, fmt.Sprintf("After transfering in amount of %d", amount))
}

func (c *Worker) GetBalance(accountAddress *string) (*big.Int, error) {
	if accountAddress != nil {
		return c.contract.Balances(nil, common.HexToAddress(*accountAddress))
	} else {
		aa, err := c.contract.Owner(nil)
		if err != nil {
			return nil, err
		}
		return c.contract.Balances(nil, aa)
	}
}
func (c *Worker) printBalance(adr *common.Address, operation string) (err error) {
	time.Sleep(time.Millisecond * 500)
	senderBalance, err := c.contract.Balances(nil, c.owner)
	if err != nil {
		return err
	}

	if adr != nil {
		recBalance, err := c.contract.Balances(nil, *adr)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("%s: \t sender balance: %d \t receiver balance: %d", operation, senderBalance, recBalance))
	} else {
		fmt.Println(fmt.Sprintf("%s: \t sender balance: %d \t", operation, senderBalance))
	}
	return nil
}
