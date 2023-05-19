package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kevinfinalboss/ip-monitoring/services"
)

type StatusHandler struct {
	Services services.IPStatusGetter
}

func (h *StatusHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}

	status, err := h.Services.GetIPStatus(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
