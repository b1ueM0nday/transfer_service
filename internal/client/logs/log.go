package logs

import (
	"context"
	"encoding/json"
	bo "github.com/b1uem0nday/transfer_service/contract"
	"github.com/b1uem0nday/transfer_service/internal/repository"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Logger interface {
		Run(rawurl string, address common.Address) error
		HandleTransaction(ts *types.Transaction)
	}
	logger struct {
		logs         chan types.Log
		Transactions chan *types.Transaction
		db           repository.Repo
		Logger
	}
)

func NewLogger(db repository.Repo, tx chan *types.Transaction) *logger {
	return &logger{
		logs:         make(chan types.Log),
		Transactions: tx,
		db:           db,
	}
}

func (l *logger) Run(rawurl string, address common.Address) error {
	logs := make(chan types.Log)
	var op string
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	ws, err := ethclient.DialContext(ctx, rawurl)
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
				err = l.db.InsertReceipt(time.Now(), ops[txReceipt.Logs[0].Topics[0].Hex()], b)
			} else {
				err = l.db.InsertReceipt(time.Now(), opUndefined, b)
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
			l.db.InsertLog(now, op, byteLog)
		}
	}

	return nil
}

func (l *logger) HandleTransaction(tx *types.Transaction) {
	l.Transactions <- tx
}
