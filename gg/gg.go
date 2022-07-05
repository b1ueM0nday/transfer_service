package gg

import (
	"context"
	p "github.com/b1uem0nday/transfer_service/proto"
	"github.com/b1uem0nday/transfer_service/transfer"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/big"
	"net"
)

type GrpcServer struct {
	transferClient *transfer.Client
	ctx            context.Context
	gs             *grpc.Server
	listener       net.Listener
	p.UnimplementedTransferServiceServer
}

func (gg *GrpcServer) Deposit(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.transferClient.Deposit(big.NewInt(int64(request.Amount)))
}

func (gg *GrpcServer) Withdraw(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.transferClient.Withdraw(big.NewInt(int64(request.Amount)))
}

func (gg *GrpcServer) GetBalance(ctx context.Context, request *p.BalanceRequest) (*p.BalanceReply, error) {
	reply, err := gg.transferClient.GetBalance(request.AccountAddress)
	if err != nil {
		return nil, err
	}
	return &p.BalanceReply{Balance: reply.Uint64()}, nil

}

func (gg *GrpcServer) Transfer(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.transferClient.Transfer(*request.AccountAddress, big.NewInt(int64(request.Amount)))
}

func New(ctx context.Context, client *transfer.Client) *GrpcServer {
	return &GrpcServer{
		transferClient:                     client,
		ctx:                                ctx,
		UnimplementedTransferServiceServer: p.UnimplementedTransferServiceServer{},
	}
}

func (gg *GrpcServer) Connect(port string) (err error) {
	gg.listener, err = net.Listen("tcp", ":"+port)

	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	gg.gs = grpc.NewServer(opts...)
	p.RegisterTransferServiceServer(gg.gs, gg)

	return gg.gs.Serve(gg.listener)
}
