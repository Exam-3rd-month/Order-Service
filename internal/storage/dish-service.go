package storage

import (
	"context"
	"encoding/json"
	"log"
	pb "order-service/genprotos/order_pb"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		return nil, err
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
		return nil, err
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute update query", "error", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Failed to get rows affected", "error", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		s.logger.Error("Failed to commit transaction", "error", err)
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
	var total int32
	countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)").
		From("dishes").
		Where("kitchen_id = ?", in.KitchenId).
		Where("deleted_at IS NULL").
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

	query, args, err := s.queryBuilder.Select("dish_id", "name", "price", "category", "available").
		From("dishes").
		Where("kitchen_id = ?", in.KitchenId).
		Where("deleted_at IS NULL").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build query", "error", err)
		return nil, err
	}

	log.Println(query)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute query", "error", err)
		return nil, err
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
			return nil, err
		}
		dishes = append(dishes, dish)
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error after scanning rows", "error", err)
		return nil, err
	}

	return &pb.ListDishesResponse{
		Dishes: dishes,
		Total:  total,
		Page:   page,
		Limit:  limit,
	}, nil
}

// 2.5
func (s *OrderSt) UpdateDishNutritionInfo(ctx context.Context, in *pb.UpdateDishNutritionInfoRequest) (*pb.UpdateDishNutritionInfoResponse, error) {
	updated_at := time.Now()

	allergens := make([]string, len(in.Allergens))
	for i, a := range in.Allergens {
		allergens[i] = a.Name 
	}

	dietaryInfo := make([]string, len(in.DietaryInfo))
	for i, d := range in.DietaryInfo {
		dietaryInfo[i] = d.Name 
	}

	nutritionInfoJSON, err := json.Marshal(in.NutritionInfo)
	if err != nil {
		s.logger.Error("Failed to marshal nutrition info", "error", err)
		return nil, err
	}

	query, args, err := s.queryBuilder.
		Update("dishes").
		Set("allergens", pq.Array(allergens)).
		Set("nutrition_info", nutritionInfoJSON).
		Set("dietary_info", pq.Array(dietaryInfo)).
		Set("updated_at", updated_at).
		Suffix("RETURNING name").
		Where(sq.Eq{"dish_id": in.DishId}).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build update query", "error", err)
		return nil, err
	}

	var name string

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&name)
	if err != nil {
		s.logger.Error("Failed to execute update query", "error", err)
		return nil, err
	}

	return &pb.UpdateDishNutritionInfoResponse{
		DishId:        in.DishId,
		Name:          name,
		Allergens:     in.Allergens,
		NutritionInfo: in.NutritionInfo,
		DietaryInfo:   in.DietaryInfo,
		UpdatedAt: updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}
