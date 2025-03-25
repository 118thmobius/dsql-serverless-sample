package service

import "github.com/118thmobius/dsql-serverless-sample/internal/domain"

type TransactionDomainService interface {
	CanTransfer(account domain.StatefulAccount, amount int) bool
}

type TransactionDomainServiceImpl struct {
}

func NewTransactionDomainService() TransactionDomainService {
	return &TransactionDomainServiceImpl{}
}

func (s *TransactionDomainServiceImpl) CanTransfer(account domain.StatefulAccount, amount int) bool {
	return account.Balance >= amount
}
