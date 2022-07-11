package main

import (
	"context"
	"github.com/b1uem0nday/transfer_service/proto"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

const BuildNumber = "0.1"
const EnvVarPrefix = "TRANSFER_SERVICE_"

var client proto.TransferServiceClient

func main() {

	app := cli.NewApp()
	app.Name = "TRANSFER_SERVICE"
	app.Usage = "Test service for simple smart contract"
	app.Version = BuildNumber
	servAddress := cli.StringFlag{
		Name:   "grpc-addr, ga",
		Usage:  "grpc server address",
		Value:  "127.0.0.1:3000",
		EnvVar: EnvVarPrefix + "GRPC ADDRESS",
	}
	app.Commands = []cli.Command{
		{
			Name:    "deposit",
			Aliases: []string{"d"},
			Usage:   "deposit amount to the owner account",
			Before:  connect,
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
			Before:  connect,
			Action:  withdraw,
			Flags: []cli.Flag{cli.Uint64Flag{
				Name:   "withdraw-amount, wa",
				Value:  0,
				Usage:  "Amount of money to withdraw",
				EnvVar: EnvVarPrefix + "WITHDRAW AMOUNT",
			}, servAddress},
		},

		{
			Name:    "balance",
			Aliases: []string{"b"},
			Usage:   "check balance of user's account",
			Before:  connect,
			Action:  getBalance,
			Flags: []cli.Flag{cli.StringFlag{
				Name:   "account, a",
				Value:  "",
				Usage:  "Address of the account",
				EnvVar: EnvVarPrefix + "ADDRESS",
			}, servAddress},
		},
		{
			Name:    "transfer",
			Aliases: []string{"t"},
			Usage:   "transfer amount from the owner account to receiver account",
			Before:  connect,
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
				}, servAddress,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func connect(c *cli.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	conn, err := grpc.DialContext(ctx, c.String("ga"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer cancel()
	if err != nil {
		return err
	}
	client = proto.NewTransferServiceClient(conn)

	return nil
}
func deposit(c *cli.Context) {
	_, err := client.Deposit(context.Background(), &proto.BalanceOperationRequest{
		Amount: c.Uint64("deposit-amount"),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func withdraw(c *cli.Context) {
	_, err := client.Withdraw(context.Background(), &proto.BalanceOperationRequest{
		Amount: c.Uint64("withdraw-amount"),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func transfer(c *cli.Context) {
	receiver := c.String("transfer-receiver")
	_, err := client.Transfer(context.Background(), &proto.BalanceOperationRequest{
		AccountAddress: &receiver,
		Amount:         c.Uint64("transfer-amount"),
	})

	if err != nil {
		log.Fatal(err)
	}
}

func getBalance(c *cli.Context) {

	if account := c.String("account"); account == "" {
		balance, err := client.GetBalance(context.Background(), &proto.BalanceRequest{AccountAddress: nil})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Your balance: %d", balance.Balance)
	} else {
		balance, err := client.GetBalance(context.Background(), &proto.BalanceRequest{AccountAddress: &account})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Account %s balance: %d", account, balance.Balance)
	}
}
