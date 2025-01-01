package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"sego/internal/entity"
	"time"
)

type TransactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Store(t *entity.Transaction) error {
	query := `INSERT INTO transaction.transactions(timestamp, status, type, amount) VALUES ($1, $2, $3, $4)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := r.db.Exec(ctx, query, t.Date, t.Status, t.Typ, t.Amount)
	return err
}

func (r *TransactionRepository) GetAll() ([]entity.Transaction, error) {
	query := `SELECT timestamp, status, type, amount FROM transaction.transactions ORDER BY timestamp DESC`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		var t entity.Transaction
		err := rows.Scan(&t.Date, &t.Status, &t.Typ, &t.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return transactions, nil
}
