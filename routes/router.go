package router

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/handlers"
	authHandler "github.com/Maxxxxxx-x/iris-swift/handlers/auth"
	rootHandler "github.com/Maxxxxxx-x/iris-swift/handlers/root"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/labstack/echo/v4"
)

func showAllRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", " ")
	if err != nil {
		logger.Error().Err(err).Msg("Error occured showing routes")
	}

	os.WriteFile("routes.json", data, 0644)
}

func RegisterRoutes(ctx context.Context, e *echo.Echo, querier sqlc.Querier) {
	registry := []any{
		rootHandler.New(ctx),
		authHandler.New(ctx, querier),
	}

	api := e.Group("/api")

	for _, handler := range registry {
		if uiRegisterer, ok := handler.(handlers.UIRegisterer); ok {
			uiRegisterer.RegisterUIRoutes(e)
		}
		if apiRegisterer, ok := handler.(handlers.APIRegisterer); ok {
			apiRegisterer.RegisterV1APIRoutes(api.Group("/v1"))
		}
	}

	logger.Info().Msg("All routes registered!")

	showAllRoutes(e)
}
