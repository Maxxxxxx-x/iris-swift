package rootHandler

import (
	"context"

	"github.com/Maxxxxxx-x/iris-swift/handlers"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/Maxxxxxx-x/iris-swift/views/pages"
	"github.com/labstack/echo/v4"
)

type IRootHandler interface {
	RegisterUIRoutes(e *echo.Echo)
}

type RootHandler struct {
	ctx context.Context
}

func New(ctx context.Context) IRootHandler {
	return &RootHandler{
		ctx: ctx,
	}
}

func (handler *RootHandler) RegisterUIRoutes(e *echo.Echo) {
	logger.Info().Msg("Registering Root rotues")
	e.GET("/", handler.ShowIndexPage)
}

func (handler *RootHandler) ShowIndexPage(ctx echo.Context) error {
	component := pages.IndexPage()
	handlers.RenderTempl(ctx, &component)
	return nil
}
