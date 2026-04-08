package gateway

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h *Handler) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.svc.RateLimiter.Allow() {
			h.svc.Monitor.Inc()
			next.ServeHTTP(w, r)
		} else {
			t, err := h.svc.Cache.GetUser(context.Background(), "something")
			if err != nil {
				log.Println("no good")
			}

			empty := time.Time{}

			if t == empty {
				h.svc.Email.SendEmail("error panic too many requests")
				h.svc.Cache.SetUser(context.Background(), "something", time.Now())
			}

			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	})
}

func (h *Handler) Image(w http.ResponseWriter, r *http.Request) {
	job, _ := h.svc.Gateway.hosts["/image"]
	fmt.Println(h.svc.Gateway.hosts)
	job.ServeHTTP(w, r)
}

func (h *Handler) Report(w http.ResponseWriter, r *http.Request) {
	job, _ := h.svc.Gateway.hosts["/report"]
	job.ServeHTTP(w, r)
}

func (h *Handler) Task(w http.ResponseWriter, r *http.Request) {
	job, _ := h.svc.Gateway.hosts["/task"]
	job.ServeHTTP(w, r)
}
