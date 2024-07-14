package service

import (
	"context"
	pb "order-service/genprotos/order_pb"
)

// 1
func (s *OrderServiceSt) AddDish(ctx context.Context, in *pb.AddDishRequest) (*pb.AddDishResponse, error) {
	s.logger.Info("AddDish")
	return s.service.AddDish(ctx, in)
}

// 2
func (s *OrderServiceSt) UpdateDish(ctx context.Context, in *pb.UpdateDishRequest) (*pb.UpdateDishResponse, error) {
	s.logger.Info("UpdateDish")
	return s.service.UpdateDish(ctx, in)
}

// 3
func (s *OrderServiceSt) DeleteDish(ctx context.Context, in *pb.DeleteDishRequest) (*pb.DeleteDishResponse, error) {
	s.logger.Info("DeleteDish")
	return s.service.DeleteDish(ctx, in)
}

// 4
func (s *OrderServiceSt) ListDishes(ctx context.Context, in *pb.ListDishesRequest) (*pb.ListDishesResponse, error) {
	s.logger.Info("ListDishes")
	return s.service.ListDishes(ctx, in)
}
