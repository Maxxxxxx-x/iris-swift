package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/Maxxxxxx-x/iris-swift/db"
	"github.com/Maxxxxxx-x/iris-swift/db/sql/sqlc"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/Maxxxxxx-x/iris-swift/server"
	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load environment: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	appEnv := config.GetAppEnv()

	if appEnv == "dev" {
		loadDotEnv()
	}

	config, err := config.GetConfig("./config.yml")
	config.Env = appEnv
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger.Init(config.Logger, config.Env)

	var conn *sql.DB

	if conn, err = db.ConnectDatabase(config.Database); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	db.TestConnection(conn)
	querier := sqlc.New(conn)

	token.Init(config.JwtConfig)

	server := server.New(querier, appEnv)
	go func() {
		if err := server.Start(config.App.Host, config.App.Port); err != nil {
			logger.Fatal().Err(err).Msg("An error occured starting the server`")
		}
	}()
	server.Shutdown()
}
