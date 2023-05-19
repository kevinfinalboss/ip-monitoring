package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func Recoverer(next http.Handler) http.Handler {
	return middleware.Recoverer(next)
}

func Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

func Timeout(duration time.Duration) func(next http.Handler) http.Handler {
	return middleware.Timeout(duration)
}
