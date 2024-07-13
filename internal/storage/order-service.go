package storage

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"order-service/internal/config"
)

type (
	OrderSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*OrderSt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &OrderSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}
