package sql_cm

import "time"

type TaskStruct struct {
	Id int
	Title string
	Description string
	Completed bool
	Created_at time.Time
	Completed_at *time.Time
}
