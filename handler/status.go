package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kevinfinalboss/ip-monitoring/services"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}

	status, err := services.GetIPStatus(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
