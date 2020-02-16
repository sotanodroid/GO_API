package models

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

//InitDB initialize db connection
func InitDB(dataSourceName string) {
	var err error
	ctx := context.Background()

	pool, err = pgxpool.Connect(ctx, dataSourceName)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

}
