package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/timmyjinks/distributed-system/config"
	"github.com/timmyjinks/distributed-system/queue"
	"github.com/timmyjinks/distributed-system/workers/image"
)

func main() {
	_ = config.Load()

	connStr := "host=db-service port=5432 user=postgres password=password sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	kafka := queue.NewKafkaService("image")

	ctx, cancel := context.WithCancel(context.Background())
	svc := image.NewService(ctx, conn, kafka)
	defer cancel()

	app := application{
		svc: svc,
	}

	app.Run(":80")
}
