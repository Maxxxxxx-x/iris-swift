package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	router "github.com/Maxxxxxx-x/iris-swift/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start(host, port string) error
	Shutdown()
}

type EchoServer struct {
	echo *echo.Echo
	ctx  context.Context
}

func (server *EchoServer) Start(host, port string) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	return server.echo.Start(addr)
}

func (server *EchoServer) Shutdown() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.echo.Shutdown(ctx); err != nil {
		server.echo.Logger.Fatal(err)
	}
}

func New(querier sqlc.Querier, env string) Server {
	server := &EchoServer{
		ctx: context.Background(),
	}
	server.echo = echo.New()

	server.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogLatency:   true,
		LogProtocol:  true,
		LogRemoteIP:  true,
		LogMethod:    true,
		LogURIPath:   true,
		LogHost:      true,
		LogRequestID: true,
		LogReferer:   true,
		LogUserAgent: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("host", v.Host).
				Str("protocol", v.Protocol).
				Str("request_id", v.RequestID).
				Str("request_ip", v.RemoteIP).
				Str("referer", v.Referer).
				Str("path", v.URIPath).
				Str("method", v.Method).
				Str("user_agent", v.UserAgent).
				Int("status", v.Status).
				Str("latency", v.Latency.String()).
				Msg("")
			return nil
		},
	}))

	server.echo.Use(middleware.Secure())
	server.echo.Use(middleware.Recover())

	server.echo.HideBanner = true

	staticRootDir := "./views/static"
	if env != "dev" {
		staticRootDir = "./static"
	}

	staticRoot := server.echo.Group("/static")
	staticRoot.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   staticRootDir,
		Browse: true,
	}))

	router.RegisterRoutes(server.ctx, server.echo, querier)

	return server
}
