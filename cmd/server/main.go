package main

import (
	"context"
	"github.com/b1uem0nday/transfer_service/internal/base"
	"github.com/b1uem0nday/transfer_service/internal/contracts"
	"github.com/b1uem0nday/transfer_service/internal/gg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const configPath = "./config/config.yaml"

type config struct {
	Grpc     gg.Config        `yaml:"grpc"`
	Contract contracts.Config `yaml:"node"`
	Base     struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		Port     uint   `yaml:"port"`
	} `yaml:"base"`
}

/*
const BuildNumber = "0.1"
const EnvVarPrefix = "TRANSFER_SERVICE_"
*/

func main() {
	cfg := loadConfig()
	ctx := context.Background()
	db := base.NewRepository(ctx)
	err := db.Connect(cfg.Base.Address, cfg.Base.Login, cfg.Base.Password, cfg.Base.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	client := contracts.NewClient(db, ctx)

	err = client.Prepare(&cfg.Contract)
	if err != nil {
		log.Fatal(err)
	}

	gs := gg.New(ctx, client)
	err = gs.Connect(cfg.Grpc.Port)
	if err != nil {
		log.Fatal(err)
	}
	gs.Run()
	<-ctx.Done()
	/*
		app := cli.NewApp()
		app.Name = "TRANSFER_SERVICE"
		app.Usage = "Test service for simple smart contract"
		app.Version = BuildNumber
		app.Flags = []cli.Flag{cli.StringFlag{
			Name:  "grpc-addr, ga",
			Usage: "grpc server address",
			Value: "127.0.0.1:3000",
		}}
		app.Before = connect
		app.Commands = []cli.Command{
			{
				Name:    "deposit",
				Aliases: []string{"d"},
				Usage:   "deposit amount to the owner account",
				Action:  deposit,
				Flags: []cli.Flag{cli.Uint64Flag{
					Name:   "deposit-amount, da",
					Value:  0,
					Usage:  "Amount of money to deposit",
					EnvVar: EnvVarPrefix + "DEPOSIT AMOUNT",
				}},
			},
			{
				Name:    "withdraw",
				Aliases: []string{"w"},
				Usage:   "withdraw amount from the owner account",
				Action:  withdraw,
				Flags: []cli.Flag{cli.Uint64Flag{
					Name:   "withdraw-amount, wa",
					Value:  0,
					Usage:  "Amount of money to withdraw",
					EnvVar: EnvVarPrefix + "WITHDRAW AMOUNT",
				}},
			},

			{
				Name:    "balance",
				Aliases: []string{"b"},
				Usage:   "check balance of user's account",
				Action:  getBalance,
				Flags: []cli.Flag{cli.StringFlag{
					Name:   "account, a",
					Value:  "",
					Usage:  "Address of the account",
					EnvVar: EnvVarPrefix + "ADDRESS",
				}},
			},
			{
				Name:    "transfer",
				Aliases: []string{"t"},
				Usage:   "transfer amount from the owner account to receiver account",
				Action:  transfer,
				Flags: []cli.Flag{
					cli.Uint64Flag{
						Name:   "transfer-amount, ta",
						Value:  0,
						Usage:  "Amount of money to withdraw",
						EnvVar: EnvVarPrefix + "TRANSFER AMOUNT",
					},
					cli.StringFlag{
						Name:   "transfer-receiver, tr",
						Value:  "",
						Usage:  "Amount of money to withdraw",
						EnvVar: EnvVarPrefix + "TRANSFER RECEIVER",
					},
				},
			},
		}

		if err := app.Run(os.Args); err != nil {
			os.Exit(1)
		}*/
}

func loadConfig() *config {

	var cfg config
	b, err := ioutil.ReadFile(configPath)
	if err != nil || b == nil {
		log.Println("run using default config")
		cfg = *defaultConfig()
	} else {
		err = yaml.Unmarshal(b, &cfg)
		if err != nil {
			log.Println("run using default config")
			cfg = *defaultConfig()
		}
	}
	return &cfg
}

func defaultConfig() *config {
	return &config{
		Grpc:     gg.DefaultConfig,
		Contract: contracts.DefaultConfig,
	}
}
