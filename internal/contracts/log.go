package contracts

import (
	"context"
	"encoding/json"
	"github.com/b1uem0nday/transfer_service/internal/base"
	bo "github.com/b1uem0nday/transfer_service/internal/contracts/balance_operations"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"time"
)

var ops = map[string]string{
	bo.TransferTopicHash:   "Transfer",
	bo.DepositTopicHash:    "Deposit",
	bo.WithdrawalTopicHash: "Withdraw",
}

const opUndefined = "Undefined"

type (
	logger struct {
		logs         chan types.Log
		Transactions chan *types.Transaction
		base         *base.Database
	}
)

func NewLogger(db *base.Database, tx chan *types.Transaction) *logger {
	return &logger{
		logs:         make(chan types.Log),
		Transactions: tx,
		base:         db,
	}
}

func (l *logger) Run(rawurl string, address common.Address) error {
	logs := make(chan types.Log)
	var op string
	ws, err := connect(rawurl)
	if err != nil {
		return err
	}
	filter := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	sub, err := ws.SubscribeFilterLogs(context.Background(), filter, logs)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-context.Background().Done():
			return nil
		case err := <-sub.Err():
			log.Fatal(err)
		case tx := <-l.Transactions:
			if tx == nil {
				break
			}

			txReceipt, err := bind.WaitMined(context.Background(), ws, tx)
			if err != nil {
				log.Println(err)
				continue
			}

			if txReceipt.Status == types.ReceiptStatusSuccessful {
				log.Printf("transaction completed succesfully, hash: %s", txReceipt.TxHash)
			} else {
				log.Printf("transaction execution failed, hash: %s", txReceipt.TxHash)
			}

			b, err := json.Marshal(txReceipt)
			log.Printf("block hash: %s \t block number: %d\n", txReceipt.BlockHash, txReceipt.BlockNumber)
			log.Printf("gas used: %d \t cumulitative gas used: %d\n", txReceipt.GasUsed, txReceipt.CumulativeGasUsed)

			if _, ok := ops[txReceipt.Logs[0].Topics[0].Hex()]; ok {
				err = l.base.InsertReceipt(time.Now(), ops[txReceipt.Logs[0].Topics[0].Hex()], b)
			} else {
				err = l.base.InsertReceipt(time.Now(), opUndefined, b)
			}
			if err != nil {
				log.Println(err)
			}
		case vLog := <-logs:
			now := time.Now()
			for i := range vLog.Topics {
				if _, ok := ops[vLog.Topics[i].Hex()]; ok {
					op = ops[vLog.Topics[i].Hex()]
				} else {
					op = "Undefined"
					log.Println("Unhandled  topic", vLog.Topics[i].Hex())
					continue
				}
			}
			byteLog, err := json.Marshal(vLog)
			if err != nil {
				log.Println(err)
				continue
			}
			l.base.InsertLog(now, op, byteLog)
		}
	}

	return nil
}
