package service

import (
	"log"
	"log/slog"

	"order-service/genprotos/auth_pb"
	pb "order-service/genprotos/order_pb"
	"order-service/internal/config"
	"order-service/internal/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	OrderServiceSt struct {
		pb.UnimplementedOrderServiceServer
		aut_client auth_pb.AuthServiceClient
		service    storage.OrderSt
		logger     *slog.Logger
	}
)

func connect(port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func New(config *config.Config, service storage.OrderSt, logger *slog.Logger) *OrderServiceSt {
	return &OrderServiceSt{
		aut_client: auth_pb.NewAuthServiceClient(connect(config.Server.Auth_Port)),
		service:    service,
		logger:     logger,
	}
}
