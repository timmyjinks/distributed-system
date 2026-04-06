package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter interface {
	Allow() bool
}

type SlidingWindowRateLimiter struct {
	capacity   int
	window     time.Time
	windowSize time.Duration

	mu  sync.Mutex
	log []time.Time
}

func NewSlidingWindowRateLimiter(capacity int, windowSize time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		capacity:   capacity,
		windowSize: windowSize,
		log:        make([]time.Time, 0),
	}
}

func (r *SlidingWindowRateLimiter) Allow() bool {
	r.slideWindow()
	if len(r.log) > r.capacity {
		return false
	}

	return true
}

func (r *SlidingWindowRateLimiter) slideWindow() {
	now := time.Now()

	r.window = now.Add(-1 * r.windowSize)

	r.mu.Lock()
	for _, log := range r.log {
		if log.Before(r.window) {
			r.log = r.log[1:]
		}
	}
	r.mu.Unlock()

	r.mu.Lock()
	r.log = append(r.log, now)
	r.mu.Unlock()
}
