package main

import (
	"context"
	"fmt"
	"time"

	postgres "github.com/m4nsur/postgres-lrn/postgres/connection"
	"github.com/m4nsur/postgres-lrn/postgres/sql_cm"
)



func main() {
	ctx := context.Background()
	conn, err := postgres.CreateConnection(ctx)
	if err != nil {
		fmt.Println(err)
	}

	if err := sql_cm.CreateTable(ctx, conn); err !=nil {
		fmt.Println(err)
	} else {
		fmt.Println("created table")

	}

	taskValues := &sql_cm.TaskStruct{
	   Title:       "Купить книгу",
	   Description: "на озон",
	   Completed:   false,
	   Created_at:  time.Now(),
	}

}