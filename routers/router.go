package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kevinfinalboss/ip-monitoring/handler"
	"github.com/kevinfinalboss/ip-monitoring/middleware"
	"github.com/kevinfinalboss/ip-monitoring/services"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60))
	r.Use(middleware.RequestLogger)

	s := &services.Service{}
	h := &handler.StatusHandler{
		Services: s,
	}
	r.Get("/status", h.GetStatus)

	return r
}
