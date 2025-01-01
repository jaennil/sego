package repository

import (
	"github.com/jackc/pgx/v5"
	"sego/internal/entity"
)

type Repository struct {
	Transaction
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Transaction: NewTransactionRepository(db),
	}
}

type Transaction interface {
	Store(t *entity.Transaction) error
	GetAll() ([]entity.Transaction, error)
}
