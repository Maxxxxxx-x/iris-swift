package authHandler

import (
	"database/sql"
	"errors"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/handlers"
	"github.com/Maxxxxxx-x/iris-swift/services/auth"
	"github.com/Maxxxxxx-x/iris-swift/services/session"
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
	}

	if err := handler.ValidateRequest(params); err != nil {
		handler.logger.Warn().Err(err).Msg("invalid request")
	}

	user, err := handler.querier.GetUserByIdentifier(ctx.Request().Context(), sqlc.GetUserByIdentifierParams{
		Username: params.Identifier,
		Email:    params.Identifier,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.logger.Warn().Msgf("Account not found wtih identifier %s", params.Identifier)
		}
		handler.logger.Error().Err(err).Msg("Server Error")
	}

	if !auth.ComparePassword(params.Password, user.Password) {
		handler.logger.Warn().Msgf("invalid password for account %s", user.ID)
	}

	// logged in! create cookie
	cookie, err := session.CreateCookie(ctx, COOKIE_NAME, map[string][string]{
	}, 86400)

	return nil
}
