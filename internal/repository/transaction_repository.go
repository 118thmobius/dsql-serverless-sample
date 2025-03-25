package repository

import (
	"context"
	"github.com/118thmobius/dsql-serverless-sample/internal/domain"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, q db.Queryer, from domain.Account, to domain.Account, amount int) (domain.Transaction, error)
}
