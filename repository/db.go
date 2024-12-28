package repository

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/jackc/pgx/v5"
	"os"
)

var conn *pgx.Conn

func PgConnect() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/sego")
	if err != nil {
		fyne.LogError("Could not connect to database", err)
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			fyne.LogError("Could not close connection", err)
		}
	}(conn, context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 1").Scan(&greeting)
	if err != nil {
		fyne.LogError("Could not execute query", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
