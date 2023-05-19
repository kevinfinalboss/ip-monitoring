package services

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/kevinfinalboss/ip-monitoring/models"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func getUrlsFromFile() ([]string, error) {
	file, err := os.Open("urls.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func GetIPStatus(rawUrl string) (*models.IPStatus, error) {
	status := &models.IPStatus{}

	if !strings.HasPrefix(rawUrl, "http://") && !strings.HasPrefix(rawUrl, "https://") {
		rawUrl = "http://" + rawUrl
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("could not parse URL: %v", err)
	}

	ips, err := net.LookupIP(parsedUrl.Hostname())
	if err != nil {
		return nil, fmt.Errorf("could not get IPs: %v", err)
	}

	for _, ip := range ips {
		status.IPAddress = ip.String()
		break
	}

	resp, err := http.Get(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get HTTP status: %v", err)
	}
	status.HTTPStatus = resp.StatusCode

	start := time.Now()
	_, err = http.Get(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get latency: %v", err)
	}
	status.Latency = time.Since(start).String()

	whoisRecord, err := whois.Whois(parsedUrl.Hostname())
	if err != nil {
		return nil, fmt.Errorf("could not get whois record: %v", err)
	}
	parsedWhois, err := whoisparser.Parse(whoisRecord)
	if err != nil {
		return nil, fmt.Errorf("could not parse whois record: %v", err)
	}
	status.WhoisRegistrar = parsedWhois.Registrar.Name
	status.WhoisCreationDate = parsedWhois.Domain.CreatedDate
	status.WhoisExpirationDate = parsedWhois.Domain.ExpirationDate

	return status, nil
}
