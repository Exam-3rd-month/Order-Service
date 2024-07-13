package storage

import (
	"context"
	pb "order-service/genprotos/order_pb"
	"time"

	"github.com/google/uuid"
)

func (s *OrderSt) GetDishPriceById(ctx context.Context, dish_id string) (float64, error) {
	query, args, err := s.queryBuilder.Select("price").
		From("dishes").
		Where("dish_id =?", dish_id).
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
			"updates_at").
		Values(
			dish_id,
			in.KitchenId,
			in.Name,
			in.Description,
			in.Price,
			in.Category,
			in.Ingredients,
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

	query, args, err := s.queryBuilder.Update("dishes").
		Set("name", in.Name).
		Set("description", in.Description).
		Set("price", in.Price).
		Set("category", in.Category).
		Set("ingredients", in.Ingredients).
		Set("available", in.Available).
		Set("updated_at", updated_at).
		Where("dish_id =?", in.DishId).
		Where("kitchen_id =?", in.KitchenId).
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
	query, args, err := s.queryBuilder.Delete("dishes").
		Where("dish_id =?", in.DishId).
		Where("kitchen_id =?", in.KitchenId).
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
	query, args, err := s.queryBuilder.Select("dish_id", "name", "description", "price", "category", "ingredients", "available").
		From("dishes").
		Where("kitchen_id =?", in.KitchenId).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	dishes := []*pb.Dish{}
	for rows.Next() {
		dish := &pb.Dish{}
		err = rows.Scan(&dish.DishId, &dish.Name, &dish.Description, &dish.Price, &dish.Category, &dish.Ingredients, &dish.Available)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}
		dishes = append(dishes, dish)
	}

	return &pb.ListDishesResponse{
		Dishes: dishes,
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
