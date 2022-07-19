package gg

import (
	"context"
	"github.com/b1uem0nday/transfer_service/internal/client"
	p "github.com/b1uem0nday/transfer_service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (gg *GrpcGateway) Deposit(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.client.Deposit(request.Amount)
}

func (gg *GrpcGateway) Withdraw(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.client.Withdraw(request.Amount)
}

func (gg *GrpcGateway) GetBalance(ctx context.Context, request *p.BalanceRequest) (*p.BalanceReply, error) {
	reply, err := gg.client.GetBalance(request.AccountAddress)
	if err != nil {
		return nil, err
	}
	return &p.BalanceReply{Balance: reply}, nil

}

func (gg *GrpcGateway) Transfer(ctx context.Context, request *p.BalanceOperationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.client.Transfer(*request.AccountAddress, request.Amount)
}

func (gg *GrpcGateway) AddItem(ctx context.Context, request *p.AddItemRequest) (*p.AddItemVendorCodeReply, error) {
	item := client.MarketItem{
		VendorCode:  "",
		Name:        request.Name,
		Description: request.Desc,
		Price:       request.Price,
		Count:       request.Count,
	}
	err := gg.client.AddItem(&item)
	return &p.AddItemVendorCodeReply{Code: item.VendorCode}, err
}

func (gg *GrpcGateway) RemoveItem(ctx context.Context, of *p.RemoveItemOneOf) (*emptypb.Empty, error) {
	return nil, nil
}

func (gg *GrpcGateway) UpdateItem(ctx context.Context, request *p.UpdateItemRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.client.UpdateItem(request.Code, request.Desc, request.Price, request.Count)
}

func (gg *GrpcGateway) RegisterAccount(ctx context.Context, request *p.RegisterRequest) (*emptypb.Empty, error) {
	user := client.UserData{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Birthday: request.Birthday,
	}
	return &emptypb.Empty{}, gg.client.RegisterAccount(&user)

}

func (gg *GrpcGateway) UpgradeAccount(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.client.UpgradeAccount()
}

func (gg *GrpcGateway) UpdateAccount(ctx context.Context, request *p.UpdateAccountRequest) (*emptypb.Empty, error) {
	if request.Name != nil {
		err := gg.client.ChangeName(*request.Name)
		if err != nil {
			return &emptypb.Empty{}, err
		}
	}
	if request.Phone != nil {
		err := gg.client.ChangePhone(*request.Phone)
		if err != nil {
			return &emptypb.Empty{}, err
		}
	}
	if request.Email != nil {
		err := gg.client.ChangeEmail(*request.Email)
		if err != nil {
			return &emptypb.Empty{}, err
		}
	}
	return &emptypb.Empty{}, nil
}

func (gg *GrpcGateway) GetAccountInfo(ctx context.Context, request *p.GetAccountInfoRequest) (*p.GetAccountInfoReply, error) {
	info, err := gg.client.GetAccountInfo(request.Address)
	if err != nil {
		return nil, err
	}
	return &p.GetAccountInfoReply{
		Name:     info.Name,
		Phone:    info.Phone,
		Email:    info.Email,
		Birthday: info.Birthday,
		RegTime:  info.RegTime,
	}, nil
}

func (gg *GrpcGateway) GetSellerItemsList(ctx context.Context, request *p.GetSellerItemsListRequest) (*p.GetSellerItemsListReply, error) {
	list, err := gg.client.GetAccountItemsList()
	if err != nil {
		return nil, err
	}
	var items p.GetSellerItemsListReply
	for i := range list {
		items.Items = append(items.Items, &p.SellerItem{
			VendorCode: list[i].VendorCode,
			Name:       list[i].Name,
			Desc:       list[i].Description,
			Price:      list[i].Price,
			Count:      list[i].Count,
		})
	}
	return &items, nil
}

func (gg *GrpcGateway) Buy(ctx context.Context, request *p.BuyRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, gg.client.BuyItem(request.Seller, request.Code, request.Count)
}
