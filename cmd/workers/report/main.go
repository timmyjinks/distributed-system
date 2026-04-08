package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/timmyjinks/distributed-system/config"
	"github.com/timmyjinks/distributed-system/workers/report"
)

func main() {
	_ = config.Load()

	connStr := "host=db port=5432 user=postgres password=password sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	svc := report.NewService(conn)

	app := application{
		svc: svc,
	}

	app.Run(":80")
}
