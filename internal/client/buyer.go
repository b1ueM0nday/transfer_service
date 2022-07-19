package client

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type (
	Buyer interface {
		RegisterAccount(usr *UserData) error
		UpgradeAccount() error
		ChangeName(name string) error
		ChangePhone(phone string) error
		ChangeEmail(email string) error
		GetAccountInfo(address string) (*UserData, error)
		GetAccountItemsList() ([]MarketItem, error)
		BuyItem(seller, code string, count uint64) error
	}
	UserData struct {
		Name     string
		Phone    string
		Email    string
		Birthday uint64
		RegTime  uint64
	}
	MarketItem struct {
		VendorCode  string
		Name        string
		Description string
		Price       uint64
		Count       uint64
	}
)

func (c *Client) RegisterAccount(usr *UserData) error {
	usr.RegTime = uint64(time.Now().UnixNano())
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract.RegisterAccount(opts, usr.Name, usr.Phone, usr.Email,
		usr.Birthday, usr.RegTime)
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
	tx, err := c.contract.UpgradeAccount(opts)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}
func (c *Client) ChangeName(name string) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract.ChangeName(opts, name)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}
func (c *Client) ChangePhone(phone string) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract.ChangePhone(opts, phone)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}
func (c *Client) ChangeEmail(email string) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract.ChangeEmail(opts, email)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}

func (c *Client) GetAccountInfo(address string) (*UserData, error) {
	data, err := c.contract.GetAccountInfo(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &UserData{
		Name:     data.Name,
		Phone:    data.Phone,
		Email:    data.Email,
		Birthday: data.Birthday,
		RegTime:  data.RegTime,
	}, nil
}

func (c *Client) GetAccountItemsList() ([]MarketItem, error) {
	data, err := c.contract.GetAccountItemsList(nil)
	if err != nil {
		return nil, err
	}
	items := make([]MarketItem, len(data))
	for i := range data {
		items[i] = MarketItem{
			VendorCode:  data[i].VendorCode,
			Name:        data[i].Name,
			Description: data[i].Description,
			Price:       data[i].Price,
			Count:       data[i].Count,
		}
	}
	return items, nil

}
func (c *Client) BuyItem(seller, code string, count uint64) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract.BuyItem(opts, common.HexToAddress(seller), code, count)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return err
}
