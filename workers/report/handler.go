package report

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

func (h *Handler) Report(w http.ResponseWriter, r *http.Request) {
	var report Report
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.Job()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := h.service.Append(id, report.Title, report.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
}
