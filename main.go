package main

import (
	"context"
	"fmt"

	postgres "github.com/m4nsur/postgres-lrn/postgres/connection"
	"github.com/m4nsur/postgres-lrn/postgres/sql_cm"
)

// check shh dsds

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


}