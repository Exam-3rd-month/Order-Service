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
	query, args, err := s.queryBuilder.Select("review_id", "user_id", "rating", "comment", "created_at").
		From("reviews").
		Where(sq.Eq{"kitchen_id": in.KitchenId}).
		OrderBy("created_at DESC").
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

	var reviews []*pb.Review

	for rows.Next() {
		review := &pb.Review{}
		err = rows.Scan(
			&review.ReviewId,
			&review.UserName,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return &pb.ListReviewsResponse{
		Reviews: reviews,
	}, nil
}
