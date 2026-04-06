package gateway

import (
	"github.com/timmyjinks/distributed-system/ratelimiter"
)

type Service struct {
	Gateway     *Gateway
	RateLimiter ratelimiter.RateLimiter
}

func NewService(g *Gateway, r ratelimiter.RateLimiter) Service {
	return Service{
		Gateway:     g,
		RateLimiter: r,
	}
}
