package keysHandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	apikeys "github.com/Maxxxxxx-x/iris-swift/services/api_keys"
	"github.com/Maxxxxxx-x/iris-swift/services/auth"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)

func (handler *KeysHandler) CreateKey(ctx echo.Context) error {
	params := new(CreateApiKeyParam)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind params to context")
	}

	if err := ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	claims, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, "access-token-claims")
	if err != nil {
		if errors.Is(err, utils.ErrNoClaimsFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	tokenId, err := utils.GenerateULID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate API Key id")
	}

	apiKey, err := apikeys.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate API Key")
	}

	apiKeyHash, err := auth.HashPassword(apiKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate API Key")
	}

	queryParams := sqlc.CreateApiKeyParams{
		ID:             tokenId.String(),
		ApiKeyHash:     apiKeyHash,
		Name:           params.Name,
		AllowedDomains: params.AllowedDomains,
		CreatedBy:      claims.UserId,
	}

	if params.ExpiresAt != nil {
		duration, err := time.ParseDuration(*params.ExpiresAt)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse expires_at")
		}
		queryParams.ExpiresAt.Time = time.Now().Add(duration)
	}

	_, err = handler.querier.CreateApiKey(ctx.Request().Context(), queryParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new API Key record")
	}

	return ctx.JSON(http.StatusOK, convertToCreateApiKeyResponse(apiKey))
}
