package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"time"

	pb "order-service/genprotos/order_pb"
	"order-service/internal/config"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type (
	OrderSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
		logger       *slog.Logger
	}
)

func New(config *config.Config, logger *slog.Logger) (*OrderSt, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &OrderSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		logger:       logger,
	}, nil
}

// 5
func (s *OrderSt) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order_id := uuid.New().String()
	created_at := time.Now()
	var total_amount float64

	var orderItems []*pb.OrderResponse
	for _, item := range in.Items {
		price, err := s.GetDishPriceById(ctx, item.DishId)
		if err != nil {
			s.logger.Error("Failed to get dish price", "error", err)
			return nil, err
		}

		dishName, err := s.GetDishNameById(ctx, item.DishId)
		if err != nil {
			s.logger.Error("Failed to get dish name", "error", err)
			return nil, err
		}

		orderItems = append(orderItems, &pb.OrderResponse{
			DishId:   item.DishId,
			Name:     dishName,
			Price:    price,
			Quantity: fmt.Sprintf("%d", item.Quantity),
		})

		total_amount += price * float64(item.Quantity)
	}

	itemsJSON, err := json.Marshal(in.Items)
	if err != nil {
		s.logger.Error("Failed to marshal items", "error", err)
		return nil, err
	}

	query, args, err := s.queryBuilder.Insert("orders").
		Columns(
			"order_id",
			"user_id",
			"kitchen_id",
			"items",
			"total_amount",
			"status",
			"delivery_address",
			"delivery_time",
			"created_at").
		Values(
			order_id,
			in.UserId,
			in.KitchenId,
			itemsJSON,
			total_amount,
			"pending",
			in.DeliveryAddress,
			in.DeliveryTime,
			created_at).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build insert query", "error", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute insert query", "error", err)
		return nil, err
	}

	return &pb.CreateOrderResponse{
		OrderId:         order_id,
		UserId:          in.UserId,
		KitchenId:       in.KitchenId,
		Items:           orderItems,
		TotalAmount:     total_amount,
		Status:          "pending",
		DeliveryAddress: in.DeliveryAddress,
		DeliveryTime:    in.DeliveryTime,
		CreatedAt:       created_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 6
func (s *OrderSt) UpdateOrderStatus(ctx context.Context, in *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	updated_at := time.Now()

	query, args, err := s.queryBuilder.Update("orders").
		Set("status", in.Status).
		Set("updated_at", updated_at).
		Where(sq.Eq{"order_id": in.OrderId}).
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

	return &pb.UpdateOrderStatusResponse{
		OrderId:   in.OrderId,
		Status:    in.Status,
		UpdatedAt: updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 7
func (s *OrderSt) ListOfOrders(ctx context.Context, in *pb.ListOfOrdersRequest) (*pb.ListOfOrdersResponse, error) {
	var total int32
	countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)").
		From("orders").
		Where("user_id = ?", in.UserId).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build count query", "error", err)
		return nil, err
	}

	err = s.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		s.logger.Error("Failed to execute count query", "error", err)
		return nil, err
	}

	limit := in.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := (total + limit - 1) / limit
	page := in.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	query, args, err := s.queryBuilder.Select("order_id", "kitchen_id", "total_amount", "status", "delivery_address", "created_at").
		From("orders").
		Where("user_id = ?", in.UserId).
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}
	defer rows.Close()

	orders := []*pb.Order{}
	for rows.Next() {
		order := &pb.Order{}
		var createdAt time.Time
		err = rows.Scan(
			&order.OrderId,
			&order.KitchenId,
			&order.TotalAmount,
			&order.Status,
			&order.DeliveryAddress,
			&createdAt,
		)
		if err != nil {
			s.logger.Error("Failed to scan row", "error", err)
			return nil, err
		}
		order.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error after scanning rows", "error", err)
		return nil, err
	}

	return &pb.ListOfOrdersResponse{
		Orders: orders,
		Total:  total,
		Page:   page,
		Limit:  limit,
	}, nil
}

