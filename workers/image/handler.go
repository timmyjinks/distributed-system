package image

import (
	"context"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/timmyjinks/distributed-system/queue"
)

type Handler struct {
	service Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		service: svc,
	}
}

func (h *Handler) Image(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	id := uuid.New()

	go func() {
		if err := h.service.queue.Producer.Send(context.Background(), "image", queue.Message{
			ID:      id.String(),
			Type:    "image",
			Payload: b,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(id.String()))
}
