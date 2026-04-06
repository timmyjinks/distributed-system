package gateway

import (
	"fmt"
	"net/http"
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
			next.ServeHTTP(w, r)
		} else {
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
