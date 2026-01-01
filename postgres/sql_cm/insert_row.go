package sql_cm

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery  := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES('Купить что-то', 'завтра в магазине', FALSE, '2026-01-02 01:02');
	`

	_, err := conn.Exec(ctx, sqlQuery)


	return err
}