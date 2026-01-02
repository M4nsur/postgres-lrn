package sql_cm

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type TaskStruct struct {
	Title string
	Description string
	Completed bool
	Created_at time.Time
}


func InsertRow(ctx context.Context, conn *pgx.Conn, task *TaskStruct) error {
	sqlQuery  := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES($1, $2, $3, $4);
	`

	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.Created_at)


	return err
}