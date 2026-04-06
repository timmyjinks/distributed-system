package task

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		service: svc,
	}
}

func (h *Handler) Task(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.Job()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := h.service.Append(id, task.Type); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
}
