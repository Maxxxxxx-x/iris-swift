package keysHandler

import (
	"context"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/Maxxxxxx-x/iris-swift/middleware"
	_ "github.com/Maxxxxxx-x/iris-swift/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type IKeysHandler interface {
	RegisterUIRoutes(e *echo.Echo)
	RegisterV1APIRoutes(api *echo.Group)
}

type KeysHandler struct {
	ctx     context.Context
	logger  zerolog.Logger
	querier sqlc.Querier
}

func New(ctx context.Context, querier sqlc.Querier) IKeysHandler {
	var handler KeysHandler
	handler.ctx = ctx
	handler.logger = logger.NewLogger("keys-handler")
	handler.querier = querier

	return &handler
}

func (handler *KeysHandler) RegisterUIRoutes(e *echo.Echo) {
	handler.logger.Info().Msg("Registering UI routes")
}

func (handler *KeysHandler) RegisterV1APIRoutes(api *echo.Group) {
	handler.logger.Info().Msg("Registering API routes")
	keysGrp := api.Group("/keys")
	keysGrp.Use(middleware.AccessTokenRequired)

	api.GET("/", handler.GetKeysFromUser)
	api.GET("/:id", handler.GetKeyById)
	api.GET("/:id/usage", handler.GetKeyUsageById)

	api.POST("/", handler.CreateKey)
	api.POST("/:id/refresh", handler.RefreshKeyWithId)

	api.DELETE("/:id", handler.RevokeKeyWithId)

	handler.logger.Info().Msg("Registered GetKeys API routes")
}
