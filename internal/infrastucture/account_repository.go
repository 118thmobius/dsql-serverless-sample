package infrastructure

import (
	"context"
	"errors"
	"github.com/118thmobius/dsql-serverless-sample/internal/domain"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
	"github.com/118thmobius/dsql-serverless-sample/internal/repository"
)

type AccountRepositoryImpl struct {
}

func NewAccountRepository() repository.AccountRepository {
	return &AccountRepositoryImpl{}
}

func (r AccountRepositoryImpl) GetAccountByID(ctx context.Context, q db.Queryer, userId string) (domain.StatefulAccount, error) {
	var account domain.StatefulAccount
	row := q.QueryRow(ctx, `
		SELECT user_id, screen_name, balance
		FROM app_account
		WHERE user_id = $1
		`, userId)
	err := row.Scan(
		&account.Account.UserId,
		&account.Account.ScreenName,
		&account.Balance,
	)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (r AccountRepositoryImpl) UpdateAccount(ctx context.Context, q db.Queryer, userId string, balance int) error {
	cmdTag, err := q.Exec(ctx, `
		UPDATE app_account 
		SET balance = $1 
		WHERE user_id = $2
		`, balance, userId)
	if err != nil {
		return err
	}
	if affected := cmdTag.RowsAffected(); affected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
