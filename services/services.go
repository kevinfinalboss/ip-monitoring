package services

import (
	"fmt"
	"net/http"
	"time"
)

func GetIPStatus(url string) (int, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("erro ao realizar a requisição GET para %s: %w", url, err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
