package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kevinfinalboss/ip-monitoring/services"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	url := "https://github.com/"

	statusCode, err := services.GetIPStatus(url)
	if err != nil {
		log.Printf("Erro ao obter o status da URL: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"url":    url,
		"status": statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
