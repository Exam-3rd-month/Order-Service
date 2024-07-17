package service

import (
	"context"
	"errors"
	"fmt"
	"order-service/genprotos/auth_pb"
	pb "order-service/genprotos/order_pb"
)

// 9
func (s *OrderServiceSt) AddReview(ctx context.Context, req *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {
	exists, err := s.aut_client.DoesKitchenExist(ctx, &auth_pb.DoesKitchenExistRequest{KitchenId: req.KitchenId})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	
	if !exists.Exists {
		s.logger.Error("DoesKitchenExist failed", "error", errors.New("kitchen does not exist"))
		return nil, errors.New("kitchen does not exist")
	}

	exists1, err := s.aut_client.DoesUserExist(ctx, &auth_pb.DoesUserExistRequest{UserId: req.UserId})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	if !exists1.Exists {
		s.logger.Error("DoesUserExist failed", "error", errors.New("user does not exist"))
		return nil, errors.New("user does not exist")
	}

	s.logger.Info("add review request")

	resp, err := s.service.AddReview(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.aut_client.IncrementOrderRating(ctx, &auth_pb.IncrementOrderRatingRequest{
		KitchenId: resp.KitchenId,
		Rating:    resp.Rating,
	})

	return resp, nil
}

// 10
func (s *OrderServiceSt) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	s.logger.Info("list of reviews request")
	return s.service.ListReviews(ctx, req)
}

// 2.4
func (s *OrderServiceSt) CreateKitchenWorkingHours(ctx context.Context, req *pb.CreateKitchenWorkingHoursRequest) (*pb.CreateKitchenWorkingHoursResponse, error) {
	exists, _ := s.aut_client.DoesKitchenExist(ctx, &auth_pb.DoesKitchenExistRequest{KitchenId: req.KitchenId})
	if !exists.Exists {
		s.logger.Error("Failed to check if kitchen exists", "error", fmt.Errorf("kitchen does not exist"))
		return nil, fmt.Errorf("kitchen does not exist")
	}
	s.logger.Info("create kitchen working hours request")

	return s.service.CreateKitchenWorkingHours(ctx, req)
}
