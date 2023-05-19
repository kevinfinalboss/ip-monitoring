package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kevinfinalboss/ip-monitoring/middleware"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestRequestLogger(t *testing.T) {
	_, hook := test.NewNullLogger()
	logrus.AddHook(hook)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	req, _ := http.NewRequest(http.MethodGet, "/foo", strings.NewReader(""))
	res := httptest.NewRecorder()

	handler := middleware.RequestLogger(next)
	handler.ServeHTTP(res, req)

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, "Processed a request", hook.LastEntry().Message)

	expectedFields := logrus.Fields{
		"method":     req.Method,
		"requestURI": req.RequestURI,
		"remoteAddr": req.RemoteAddr,
		"duration":   hook.LastEntry().Data["duration"],
	}
	assert.Equal(t, expectedFields, hook.LastEntry().Data)
}
