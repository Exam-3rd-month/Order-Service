package storage

import (
	"context"
	"time"

	pb "order-service/genprotos/order_pb"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// 9
func (s *OrderSt) AddReview(ctx context.Context, in *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {
	rewiev_id := uuid.New().String()
	created_at := time.Now()

	query, args, err := s.queryBuilder.Insert("reviews").
		Columns(
			"review_id",
			"order_id",
			"user_id",
			"kitchen_id",
			"rating",
			"comment",
			"created_at").
		Values(
			rewiev_id,
			in.OrderId,
			in.UserId,
			in.KitchenId,
			in.Rating,
			in.Comment,
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

	return &pb.AddReviewResponse{
		RewievId:  rewiev_id,
		OrderId:   in.OrderId,
		UserId:    in.UserId,
		KitchenId: in.KitchenId,
		Rating:    in.Rating,
		Comment:   in.Comment,
		CreatedAt: created_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 10
func (s *OrderSt) ListReviews(ctx context.Context, in *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
    var total int32
    var averageRating float32
    countQuery, countArgs, err := s.queryBuilder.Select("COUNT(*)", "AVG(rating)").
        From("reviews").
        Where(sq.Eq{"kitchen_id": in.KitchenId}).
        ToSql()
    if err != nil {
        s.logger.Error("Failed to build count query", "error", err)
        return nil, err
    }

    err = s.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total, &averageRating)
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

    query, args, err := s.queryBuilder.Select("review_id", "user_id", "rating", "comment", "created_at").
        From("reviews").
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

    var reviews []*pb.Review

    for rows.Next() {
        review := &pb.Review{}
        var createdAt time.Time
        err = rows.Scan(
            &review.ReviewId,
            &review.UserName,
            &review.Rating,
            &review.Comment,
            &createdAt,
        )
        if err != nil {
            s.logger.Error("Failed to scan row", "error", err)
            return nil, err
        }
        review.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
        reviews = append(reviews, review)
    }

    if err = rows.Err(); err != nil {
        s.logger.Error("Error after scanning rows", "error", err)
        return nil, err
    }

    return &pb.ListReviewsResponse{
        Reviews:       reviews,
        Total:         total,
        AverageRating: averageRating,
        Page:          page,
        Limit:         limit,
    }, nil
}

