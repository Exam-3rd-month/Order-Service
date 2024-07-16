package service

import (
	"context"
	"order-service/genprotos/auth_pb"
	pb "order-service/genprotos/order_pb"
)

// 11
func (s *OrderServiceSt) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	s.logger.Info("create payment request")
	resp, err := s.service.CreatePayment(ctx, req)
	if err != nil {
		s.logger.Error("create payment failed", "error", err)
		return nil, err
	}

	incReq := &auth_pb.IncrementTotalOrdersRequest{
		KitchenId:   resp.KitchenId,
		TotalOrders: 1,
	}

	_, err = s.aut_client.IncrementTotalOrders(ctx, incReq)
	if err != nil {
		s.logger.Error("increment total orders failed", "error", err)
		return nil, err
	}

	return resp, nil
}
