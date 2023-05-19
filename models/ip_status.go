package models

type IPStatus struct {
	URL                 string `json:"url"`
	IPAddress           string `json:"ip_address"`
	HTTPStatus          int    `json:"http_status"`
	LastDown            string `json:"last_down"`
	Latency             string `json:"latency"`
	WhoisRegistrar      string `json:"whois_registrar"`
	WhoisCreationDate   string `json:"whois_creation_date"`
	WhoisExpirationDate string `json:"whois_expiration_date"`
}
