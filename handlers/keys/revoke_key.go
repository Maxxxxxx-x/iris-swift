package keysHandler

import (
	"errors"
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)

func (handler *KeysHandler) RevokeKeyWithId(ctx echo.Context) error {
	id := ctx.Param("id")

	claims, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, "access-token-claims")
	if err != nil {
		if errors.Is(err, utils.ErrNoClaimsFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := handler.querier.DeleteApiKeyByIdFromUser(ctx.Request().Context(), sqlc.DeleteApiKeyByIdFromUserParams{
		ID:        id,
		CreatedBy: claims.UserId,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to revoke token. Try again later")
	}

	return nil
}
