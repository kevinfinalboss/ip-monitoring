package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevinfinalboss/ip-monitoring/handler"
	"github.com/kevinfinalboss/ip-monitoring/models"
	"github.com/stretchr/testify/assert"
)

type MockIPStatusGetter struct{}

func (m *MockIPStatusGetter) GetIPStatus(rawUrl string) (*models.IPStatus, error) {
	return &models.IPStatus{
		IPAddress:           "127.0.0.1",
		HTTPStatus:          200,
		Latency:             "0ms",
		WhoisRegistrar:      "ICANN",
		WhoisCreationDate:   "01-01-1970",
		WhoisExpirationDate: "01-01-2070",
	}, nil
}

func TestGetStatus(t *testing.T) {
	h := &handler.StatusHandler{
		Services: &MockIPStatusGetter{},
	}

	req, err := http.NewRequest("GET", "/status?url=http://test.com", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetStatus)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var res models.IPStatus
	err = json.NewDecoder(bytes.NewReader(recorder.Body.Bytes())).Decode(&res)
	assert.NoError(t, err)

	assert.Equal(t, "127.0.0.1", res.IPAddress)
	assert.Equal(t, 200, res.HTTPStatus)
	assert.Equal(t, "0ms", res.Latency)
	assert.Equal(t, "ICANN", res.WhoisRegistrar)
	assert.Equal(t, "01-01-1970", res.WhoisCreationDate)
	assert.Equal(t, "01-01-2070", res.WhoisExpirationDate)
}
