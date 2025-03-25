package repository

import (
	"context"
	"github.com/118thmobius/dsql-serverless-sample/internal/domain"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
)

type AccountRepository interface {
	GetAccountByID(ctx context.Context, q db.Queryer, userId string) (domain.StatefulAccount, error)
	UpdateAccount(ctx context.Context, q db.Queryer, userId string, balance int) error
}
