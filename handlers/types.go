package handlers

import "github.com/labstack/echo/v4"

type UIRegisterer interface {
	RegisterUIRoutes(e *echo.Echo)
}

type APIRegisterer interface {
	RegisterV1APIRoutes(group *echo.Group)
}
