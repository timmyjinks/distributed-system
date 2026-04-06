package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timmyjinks/distributed-system/workers/task"
)

type application struct {
	svc task.Service
}

func (app *application) Run(addr string) {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	handler := task.NewHandler(app.svc)
	mux.HandleFunc("/task", handler.Task)

	fmt.Println("Listening on localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}
