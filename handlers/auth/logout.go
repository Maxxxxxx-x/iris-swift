package authHandler

import (
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/middleware"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)

func (handler *AuthHandler) RegisterLogoutAPIRoute(api *echo.Group) {
	api.POST("/logout", handler.HandleLogout, middleware.AccessTokenRequired)
}

func (handler *AuthHandler) HandleLogout(ctx echo.Context) error {
	_, err := utils.GetClaimsFromContext[token.AccessTokenClaims](ctx, ACCESS_TOKEN_NAME)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "Invalid or missing access token")
	}
	handler.removeAccessTokenFromCookie(ctx)
	handler.removeRefreshTokenFromCookie(ctx)
	return ctx.String(http.StatusOK, "bye bye")
}
