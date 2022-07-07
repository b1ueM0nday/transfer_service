package gg

import (
	"context"
	p "github.com/b1uem0nday/transfer_service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/big"
)

func (gg *GrpcGateway) Deposit(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	err := gg.transferClient.UpdateNonce()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, gg.transferClient.Deposit(big.NewInt(int64(request.Amount)))
}

func (gg *GrpcGateway) Withdraw(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	err := gg.transferClient.UpdateNonce()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, gg.transferClient.Withdraw(big.NewInt(int64(request.Amount)))
}

func (gg *GrpcGateway) GetBalance(ctx context.Context, request *p.BalanceRequest) (*p.BalanceReply, error) {
	reply, err := gg.transferClient.GetBalance(request.AccountAddress)
	if err != nil {
		return nil, err
	}
	return &p.BalanceReply{Balance: reply.Uint64()}, nil

}

func (gg *GrpcGateway) Transfer(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	err := gg.transferClient.UpdateNonce()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, gg.transferClient.Transfer(*request.AccountAddress, big.NewInt(int64(request.Amount)))
}