// 8
func (s *OrderSt) GetOrderByKitchenId(ctx context.Context, in *pb.GetOrderByKitchenIdRequest) (*pb.GetOrderByKitchenIdResponse, error) {
	var total int32
	countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)").
		From("orders").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build count query", "error", err)
		return nil, err
	}

	err = s.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		s.logger.Error("Failed to execute count query", "error", err)
		return nil, err
	}

	limit := in.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := (total + limit - 1) / limit
	page := in.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	query, args, err := s.queryBuilder.Select("order_id", "kitchen_id", "total_amount", "status", "delivery_address").
		From("orders").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Order

	for rows.Next() {
		order := &pb.Order{}
		err = rows.Scan(
			&order.OrderId,
			&order.KitchenId,
			&order.TotalAmount,
			&order.Status,
			&order.DeliveryAddress,
		)
		if err != nil {
			s.logger.Error("Failed to scan row", "error", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error after scanning rows", "error", err)
		return nil, err
	}

	return &pb.GetOrderByKitchenIdResponse{
		Orders: orders,
		Total:  total,
		Page:   page,
		Limit:  limit,
	}, nil
}

// 12
func (s *OrderSt) GetFullInfoAboutOrder(ctx context.Context, in *pb.GetFullInfoAboutOrderRequest) (*pb.GetFullInfoAboutOrderResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"user_id",
		"kitchen_id",
		"status",
		"delivery_address",
		"delivery_time",
		"created_at",
		"updated_at").
		From("orders").
		Where(sq.Eq{"order_id": in.OrderId}).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	log.Println(query)
	var order pb.GetFullInfoAboutOrderResponse
	var deliveryTime, createdAt, updatedAt time.Time

	err = s.db.QueryRowContext(ctx, query, args...).Scan(
		&order.UserId,
		&order.KitchenId,
		&order.Status,
		&order.DiliveryAddress,
		&deliveryTime,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		s.logger.Error("Failed to scan row", "error", err)
		return nil, err
	}

	order.OrderId = in.OrderId
	order.DeliveryTime = deliveryTime.Format(time.RFC3339)
	order.CreatedAt = createdAt.Format(time.RFC3339)
	order.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &order, nil
}

/*
-- Dishes jadvali
CREATE TABLE IF NOT EXISTS dishes (
    dish_id UUID PRIMARY KEY,
    kitchen_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(50),
    ingredients TEXT[],
    allergens TEXT[],
    nutrition_info JSONB,
    dietary_info TEXT[],
    available BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Orders jadvali
CREATE TABLE IF NOT EXISTS orders (
    order_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    kitchen_id UUID NOT NULL,
    items JSONB NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL,
    delivery_address TEXT NOT NULL,
    delivery_time TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Reviews jadvali
CREATE TABLE IF NOT EXISTS reviews (
    review_id UUID PRIMARY KEY,
    order_id UUID REFERENCES orders(order_id),
    user_id UUID NOT NULL,
    kitchen_id UUID NOT NULL,
    rating DECIMAL(2, 1) NOT NULL,
    comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Payments jadvali
CREATE TABLE IF NOT EXISTS payments (
    payment_id UUID PRIMARY KEY,
    order_id UUID REFERENCES orders(order_id),
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    transaction_id VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Working_Hours jadvali
CREATE TABLE IF NOT EXISTS working_hours (
    kitchen_id UUID NOT NULL,
    day_of_week INTEGER NOT NULL,
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    PRIMARY KEY (kitchen_id, day_of_week)
);

-- User_Preferences jadvali
CREATE TABLE IF NOT EXISTS user_preferences (
    user_id UUID NOT NULL,
    cuisine_type VARCHAR(50),
    dietary_preferences TEXT[],
    favorite_kitchen_ids UUID[],
    PRIMARY KEY (user_id)
);

-- Delivery_Routes jadvali
CREATE TABLE IF NOT EXISTS delivery_routes (
    delivery_id UUID PRIMARY KEY,
    order_id UUID REFERENCES orders(order_id),
    start_address TEXT NOT NULL,
    end_address TEXT NOT NULL,
    distance DECIMAL(10, 2) NOT NULL,
    duration INTEGER NOT NULL,
    route_polyline TEXT,
    waypoints JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

*/
