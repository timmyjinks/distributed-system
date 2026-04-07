package gateway

import (
	"github.com/timmyjinks/distributed-system/monitoring"
	"github.com/timmyjinks/distributed-system/ratelimiter"
)

type Service struct {
	Gateway     *Gateway
	RateLimiter ratelimiter.RateLimiter
	Monitor     *monitoring.PrometheusService
}

func NewService(g *Gateway, r ratelimiter.RateLimiter, m *monitoring.PrometheusService) Service {
	return Service{
		Gateway:     g,
		RateLimiter: r,
		Monitor:     m,
	}
}
