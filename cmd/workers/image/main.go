package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/timmyjinks/distributed-system/config"
	"github.com/timmyjinks/distributed-system/queue"
	"github.com/timmyjinks/distributed-system/workers/image"
)

func main() {
	_ = config.Load()

	connStr := "host=db port=5432 user=postgres password=password sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	kafka := queue.NewKafkaService("image")

	svc := image.NewService(conn, kafka)

	app := application{
		svc: svc,
	}

	app.Run(":80")
}
