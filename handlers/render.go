package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderTempl(ctx echo.Context, component *templ.Component) {
	(*component).Render(ctx.Request().Context(), ctx.Response().Writer)
}
