package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:el0404@localhost:5432/postgres")
	if err != nil {
		fmt.Println(err)
	}

	if err := conn.Ping(ctx); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected")
}