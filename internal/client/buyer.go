package client

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

type (
	Buyer interface {
		RegisterAccount(usr *UserData) error
		UpgradeAccount() error
		UpdateUserData(name, phone, email *string) error
		GetAccountInfo()
		GetAccountItemsList()
		BuyItem()
	}
	UserData struct {
		Name     string
		Phone    string
		Email    string
		Birthday time.Time
		RegTime  time.Time
	}
)

func (c *Client) RegisterAccount(usr *UserData) error {
	usr.RegTime = time.Now()
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract1.RegisterAccount(opts, usr.Name, usr.Phone, usr.Email,
		uint64(usr.Birthday.UnixNano()), uint64(usr.RegTime.UnixNano()))
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}

func (c *Client) UpgradeAccount() error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract1.UpgradeAccount(opts)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}

func (c *Client) UpdateUserData(name, phone, email *string) error {
	if name != nil {
		err := c.updateUserData(*name, c.contract1.ChangeName)
		if err != nil {
			return err
		}
	}
	if phone != nil {
		err := c.updateUserData(*phone, c.contract1.ChangePhone)
		if err != nil {
			return err
		}
	}
	if email != nil {
		err := c.updateUserData(*email, c.contract1.ChangeEmail)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) updateUserData(data string, updater func(*bind.TransactOpts, string) (*types.Transaction, error)) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := updater(opts, data)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}
