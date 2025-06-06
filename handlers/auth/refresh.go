package authHandler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/middleware"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)


func (handler *AuthHandler) RegisterRefreshAPIRoute(api *echo.Group) {
	api.POST("/refresh", handler.HandleRefresh, middleware.RefreshTokenRequired)
}


func (handler *AuthHandler) HandleRefresh(ctx echo.Context) error {
	claim, err := utils.GetClaimsFromContext[token.RefreshTokenClaims](ctx, "refresh-token-claims")
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "Invalid or missing refresh token")
	}
	userId := claim.UserId
	user, err := handler.querier.GetUserById(ctx.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.removeAccessTokenFromCookie(ctx)
			handler.removeRefreshTokenFromCookie(ctx)
			return ctx.JSON(http.StatusUnauthorized, "Unale to find account")
		}
		return ctx.String(http.StatusInternalServerError, "an error occured while finding your account")
	}
	accessToken, err := token.GenerateAccessToken(userId, user.AccountType)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Error occured generating access token")
	}
	handler.setAccessTokenToCookie(ctx, accessToken)
	return ctx.String(http.StatusOK, "refreshed")
}
