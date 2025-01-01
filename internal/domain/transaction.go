package domain

import (
	"sego/internal/entity"
	"sego/internal/repository"
)

type TransactionDomain struct {
	repository repository.Transaction
}

func NewTransactionDomain(repository repository.Transaction) *TransactionDomain {
	return &TransactionDomain{repository: repository}
}

func (t *TransactionDomain) CreateTransaction(transaction *entity.Transaction) error {
	return t.repository.Store(transaction)
}

func (t *TransactionDomain) GetTransactions() ([]entity.Transaction, error) {
	return t.repository.GetAll()
}
