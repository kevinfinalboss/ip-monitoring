package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kevinfinalboss/ip-monitoring/handler"
	"github.com/kevinfinalboss/ip-monitoring/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60))

	r.Get("/status", handler.GetStatus)

	log.Fatal(http.ListenAndServe(":8080", r))
}
