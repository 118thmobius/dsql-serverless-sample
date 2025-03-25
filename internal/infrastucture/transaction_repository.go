package infrastructure

import (
	"context"
	"github.com/118thmobius/dsql-serverless-sample/internal/domain"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
	"github.com/118thmobius/dsql-serverless-sample/internal/repository"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() repository.TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (t TransactionRepositoryImpl) CreateTransaction(ctx context.Context, q db.Queryer, from domain.Account, to domain.Account, amount int) (domain.Transaction, error) {
	row := q.QueryRow(ctx, `
		INSERT INTO app_tx (from_id,to_id,amount)
		VALUES ($1,$2,$3) RETURNING tx_id
		`, from.UserId, to.UserId, amount)
	var tx domain.Transaction
	tx.From = from
	tx.To = to
	tx.Amount = amount
	err := row.Scan(
		&tx.Id,
	)
	if err != nil {
		return tx, err
	}
	return tx, err
}
