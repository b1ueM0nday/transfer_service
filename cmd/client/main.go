package main

import (
	"context"
	"fmt"
	"github.com/b1uem0nday/transfer_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	cc := proto.NewTransferServiceClient(conn)
	_, err = cc.Deposit(ctx, &proto.BalanceOperationRequest{
		Amount: 100500,
	})
	if err != nil {
		log.Panic(err)
	}
	balance, err := cc.GetBalance(ctx, &proto.BalanceRequest{AccountAddress: nil})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(balance.Balance)
	return
}
