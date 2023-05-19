package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kevinfinalboss/ip-monitoring/models"
)

type DiscordEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DiscordWebhookBody struct {
	Embeds []DiscordEmbed `json:"embeds"`
}

func PostToWebhook(webhookUrl string, status *models.IPStatus) error {
	embed := DiscordEmbed{
		Title: "Status de IP",
		Description: "Endereço IP: " + status.IPAddress +
			"\nStatus HTTP: " + strconv.Itoa(status.HTTPStatus) +
			"\nLatência: " + status.Latency +
			"\nRegistrar do Whois: " + status.WhoisRegistrar +
			"\nData de criação do Whois: " + status.WhoisCreationDate +
			"\nData de expiração do Whois: " + status.WhoisExpirationDate,
	}

	body := DiscordWebhookBody{
		Embeds: []DiscordEmbed{embed},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to send data to discord webhook, status code: %d", resp.StatusCode)
	}

	return nil
}
