package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
    conn := os.Getenv("conn")
    return pgx.Connect(ctx, conn)
}