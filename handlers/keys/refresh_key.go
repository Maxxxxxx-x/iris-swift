package keysHandler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	apikeys "github.com/Maxxxxxx-x/iris-swift/services/api_keys"
	"github.com/Maxxxxxx-x/iris-swift/services/auth"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)

func (handler *KeysHandler) RefreshKeyWithId(ctx echo.Context) error {
	id := ctx.Param("id")

	claims, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, "access-token-claims")
	if err != nil {
		if errors.Is(err, utils.ErrNoClaimsFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	_, err = handler.querier.GetApiKeyFromUserByID(ctx.Request().Context(), sqlc.GetApiKeyFromUserByIDParams{
		ID:        id,
		CreatedBy: claims.UserId,
	})

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occured fetching records")
	}
	apiKey, err := apikeys.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate API Key")
	}

	apiKeyHash, err := auth.HashPassword(apiKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate API Key")
	}

	if err := handler.querier.RefreshApiKey(ctx.Request().Context(), sqlc.RefreshApiKeyParams{
		ID:         id,
		ApiKeyHash: apiKeyHash,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to refresh api key")
	}

	return nil
}
