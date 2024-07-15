package service

import (
	"context"
	pb "order-service/genprotos/order_pb"
)

// 11
func (s *OrderServiceSt) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	s.logger.Info("create payment request")
	return s.service.CreatePayment(ctx, req)
}
