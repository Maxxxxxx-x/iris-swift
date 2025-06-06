package keysHandler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)

func (handler *KeysHandler) GetKeysFromUser(ctx echo.Context) error {
	claims, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, "access-token-claims")
	if err != nil {
		if errors.Is(err, utils.ErrNoClaimsFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	keys, err := handler.querier.GetAllApiKeysFromUserId(ctx.Request().Context(), claims.UserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occured fetching records")
	}

	var response []APIKeyResponse
	for _, key := range keys {
		response = append(response, convertToApiKeyResponse(key))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (handler *KeysHandler) GetKeyById(ctx echo.Context) error {
	id := ctx.Param("id")

	claims, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, "access-token-claims")
	if err != nil {
		if errors.Is(err, utils.ErrNoClaimsFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	key, err := handler.querier.GetApiKeyFromUserByID(ctx.Request().Context(), sqlc.GetApiKeyFromUserByIDParams{
		ID:        id,
		CreatedBy: claims.UserId,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occured fetching records")
	}

	return ctx.JSON(http.StatusOK, convertToApiKeyResponse(key))
}

func (handler *KeysHandler) GetKeyUsageById(ctx echo.Context) error {
	id := ctx.Param("id")

	claims, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, "access-token-claims")
	if err != nil {
		if errors.Is(err, utils.ErrNoClaimsFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	usage, err := handler.querier.GetApiKeyUsageByKeyIdAndUserId(ctx.Request().Context(), sqlc.GetApiKeyUsageByKeyIdAndUserIdParams{
		ApiKeyID:  id,
		CreatedBy: claims.UserId,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occured fetching usage logs")
	}

	var usageRes []APIKeyUsageResponse
	for _, use := range usage {
		usageRes = append(usageRes, convertToApiUsageResponse(use))
	}

	return ctx.JSON(http.StatusOK, usageRes)
}
