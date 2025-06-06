package authHandler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/handlers"
	"github.com/Maxxxxxx-x/iris-swift/services/auth"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/Maxxxxxx-x/iris-swift/views/pages"
	"github.com/labstack/echo/v4"
)

func (handler *AuthHandler) RegisterLoginUIRoutes(e *echo.Echo) {
	e.GET("/login", handler.ShowLoginPage)

	handler.logger.Info().Msg("Registered Login UI routes")
}

func (handler *AuthHandler) RegisterLoginAPIRoutes(api *echo.Group) {
	api.POST("/login", handler.HandleLogin)
	handler.logger.Info().Msg("Registered Login API routes")
}

func (handler *AuthHandler) ShowLoginPage(ctx echo.Context) error {
	component := pages.LoginPage(pages.LoginParams{})
	handlers.RenderTempl(ctx, &component)
	return nil
}

func (handler *AuthHandler) HandleLogin(ctx echo.Context) error {
	params := new(pages.LoginParams)

	if err := ctx.Bind(params); err != nil {
		handler.logger.Error().Err(err).Msg("Error occured")
		return ctx.String(http.StatusInternalServerError, "Failed to bind to params")
	}

	if err := ctx.Validate(params); err != nil {
		handler.logger.Warn().Msg("Bad Request")
		return err
	}

	user, err := handler.querier.GetUserByEmail(ctx.Request().Context(), params.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusUnauthorized, "Invalid email or password")
		}
		return ctx.String(http.StatusInternalServerError, "Error occured getting your account")
	}

	if !auth.ComparePassword(params.Password, user.Password) {
		return ctx.String(http.StatusUnauthorized, "Invalid Email or Passwordd")
	}

	accessToken, err := token.GenerateAccessToken(user.ID, user.AccountType)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Error occured generating access token")
	}

	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Error occured generating refresh token")
	}

	handler.setAccessTokenToCookie(ctx, accessToken)
	handler.setRefreshTokenToCookie(ctx, refreshToken)
	ctx.Response().Header().Add("HX-Redirect", "/dashboard")
	return ctx.String(http.StatusOK, "ok")
}
