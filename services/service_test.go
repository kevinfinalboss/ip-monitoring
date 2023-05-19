package services

import (
	"os"
	"testing"

	"github.com/kevinfinalboss/ip-monitoring/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUrlsFromFile(t *testing.T) {
	fileName := "test_urls.txt"
	defer os.Remove(fileName)

	file, _ := os.Create(fileName)
	file.WriteString("https://google.com\n")
	file.WriteString("https://facebook.com\n")
	file.Close()

	urls, err := GetUrlsFromFile(fileName)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(urls))
	assert.Equal(t, "https://google.com", urls[0])
	assert.Equal(t, "https://facebook.com", urls[1])
}

func TestGetIPStatus(t *testing.T) {
	rawUrl := "https://google.com"
	service := Service{}
	status, err := service.GetIPStatus(rawUrl)

	assert.Nil(t, err)
	assert.IsType(t, &models.IPStatus{}, status)

	assert.NotEmpty(t, status.IPAddress)
	assert.Equal(t, 200, status.HTTPStatus)
	assert.NotEmpty(t, status.Latency)

	assert.Equal(t, "MarkMonitor Inc.", status.WhoisRegistrar)
	assert.NotEmpty(t, status.WhoisCreationDate)
	assert.NotEmpty(t, status.WhoisExpirationDate)
}
