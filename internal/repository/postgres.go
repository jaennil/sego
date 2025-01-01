package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

var conn *pgx.Conn

func New() (*pgx.Conn, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err = pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/sego")
	if err != nil {
		return nil, err
	}

	return conn, conn.Ping(ctx)
}
