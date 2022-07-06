package gg

import (
	"context"
	"github.com/b1uem0nday/transfer_service/internal/transfer"
	p "github.com/b1uem0nday/transfer_service/proto"
	"google.golang.org/grpc"
	"net"
)

const defaultPort = "5000"

type GrpcServer struct {
	transferClient *transfer.Client
	//	ctx            context.Context
	gs       *grpc.Server
	listener net.Listener
	p.UnimplementedTransferServiceServer
}

func New(ctx context.Context, client *transfer.Client) *GrpcServer {
	return &GrpcServer{
		transferClient: client,
		//	ctx:                                ctx,
		UnimplementedTransferServiceServer: p.UnimplementedTransferServiceServer{},
	}
}

func (gg *GrpcServer) Connect(port string) (err error) {
	if port == "" {
		port = defaultPort
	}
	gg.listener, err = net.Listen("tcp", ":"+port)

	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	gg.gs = grpc.NewServer(opts...)
	p.RegisterTransferServiceServer(gg.gs, gg)

	return gg.gs.Serve(gg.listener)
}
