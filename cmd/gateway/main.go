package main

import (
	"log"
	"time"

	"github.com/timmyjinks/distributed-system/config"
	"github.com/timmyjinks/distributed-system/gateway"
	"github.com/timmyjinks/distributed-system/monitoring"
	"github.com/timmyjinks/distributed-system/ratelimiter"
)

func main() {
	cfg := config.Load()

	log.Println(cfg)

	gate := gateway.NewGateway()
	ratelimiter := ratelimiter.NewSlidingWindowRateLimiter(5, time.Second)

	monitor := monitoring.NewPrometheusService("gateway_total_requests", "Total amount of requests recieved by gateway api")

	err := gate.AddHost("http://image:80", "/image")
	if err != nil {
		log.Fatal()
	}

	if err := gate.AddHost("http://report:80", "/report"); err != nil {
		log.Fatal()
	}

	if err := gate.AddHost("http://task:80", "/task"); err != nil {
		log.Fatal()
	}

	service := gateway.NewService(gate, ratelimiter, monitor)

	app := application{
		svc: service,
	}

	app.Run(cfg.Addr)
}
