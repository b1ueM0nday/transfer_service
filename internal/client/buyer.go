package client

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
		Birthday time.Time
		RegTime  time.Time
	}
	MarketItem struct {
		VendorCode  string
		Name        string
		Description string
		Price       *big.Int
		Count       uint64
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
func (c *Client) ChangeName(name string) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract1.ChangeName(opts, name)
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
	tx, err := c.contract1.ChangePhone(opts, phone)
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
	tx, err := c.contract1.ChangeEmail(opts, email)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}

func (c *Client) GetAccountInfo(address string) (*UserData, error) {
	data, err := c.contract1.GetAccountInfo(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &UserData{
		Name:     data.Name,
		Phone:    data.Phone,
		Email:    data.Email,
		Birthday: time.Unix(0, int64(data.Birthday)),
		RegTime:  time.Unix(0, int64(data.RegTime)),
	}, nil
}

func (c *Client) GetAccountItemsList() ([]MarketItem, error) {
	data, err := c.contract1.GetAccountItemsList(nil)
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
			Count:       data[i].Count.Uint64(),
		}
	}
	return items, nil

}
func (c *Client) BuyItem(seller, code string, count uint64) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract1.BuyItem(opts, common.HexToAddress(seller), code, big.NewInt(int64(count)))
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return err
}
