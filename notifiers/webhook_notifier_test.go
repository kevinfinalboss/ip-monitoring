package notifiers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevinfinalboss/ip-monitoring/models"
)

func TestPostToWebhook(t *testing.T) {
	status := &models.IPStatus{
		IPAddress:           "1.1.1.1",
		HTTPStatus:          200,
		Latency:             "1ms",
		WhoisRegistrar:      "Registrar Inc.",
		WhoisCreationDate:   "2020-01-01",
		WhoisExpirationDate: "2022-01-01",
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)

		var received DiscordWebhookBody
		err := decoder.Decode(&received)
		if err != nil {
			t.Fatalf("Could not decode received body: %v", err)
		}

		if len(received.Embeds) != 1 {
			t.Fatalf("Expected one embed, got %d", len(received.Embeds))
		}

		if received.Embeds[0].Title != "Status de IP" {
			t.Fatalf("Expected embed title to be 'Status de IP', got '%s'", received.Embeds[0].Title)
		}
	}))
	defer server.Close()

	err := PostToWebhook(server.URL, status)
	if err != nil {
		t.Fatalf("PostToWebhook returned an error: %v", err)
	}
}
