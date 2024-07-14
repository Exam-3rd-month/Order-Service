package service

import (
	"context"
	pb "order-service/genprotos/order_pb"
)

// 5
func (s *OrderServiceSt) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	s.logger.Info("create order request")
	return s.service.CreateOrder(ctx, req)
}

// 6
func (s *OrderServiceSt) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	s.logger.Info("update order status request")
	return s.service.UpdateOrderStatus(ctx, req)
}

// 7
func (s *OrderServiceSt) ListOfOrders(ctx context.Context, req *pb.ListOfOrdersRequest) (*pb.ListOfOrdersResponse, error) {
	s.logger.Info("list of orders request")
	return s.service.ListOfOrders(ctx, req)
}

// 8
func (s *OrderServiceSt) GetOrderByKitchenId(ctx context.Context, req *pb.GetOrderByKitchenIdRequest) (*pb.GetOrderByKitchenIdResponse, error) {
	s.logger.Info("get order by kitchen id request")
	return s.service.GetOrderByKitchenId(ctx, req)
}

// 12
func (s *OrderServiceSt) GetFullInfoAboutOrder(ctx context.Context, req *pb.GetFullInfoAboutOrderRequest) (*pb.GetFullInfoAboutOrderResponse, error) {
	s.logger.Info("get full info about order request")
	return s.service.GetFullInfoAboutOrder(ctx, req)
}
