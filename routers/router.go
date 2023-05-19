package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kevinfinalboss/ip-monitoring/handler"
	"github.com/kevinfinalboss/ip-monitoring/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60))
	r.Use(middleware.RequestLogger)

	r.Get("/status", handler.GetStatus)

	return r
}
