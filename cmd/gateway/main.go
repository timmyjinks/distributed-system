package main

import (
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/timmyjinks/distributed-system/cache"
	"github.com/timmyjinks/distributed-system/config"
	"github.com/timmyjinks/distributed-system/gateway"
	"github.com/timmyjinks/distributed-system/monitoring"
	"github.com/timmyjinks/distributed-system/notifications/email"
	"github.com/timmyjinks/distributed-system/ratelimiter"
)

func main() {
	cfg := config.Load()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	gate := gateway.NewGateway()
	ratelimiter := ratelimiter.NewSlidingWindowRateLimiter(5, time.Second)
	email := email.NewClient(cfg.EmailConfig.APIKey)
	monitor := monitoring.NewPrometheusService("gateway_total_requests", "Total amount of requests recieved by gateway api")
	cache := cache.NewRedisServcie(rdb)

	err := gate.AddHost("http://image-service:80", "/image")
	if err != nil {
		log.Fatal()
	}

	if err := gate.AddHost("http://report-service:80", "/report"); err != nil {
		log.Fatal()
	}

	if err := gate.AddHost("http://task-service:80", "/task"); err != nil {
		log.Fatal()
	}

	service := gateway.NewService(gate, ratelimiter, monitor, email, cache)

	app := application{
		svc: service,
	}

	app.Run(cfg.Addr)
}
