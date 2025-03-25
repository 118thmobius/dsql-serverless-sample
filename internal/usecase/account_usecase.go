package usecase

import (
	"context"
	"errors"
	"github.com/118thmobius/dsql-serverless-sample/internal/domain"
	infrastructure "github.com/118thmobius/dsql-serverless-sample/internal/infrastucture"
	"github.com/118thmobius/dsql-serverless-sample/internal/infrastucture/db"
	"github.com/118thmobius/dsql-serverless-sample/internal/repository"
	"github.com/118thmobius/dsql-serverless-sample/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountUseCase interface {
	GetAccountByID(ctx context.Context, userId string) (domain.StatefulAccount, error)
	Transfer(ctx context.Context, fromId, toId string, amount int) (domain.Transaction, error)
}

type AccountUseCaseImpl struct {
	pool                  *pgxpool.Pool
	txManager             db.TransactionManager
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
	txDomainService       service.TransactionDomainService
}

func NewAccountUseCase(
	txManager db.TransactionManager,
) *AccountUseCaseImpl {
	return &AccountUseCaseImpl{
		txManager:             txManager,
		accountRepository:     infrastructure.NewAccountRepository(),
		transactionRepository: infrastructure.NewTransactionRepository(),
		txDomainService:       service.NewTransactionDomainService(),
	}
}

func (u *AccountUseCaseImpl) GetAccountByID(ctx context.Context, userId string) (domain.StatefulAccount, error) {
	var account domain.StatefulAccount
	err := u.txManager.Do(ctx, func(ctx context.Context, q db.Queryer) error {
		acc, err := u.accountRepository.GetAccountByID(ctx, q, userId)
		if err != nil {
			return err
		}
		account = acc
		return nil
	})
	if err != nil {
		return account, err
	}
	return account, nil
}

func (u *AccountUseCaseImpl) Transfer(ctx context.Context, fromId, toId string, amount int) (domain.Transaction, error) {
	var tx domain.Transaction
	err := u.txManager.DoTx(ctx, func(ctx context.Context, q db.Queryer) error {
		fromAccount, err := u.accountRepository.GetAccountByID(ctx, q, fromId)
		if err != nil {
			return err
		}

		toAccount, err := u.accountRepository.GetAccountByID(ctx, q, toId)
		if err != nil {
			return err
		}

		if u.txDomainService.CanTransfer(fromAccount, amount) {
			fromAccount.Balance -= amount
			toAccount.Balance += amount

			if e := u.accountRepository.UpdateAccount(ctx, q, fromAccount.UserId, fromAccount.Balance); e != nil {
				return e
			}
			if e := u.accountRepository.UpdateAccount(ctx, q, toAccount.UserId, toAccount.Balance); e != nil {
				return e
			}
			t, e := u.transactionRepository.CreateTransaction(ctx, q, fromAccount.Account, toAccount.Account, amount)
			if e != nil {
				return e
			}
			tx = t
			return nil
		}
		return errors.New("Insufficient funds")
	})
	return tx, err
}
