package sql_cm

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetRows(ctx context.Context, conn *pgx.Conn) ([]TaskStruct, error) {
	sqlQuery  := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC	
	`

	rows, err := conn.Query(ctx,sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]TaskStruct, 0)

	for rows.Next() {
		var task TaskStruct 
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, err
}