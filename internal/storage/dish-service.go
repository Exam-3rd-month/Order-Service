package storage

import (
	"context"
	"log"
	pb "order-service/genprotos/order_pb"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderSt) GetDishPriceById(ctx context.Context, dish_id string) (float64, error) {
	query, args, err := s.queryBuilder.Select("price").
		From("dishes").
		Where("dish_id =?", dish_id).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	var price float64
	err = row.Scan(&price)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	return price, nil
}

func (s *OrderSt) GetDishNameById(ctx context.Context, dishID string) (string, error) {
    query, args, err := s.queryBuilder.Select("name").
        From("dishes").
        Where("dish_id = ?", dishID).
        ToSql()
    if err != nil {
        return "", err
    }

    var name string
    err = s.db.QueryRowContext(ctx, query, args...).Scan(&name)
    if err != nil {
        return "", err
    }

    return name, nil
}

// 1
func (s *OrderSt) AddDish(ctx context.Context, in *pb.AddDishRequest) (*pb.AddDishResponse, error) {
	dish_id := uuid.New().String()
	created_at := time.Now()

	query, args, err := s.queryBuilder.Insert("dishes").
		Columns(
			"dish_id",
			"kitchen_id",
			"name",
			"description",
			"price",
			"category",
			"ingredients",
			"available",
			"created_at",
			"updated_at").
		Values(
			dish_id,
			in.KitchenId,
			in.Name,
			in.Description,
			in.Price,
			in.Category,
			pq.Array(in.Ingredients),
			in.Available,
			created_at,
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

	return &pb.AddDishResponse{
		DishId:      dish_id,
		KitchenId:   in.KitchenId,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Category:    in.Category,
		Ingredients: in.Ingredients,
		Available:   in.Available,
		CreatedAt:   created_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 2
func (s *OrderSt) UpdateDish(ctx context.Context, in *pb.UpdateDishRequest) (*pb.UpdateDishResponse, error) {
	updated_at := time.Now()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		s.logger.Error("Failed to begin transaction", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	defer tx.Rollback()

	query, args, err := s.queryBuilder.Update("dishes").
		Set("name", in.Name).
		Set("description", in.Description).
		Set("price", in.Price).
		Set("category", in.Category).
		Set("ingredients", pq.Array(in.Ingredients)).
		Set("available", in.Available).
		Set("updated_at", updated_at).
		Where("dish_id = ?", in.DishId).
		Where("kitchen_id = ?", in.KitchenId).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build update query", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute update query", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Failed to get rows affected", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "Dish not found or already deleted")
	}

	if err := tx.Commit(); err != nil {
		s.logger.Error("Failed to commit transaction", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.UpdateDishResponse{
		DishId:      in.DishId,
		KitchenId:   in.KitchenId,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Category:    in.Category,
		Ingredients: in.Ingredients,
		Available:   in.Available,
		UpdatedAt:   updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 3
func (s *OrderSt) DeleteDish(ctx context.Context, in *pb.DeleteDishRequest) (*pb.DeleteDishResponse, error) {
	query, args, err := s.queryBuilder.Update("dishes").
		Set("deleted_at", time.Now()).
		Where("dish_id =?", in.DishId).
		Where("kitchen_id =?", in.KitchenId).
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

	return &pb.DeleteDishResponse{
		Message: "Dish deleted",
	}, nil
}

// 4
func (s *OrderSt) ListDishes(ctx context.Context, in *pb.ListDishesRequest) (*pb.ListDishesResponse, error) {
	// Umumiy taomlar sonini hisoblash
	var total int32
	countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)").
		From("dishes").
		Where("kitchen_id = ?", in.KitchenId).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build count query", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	err = s.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		s.logger.Error("Failed to execute count query", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	// Paginatsiya parametrlarini hisoblash va tekshirish
	limit := in.Limit
	if limit <= 0 {
		limit = 10 // Default limit
	}

	totalPages := (total + limit - 1) / limit
	page := in.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		// Agar so'ralgan sahifa mavjud sahifalardan ko'p bo'lsa, oxirgi sahifani qaytaramiz
		page = totalPages
	}

	offset := (page - 1) * limit

	// Asosiy so'rov
	query, args, err := s.queryBuilder.Select("dish_id", "name", "price", "category", "available").
		From("dishes").
		Where("kitchen_id = ?", in.KitchenId).
		Where("deleted_at IS NULL").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	log.Println(query)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	defer rows.Close()

	dishes := []*pb.Dish{}
	for rows.Next() {
		dish := &pb.Dish{}
		err = rows.Scan(
			&dish.DishId,
			&dish.Name,
			&dish.Price,
			&dish.Category,
			&dish.Available,
		)
		if err != nil {
			s.logger.Error("Failed to scan row", "error", err)
			return nil, status.Error(codes.Internal, "Internal server error")
		}
		dishes = append(dishes, dish)
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error after scanning rows", "error", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.ListDishesResponse{
		Dishes: dishes,
		Total:  total,
		Page:   page,
		Limit:  limit,
	}, nil
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
