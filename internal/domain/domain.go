package domain

import (
	"sego/internal/entity"
	"sego/internal/repository"
)

type Domain struct {
	Transaction
}

func NewDomain(repositories *repository.Repository) *Domain {
	return &Domain{
		Transaction: NewTransactionDomain(repositories.Transaction),
	}
}

type Transaction interface {
	CreateTransaction(t *entity.Transaction) error
	GetTransactions() ([]entity.Transaction, error)
}
