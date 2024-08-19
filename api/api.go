package api

import (
	"log"
	"net"

	"order-service/genprotos/order_pb"
	"order-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		service order_pb.OrderServiceServer
	}
)

func New(service order_pb.OrderServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", "order"+config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	order_pb.RegisterOrderServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
