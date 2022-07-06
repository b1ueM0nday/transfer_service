package main

import (
	"context"
	"flag"
	"github.com/b1uem0nday/transfer_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	port := flag.String("p", "3000", "port for grpc client")
	tReceiver := flag.String("tr", "", "transfer receiver")
	dAmount := flag.Uint64("da", 0, "deposit amount")
	wAmount := flag.Uint64("wa", 0, "withdraw amount")
	tAmount := flag.Uint64("ta", 0, "transfer amount")
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:"+*port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	cc := proto.NewTransferServiceClient(conn)
	_, err = cc.Deposit(ctx, &proto.BalanceOperationRequest{
		Amount: *dAmount,
	})
	if err != nil {
		log.Panic(err)
	}
	_, err = cc.Withdraw(ctx, &proto.BalanceOperationRequest{
		Amount: *wAmount,
	})
	if err != nil {
		log.Panic(err)
	}
	_, err = cc.Transfer(ctx, &proto.BalanceOperationRequest{
		Amount:         *tAmount,
		AccountAddress: tReceiver,
	})
	if err != nil {
		log.Panic(err)
	}

	balance, err := cc.GetBalance(ctx, &proto.BalanceRequest{AccountAddress: nil})
	if err != nil {
		log.Panic(err)
	}
	log.Println("sender balance", balance)
	balance, err = cc.GetBalance(ctx, &proto.BalanceRequest{AccountAddress: tReceiver})
	if err != nil {
		log.Panic(err)
	}
	log.Println("receiver balance", balance)
	return
}
