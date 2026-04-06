package image

import (
	"context"
	"fmt"
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

	fmt.Println("production", string(b))
	if err := h.service.queue.Producer.Send(context.Background(), "image", queue.Message{
		ID:      id.String(),
		Type:    "image",
		Payload: b,
	}); err != nil {
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(id.String()))
}
