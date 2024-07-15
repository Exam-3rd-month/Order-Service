package service

import (
	"context"
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
