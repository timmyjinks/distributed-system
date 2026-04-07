package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timmyjinks/distributed-system/gateway"
)

type application struct {
	svc gateway.Service
}

func (app *application) Run(addr string) {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	handler := gateway.NewHandler(app.svc)

	app.svc.Monitor.Start(mux)
	mux.Handle("/image", handler.RateLimit(http.HandlerFunc(handler.Image)))
	mux.Handle("/report", handler.RateLimit(http.HandlerFunc(handler.Report)))
	mux.Handle("/task", handler.RateLimit(http.HandlerFunc(handler.Task)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	fmt.Println("Listening on localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}
