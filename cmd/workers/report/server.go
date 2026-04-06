package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timmyjinks/distributed-system/workers/report"
)

type application struct {
	svc report.Service
}

func (app *application) Run(addr string) {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	handler := report.NewHandler(app.svc)
	mux.HandleFunc("/report", handler.Report)

	fmt.Println("Listening on localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}
