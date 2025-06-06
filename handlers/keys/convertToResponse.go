package keysHandler

import (
	"database/sql"
	"time"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
)

func convertToApiKeyResponse(apiKey sqlc.ApiKey) APIKeyResponse {
	return APIKeyResponse{
		ID:            apiKey.ID,
		Name:          apiKey.Name,
		AlloweDomains: apiKey.AllowedDomains,
		UsageCount:    int(apiKey.UsageCount),
		ExpiresAt:     nullableTime(apiKey.ExpiresAt),
	}
}

func convertToApiUsageResponse(apiUsage sqlc.ApiUsage) APIKeyUsageResponse {
	return APIKeyUsageResponse{
		KeyId:     apiUsage.ApiKeyID,
		Timestamp: apiUsage.Timestamp,
		RequestIp: apiUsage.RequestIp,
		From:      apiUsage.FromAddr,
		To:        apiUsage.ToAddr,
		Subject:   apiUsage.Subject,
	}
}

func convertToCreateApiKeyResponse(apiKey string) CreateApiKeyResponse {
	return CreateApiKeyResponse{
		ApiKey: apiKey,
	}
}

func nullableTime(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}
