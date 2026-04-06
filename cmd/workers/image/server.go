package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timmyjinks/distributed-system/workers/image"
)

type application struct {
	svc image.Service
}

func (app *application) Run(addr string) {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	handler := image.NewHandler(app.svc)
	mux.HandleFunc("/image", handler.Image)

	fmt.Println("Listening on localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}
