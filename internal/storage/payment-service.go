package storage

import (
	"context"
	"fmt"
	pb "order-service/genprotos/order_pb"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type getTotalAmountByOrderId struct {
	total_amount float64
	kitchen_id   string
}

func isValidNumber(cardNumber string, number int) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	match, _ := regexp.MatchString("^[0-9]+$", cardNumber)
	if !match {
		return false
	}
	return len(cardNumber) == number
}

func isCardValid(expirationDate string) bool {
	parts := strings.Split(expirationDate, "/")
	if len(parts) != 2 {
		return false
	}

	expMonth, err1 := strconv.Atoi(parts[0])
	expYear, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || expMonth < 1 || expMonth > 12 {
		return false
	}

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()

	if expYear+2000 > currentYear || (expYear+2000 == currentYear && expMonth >= int(currentMonth)) {
		return true
	}

	return false
}

func (s *OrderSt) getTotalAmountByOrderId(order_id string) (*getTotalAmountByOrderId, error) {
	query, args, err := s.queryBuilder.Select("total_amount", "kitchen_id").
		From("orders").
		Where("order_id =?", order_id).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var str getTotalAmountByOrderId
	row := s.db.QueryRowContext(context.Background(), query, args...)
	err = row.Scan(&str.total_amount, &str.kitchen_id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &str, nil
}

// 11
func (s *OrderSt) CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	if in.PaymentMethod != "credit_card" {
		return nil, fmt.Errorf("invalid payment method")
	} else if !isValidNumber(in.CardNumber, 16) {
		return nil, fmt.Errorf("invalid card number")
	} else if !isValidNumber(in.Cvv, 3) {
		return nil, fmt.Errorf("invalid cvv")
	} else if !isCardValid(in.ExpireDate) {
		return nil, fmt.Errorf("invalid expiration date")
	}

	str, err := s.getTotalAmountByOrderId(in.OrderId)
	if err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	payment_id := uuid.New().String()
	created_at := time.Now()
	transaction_id := uuid.New().String()

	query, args, err := s.queryBuilder.Insert("payments").
		Columns(
			"payment_id",
			"order_id",
			"amount",
			"status",
			"payment_method",
			"transaction_id",
			"created_at").
		Values(
			payment_id,
			in.OrderId,
			str.total_amount,
			"pending",
			in.PaymentMethod,
			transaction_id,
			created_at).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	query, args, err = s.queryBuilder.Update("orders").
		Set("is_done", true).
		Where("order_id =?", in.OrderId).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.CreatePaymentResponse{
		PaymentId:     payment_id,
		OrderId:       in.OrderId,
		KitchenId:     str.kitchen_id,
		Amount:        str.total_amount,
		Status:        "pending",
		TransactionId: transaction_id,
		CreatedAt:     created_at.Format("2006-01-02 15:04:05"),
	}, nil
}
