package contracts

import (
	"encoding/json"
	"fmt"
	balance_op "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
)

func (c *Client) log() {
	logs := make(chan types.Log)
	sub, err := c.wsClient.SubscribeFilterLogs(c.ctx, c.fq, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(balance_op.BalanceOpABI))
	if err != nil {
		log.Fatal(err)
	}
	var e []interface{}
	for {
		select {
		case <-c.ctx.Done():
			log.Println("contract context was closed")
			return
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println()
			for i := range vLog.Topics {
				switch vLog.Topics[i].Hex() {
				case balance_op.DepositTopicHash:

					e, err = contractAbi.Unpack("Deposit", vLog.Data)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Printf("Deposited %d", e[0])

					c.printBalance(nil, "Deposit")

				case balance_op.WithdrawalTopicHash:
					e, err = contractAbi.Unpack("Withdrawal", vLog.Data)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Printf("Withdrawed %d", e[0])
					c.printBalance(nil, "Withdraw")
				case balance_op.TransferTopicHash:
					e, err = contractAbi.Unpack("Transfer", vLog.Data)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Printf("Transferred %d to %s", e[1], e[0])
					if addr, ok := e[0].(common.Address); ok {
						c.printBalance(&addr, "Transfer")
					}

				default:
					log.Println("Unhandled  topic", vLog.Topics[i].Hex())

				}

			}
			txReceipt, err := c.wsClient.TransactionReceipt(c.ctx, vLog.TxHash)
			if err != nil {
				log.Println(err)
				continue
			}
			if txReceipt.Status == types.ReceiptStatusSuccessful {
				log.Printf("transaction completed succesfully, hash: %s", txReceipt.TxHash)
			} else {
				log.Printf("transaction execution failed, hash: %s", txReceipt.TxHash)
			}
			log.Printf("block hash: %s \t block number: %d\n", txReceipt.BlockHash, txReceipt.BlockNumber)
			log.Printf("gas used: %d \t cumulitative gas used: %d\n", txReceipt.GasUsed, txReceipt.CumulativeGasUsed)
			m, err := json.Marshal("some data")
			c.db.InsertLog(vLog.Topics[0].Hex(), m)
		}
	}
}
