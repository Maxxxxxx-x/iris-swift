package authHandler

import (
	"context"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type IAuthHandler interface {
	RegisterUIRoutes(e *echo.Echo)
}

type AuthHandler struct {
	ctx       context.Context
	logger    zerolog.Logger
	querier   sqlc.Querier
	validator *validator.Validate
}

const COOKIE_NAME = "iris_auth"

func New(ctx context.Context, querier sqlc.Querier) IAuthHandler {
	var handler AuthHandler
	handler.ctx = ctx
	handler.logger = logger.NewLogger("auth")
	handler.validator = validator.New()
	handler.querier = querier

	return &handler
}

func (handler *AuthHandler) ValidateRequest(i any) error {
	return handler.validator.Struct(i)
}

func (handler *AuthHandler) RegisterUIRoutes(e *echo.Echo) {
	handler.logger.Info().Msg("Registering UI routes")
	handler.RegisterLoginUIRoutes(e)
}

func (handler *AuthHandler) RegisterV1APIRoutes(api *echo.Group) {
	handler.logger.Info().Msg("Registering API routes")
	authGrp := api.Group("/auth")
	handler.RegisterLoginAPIRoutes(authGrp)
}
