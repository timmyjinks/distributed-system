package gateway

import (
	"github.com/timmyjinks/distributed-system/cache"
	"github.com/timmyjinks/distributed-system/monitoring"
	"github.com/timmyjinks/distributed-system/notifications/email"
	"github.com/timmyjinks/distributed-system/ratelimiter"
)

type Service struct {
	Gateway     *Gateway
	RateLimiter ratelimiter.RateLimiter
	Monitor     *monitoring.PrometheusService
	Email       *email.ResendClient
	Cache       *cache.RedisService
}

func NewService(g *Gateway, r ratelimiter.RateLimiter, m *monitoring.PrometheusService, e *email.ResendClient, c *cache.RedisService) Service {
	return Service{
		Gateway:     g,
		RateLimiter: r,
		Monitor:     m,
		Email:       e,
		Cache:       c,
	}
}
