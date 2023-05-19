package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(startTime)
		logger := logrus.WithFields(logrus.Fields{
			"method":     r.Method,
			"requestURI": r.RequestURI,
			"remoteAddr": r.RemoteAddr,
			"duration":   duration,
		})
		logger.Info("Processed a request")
	})
}
