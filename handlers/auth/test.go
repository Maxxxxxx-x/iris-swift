package authHandler

import (
	"net/http"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/services/auth"
	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/labstack/echo/v4"
)

func (handler *AuthHandler) RegisterTestRoutes(api *echo.Group) {
	api.GET("/test", handler.CreateDummyAccount)
	handler.logger.Info().Msg("Registered test API routes")
}

const (
	email    = "test@example.com"
	password = "Password"
)

func (handler *AuthHandler) CreateDummyAccount(ctx echo.Context) error {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}

	userId, err := utils.GenerateULID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate user id")
	}

	err = handler.querier.CreateUser(ctx.Request().Context(), sqlc.CreateUserParams{
		ID:       userId.String(),
		Email:    email,
		Password: hash,
	})

	if err != nil {
		handler.logger.Error().Err(err).Msg("Faild to add user to db")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add user to db")
	}

	handler.logger.Info().Msg("Created dummy account")

	return nil
}
