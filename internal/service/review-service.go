package service

import (
	"context"
	"fmt"
	"log"
	"order-service/genprotos/auth_pb"
	pb "order-service/genprotos/order_pb"
)

// 9
func (s *OrderServiceSt) AddReview(ctx context.Context, req *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {
	s.logger.Info("add review request")
	return s.service.AddReview(ctx, req)
}

// 10
func (s *OrderServiceSt) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	s.logger.Info("list of reviews request")
	return s.service.ListReviews(ctx, req)
}

// 2.1
func (s *OrderServiceSt) GetDishRecommendations(ctx context.Context, req *pb.GetDishRecommendationsRequest) (*pb.GetDishRecommendationsResponse, error) {
	s.logger.Info("get dish recommendations request")
	return s.service.GetDishRecommendations(ctx, req)
}

// 2.4
func (s *OrderServiceSt) CreateKitchenWorkingHours(ctx context.Context, req *pb.CreateKitchenWorkingHoursRequest) (*pb.CreateKitchenWorkingHoursResponse, error) {
	exists, _ := s.aut_client.DoesKitchenExist(ctx, &auth_pb.DoesKitchenExistRequest{KitchenId: req.KitchenId})
	log.Println(exists)
	if !exists.Exists {
		s.logger.Error("Failed to check if kitchen exists", "error", fmt.Errorf("kitchen does not exist"))
		return nil, fmt.Errorf("kitchen does not exist")
	}
	s.logger.Info("create kitchen working hours request")

	return s.service.CreateKitchenWorkingHours(ctx, req)
}
