package keysHandler

import "time"

type CreateApiKeyParam struct {
	Name           string  `json:"name" validate:"required"`
	AllowedDomains string  `json:"allowed_domains" validate:"required"`
	ExpiresAt      *string `json:"expires_at"`
}

type CreateApiKeyResponse struct {
	ApiKey string `json:"api_key"`
}

type APIKeyResponse struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	AlloweDomains string     `json:"allowed_domains"`
	UsageCount    int        `json:"usage_count"`
	ExpiresAt     *time.Time `json:"expires_at"`
}

type APIKeyUsageResponse struct {
	KeyId     string    `json:"key_id"`
	Timestamp time.Time `json:"timestamp"`
	RequestIp string    `json:"request_ip"`
	From      string    `json:"From"`
	To        string    `json:"To"`
	Subject   string    `json:"Subject"`
}
