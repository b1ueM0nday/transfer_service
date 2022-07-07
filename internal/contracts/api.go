package contracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func (c *Contract) Deposit(amount *big.Int) (err error) {

	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	_, err = c.contract.Deposit(c.defOpts, amount)

	return err

}
func (c *Contract) Withdraw(amount *big.Int) (err error) {

	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	_, err = c.contract.Withdraw(c.defOpts, amount)

	return err
}
func (c *Contract) Transfer(receiver string, amount *big.Int) (err error) {
	receiverAddress := common.HexToAddress(receiver)
	if c.defOpts.Nonce, err = c.getNonce(); err != nil {
		return err
	}

	_, err = c.contract.Transfer(c.defOpts, receiverAddress, amount)
	return err

}

func (c *Contract) GetBalance(accountAddress *string) (*big.Int, error) {
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
func (c *Contract) printBalance(adr *common.Address, operation string) (err error) {
	senderBalance, err := c.contract.Balances(nil, c.owner)
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
