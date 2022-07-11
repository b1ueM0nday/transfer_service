package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/b1uem0nday/transfer_service/internal/base"
	balance_op "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	Client struct {
		chainId *big.Int
		cfg     *Config
		ctx     context.Context
		db      *base.Database
		fq      ethereum.FilterQuery

		owner     *ecdsa.PrivateKey
		ownerAddr common.Address
		contract  *balance_op.BalanceOp

		ethCli   *ethclient.Client
		wsClient *ethclient.Client
	}
	txOpts struct {
		gasPrice *big.Int
		gasLimit uint64
		value    *big.Int
	}
)

const ping = time.Second * 5

var DefaultConfig = Config{
	IP:             "localhost",
	HttpPort:       "22000",
	WsPort:         "32000",
	AddressPath:    "",
	PrivateKeyPath: "",
}

func NewClient(base *base.Database, ctx context.Context) *Client {
	return &Client{db: base, ctx: ctx}
}

func (c *Client) Prepare(cfg *Config) (err error) {

	pk, err := ioutil.ReadFile(cfg.PrivateKeyPath)
	if err != nil {
		return err
	}
	c.owner, err = crypto.HexToECDSA(string(pk))
	if err != nil {
		return err
	}
	c.ethCli, err = connect(fmt.Sprintf("http://%s:%s", cfg.IP, cfg.HttpPort)) //json-rpc
	if err != nil {
		return err
	}
	c.chainId, err = c.ethCli.ChainID(c.ctx)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(c.owner, c.chainId)
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

	if c.ownerAddr, err = c.contract.Owner(nil); err != nil {
		return err
	}

	c.wsClient, err = connect(fmt.Sprintf("ws://%s:%s", cfg.IP, cfg.WsPort))
	if err != nil {
		log.Printf("unable to create connection via web socket on port %s, run without logging\n", c.cfg.WsPort)
	} else {
		c.fq = ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}
		go c.log()
	}
	return nil
}

func (c *Client) deploy(path string, opts *bind.TransactOpts) (address common.Address, err error) {

	if nonce, err := c.getNonce(); err != nil {
		return common.Address{}, err
	} else {
		opts.Nonce = nonce
	}
	addr, _, _, err := balance_op.DeployBalanceOp(opts, c.ethCli)
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
	nonce, err := c.ethCli.PendingNonceAt(context.Background(), c.ownerAddr)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), err
}

func (c *Client) setInstance(contractAddress common.Address, opts *bind.TransactOpts) (err error) {
	if c.contract, err = balance_op.NewBalanceOp(contractAddress, c.ethCli); err != nil || c.contract == nil {
		log.Printf("contract %s was not deployed, deploy again", contractAddress)
		contractAddress, err = c.deploy(c.cfg.AddressPath, opts)
		if err != nil {
			return err
		}
		c.contract, err = balance_op.NewBalanceOp(contractAddress, c.ethCli)
		if err != nil {
			return err
		}
	}
	return nil
}

func connect(rawurl string) (c *ethclient.Client, err error) {
	for {
		c, err = ethclient.Dial(rawurl)
		if err == nil && c != nil {
			return c, err
		}
		time.Sleep(ping)
	}
}

func (c *Client) newTxOpts(opts ...txOpts) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(c.owner, c.chainId)
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
