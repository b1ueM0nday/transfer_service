package client

import (
	"bytes"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type (
	Seller interface {
		AddItem(item *MarketItem) error
		RemoveItemByName(name string) error
		RemoveItemByCode(code string) error
		UpdateItem(code, newDesc string, newPrice *big.Int, newCount uint64)
	}
)

func (c *Client) AddItem(item *MarketItem) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}

	hash := getHash([]byte(item.Name), crypto.FromECDSA(c.owner.pk))
	item.VendorCode = string(hash[:])
	tx, err := c.contract1.AddItem(opts, item.VendorCode, item.Name, item.Description,
		item.Price, big.NewInt(int64(item.Count)))
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}

func (c *Client) RemoveItemByName(name string) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	hash := getHash([]byte(name), crypto.FromECDSA(c.owner.pk))
	vendorCode := string(hash[:])
	tx, err := c.contract1.RemoveItem(opts, vendorCode)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}
func (c *Client) RemoveItemByCode(code string) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract1.RemoveItem(opts, code)
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}

func (c *Client) UpdateItem(code, newDesc string, newPrice *big.Int, newCount uint64) error {
	opts, err := c.NewTxOpts()
	if err != nil {
		return err
	}
	tx, err := c.contract1.UpdateItem(opts, code, newDesc, newPrice, big.NewInt(int64(newCount)))
	if err != nil {
		return err
	}
	c.log.HandleTransaction(tx)
	return nil
}
func getHash(opts ...[]byte) [32]byte {
	var data [][]byte
	for i := range opts {
		data = append(data, opts[i])
	}
	return sha256.Sum256(bytes.Join(data, []byte{}))
}
