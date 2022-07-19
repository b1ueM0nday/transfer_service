package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	market "github.com/b1uem0nday/transfer_service/contract"
	"github.com/b1uem0nday/transfer_service/internal/client/logs"
	"github.com/b1uem0nday/transfer_service/internal/repository"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"time"
)

type (
	Config struct {
		IP             string `yaml:"ip"`
		HttpPort       string `yaml:"http_port"`
		WsPort         string `yaml:"ws_port"`
		AddressPath    string `yaml:"address_path"`
		PrivateKeyPath string `yaml:"private_key_path"`
	}

	Api interface {
		Seller
		Buyer
		Deposit(amount uint64) error
		Withdraw(amount uint64) error
		Transfer(receiver string, amount uint64) error
		GetBalance(accountAddress *string) (uint64, error)
	}
	Client struct {
		chainId *big.Int
		cfg     *Config
		ctx     context.Context

		log      logs.Logger
		owner    ownerData
		contract *market.Market

		ethCli *ethclient.Client
		Api
	}
	txOpts struct {
		gasPrice *big.Int
		gasLimit uint64
		value    *big.Int
	}
	ownerData struct {
		pk      *ecdsa.PrivateKey
		address common.Address
	}
)

var DefaultConfig = Config{
	IP:             "localhost",
	HttpPort:       "22000",
	WsPort:         "32000",
	AddressPath:    "",
	PrivateKeyPath: "",
}

func NewClient(db repository.Repo, ctx context.Context) *Client {
	return &Client{log: logs.NewLogger(db, make(chan *types.Transaction)), ctx: ctx}
}

func (c *Client) Prepare(cfg *Config) (err error) {
	if c.owner.pk, err = readKeyFromFile(cfg.PrivateKeyPath); err != nil {
		return err
	}
	c.owner.address = crypto.PubkeyToAddress(c.owner.pk.PublicKey)
	c.ethCli, err = connect(fmt.Sprintf("http://%s:%s", cfg.IP, cfg.HttpPort)) //json-rpc
	if err != nil {
		return err
	}
	c.chainId, err = c.ethCli.ChainID(c.ctx)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(c.owner.pk, c.chainId)
	if err != nil {
		return err
	}
	var contractAddress common.Address
	if b, err := ioutil.ReadFile(cfg.AddressPath); b == nil || err != nil {
		contractAddress, err = c.deploy(cfg.AddressPath, auth)
		if err != nil {
			return err
		}
	} else {
		contractAddress = common.BytesToAddress(b)
		log.Println("using existent contract", contractAddress)
	}
	if err = c.setInstance(contractAddress, auth); err != nil {
		return err
	}

	go c.log.Run(fmt.Sprintf("ws://%s:%s", cfg.IP, cfg.WsPort), contractAddress)

	return nil
}

func (c *Client) deploy(path string, opts *bind.TransactOpts) (address common.Address, err error) {

	nonce, err := c.getNonce()
	if err != nil {
		return common.Address{}, err
	} else {
		opts.Nonce = nonce
	}
	addr, _, _, err := market.DeployMarket(opts, c.ethCli)
	if err != nil {
		return common.Address{}, err
	}
	log.Println("contract deployed")
	if err = ioutil.WriteFile(path, addr.Bytes(), 0777); err != nil {
		log.Println("contract deployed, but address wasn't saved")
	}
	return addr, nil
}

func (c *Client) getNonce() (*big.Int, error) {
	nonce, err := c.ethCli.PendingNonceAt(context.Background(), c.owner.address)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), err
}

func (c *Client) setInstance(contractAddress common.Address, opts *bind.TransactOpts) (err error) {
	if c.contract, err = market.NewMarket(contractAddress, c.ethCli); err != nil || c.contract == nil {
		log.Printf("contract %s was not deployed, deploy again", contractAddress)
		contractAddress, err = c.deploy(c.cfg.AddressPath, opts)
		if err != nil {
			return err
		}
		c.contract, err = market.NewMarket(contractAddress, c.ethCli)
		if err != nil {
			return err
		}
	}
	return nil
}

func connect(rawurl string) (c *ethclient.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	return ethclient.DialContext(ctx, rawurl)
}

func (c *Client) newTxOpts(opts ...txOpts) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(c.owner.pk, c.chainId)
	if err != nil {
		return nil, err
	}
	if txOpts.Nonce, err = c.getNonce(); err != nil {
		return nil, err
	}

	if len(opts) > 0 {
		if opts[0].gasPrice == nil {
			txOpts.GasPrice, err = c.ethCli.SuggestGasPrice(c.ctx)
			if err != nil {
				log.Printf("Cannot calculate suggested gas price due the transaction")
				txOpts.GasPrice = big.NewInt(0)
			}
		}
		txOpts.GasLimit = opts[0].gasLimit
		txOpts.Value = opts[0].value
	}

	return txOpts, nil
}

func (c *Client) NewTxOpts() (*bind.TransactOpts, error) {
	return c.newTxOpts()
}

func readKeyFromFile(path string) (key *ecdsa.PrivateKey, err error) {
	pk, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return crypto.HexToECDSA(string(pk))

}
