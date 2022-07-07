package contracts

import (
	"context"
	"fmt"
	balance_op "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
)

type (
	Config struct {
		IP             string `yaml:"ip"`
		HttpPort       string `yaml:"http_port"`
		WsPort         string `yaml:"ws_port"`
		AddressPath    string `yaml:"address_path"`
		PrivateKeyPath string `yaml:"private_key_path"`
	}
	Contract struct {
		cfg *Config
		ctx context.Context
		fq  ethereum.FilterQuery
		//owner    common.Address
		contract *balance_op.BalanceOp
		ethCli   *ethclient.Client
		defOpts  *bind.TransactOpts
		wsClient *ethclient.Client
	}
)

var DefaultConfig = Config{
	IP:             "localhost",
	HttpPort:       "22000",
	WsPort:         "32000",
	AddressPath:    "",
	PrivateKeyPath: "",
}

func NewContract(ctx context.Context) *Contract {
	return &Contract{ctx: ctx}
}

func (c *Contract) Prepare(cfg *Config) (err error) {

	pk, err := ioutil.ReadFile(cfg.PrivateKeyPath)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(string(pk))
	if err != nil {
		return err
	}
	/*	publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return errors.New("error casting public key to ECDSA")
		}

		c.defOpts.From = crypto.PubkeyToAddress(*publicKeyECDSA)*/
	c.ethCli, err = ethclient.Dial(fmt.Sprintf("http://%s:%s", cfg.IP, cfg.HttpPort)) //json-rpc
	if err != nil {
		return err
	}
	chId, err := c.ethCli.ChainID(c.ctx)
	if err != nil {
		return err
	}
	c.defOpts, err = bind.NewKeyedTransactorWithChainID(privateKey, chId)
	gasPrice, err := c.ethCli.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	c.defOpts.Value = big.NewInt(0)      // in wei
	c.defOpts.GasLimit = uint64(3000000) // in units
	c.defOpts.GasPrice = gasPrice
	var contractAddress common.Address
	if b, err := ioutil.ReadFile(cfg.AddressPath); b == nil || err != nil {
		contractAddress, err = c.deploy(cfg.AddressPath)
		if err != nil {
			return err
		}
	} else {
		contractAddress = common.BytesToAddress(b)

		log.Println("using existent contract", contractAddress)
	}
	if err = c.setInstance(contractAddress); err != nil {
		return err
	}
	c.wsClient, err = ethclient.Dial(fmt.Sprintf("ws://%s:%s", cfg.IP, cfg.WsPort))
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

func (c *Contract) deploy(path string) (address common.Address, err error) {

	if nonce, err := c.getNonce(); err != nil {
		return common.Address{}, err
	} else {
		c.defOpts.Nonce = nonce
	}
	addr, _, _, err := balance_op.DeployBalanceOp(c.defOpts, c.ethCli)
	if err != nil {
		return common.Address{}, err
	}
	log.Println("contract deployed")
	if err = ioutil.WriteFile(path, addr.Bytes(), 0777); err != nil {
		log.Println("contract deployed, but address wasn't saved")
	}
	return addr, nil
}

func (c *Contract) getNonce() (*big.Int, error) {
	nonce, err := c.ethCli.PendingNonceAt(context.Background(), c.defOpts.From)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), err
}

func (c *Contract) setInstance(contractAddress common.Address) (err error) {
	if c.contract, err = balance_op.NewBalanceOp(contractAddress, c.ethCli); err != nil || c.contract == nil {
		log.Printf("contract %s was not deployed, deploy again", contractAddress)
		contractAddress, err = c.deploy(c.cfg.AddressPath)
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

func (c *Contract) UpdateNonce() error {
	nonce, err := c.getNonce()
	if err != nil {
		return err
	}
	c.defOpts.Nonce = nonce
	return nil
}
