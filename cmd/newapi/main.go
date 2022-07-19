package main

import (
	"context"
	"fmt"
	"github.com/b1uem0nday/transfer_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {
	fmt.Println("211")
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	//birthday, err := time.Parse(time.RFC822, "02 Jan 96 15:04 MST")
	if err != nil {
		log.Panic(err)
	}
	cc := proto.NewTransferServiceClient(conn)
	/*_, err = cc.RegisterAccount(ctx, &proto.RegisterRequest{
		Name:     "John Doe",
		Phone:    "1499300024",
		Email:    "vitunsaatanaperkele@gmail.com",
		Birthday: uint64(birthday.UnixNano()),
	})*/
	_, err = cc.Deposit(ctx, &proto.BalanceOperationRequest{
		AccountAddress: nil,
		Amount:         10000000000000,
	})
	_, err = cc.UpgradeAccount(ctx, &emptypb.Empty{})
	if err != nil {
		log.Panic(err)
	}
	address := "c04a83e21e4993792e8ee2f71bc5965d7d59123bc0d407c04053c1fdb8500c16"
	user, err := cc.GetAccountInfo(ctx, &proto.GetAccountInfoRequest{Address: address})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(user)
}
