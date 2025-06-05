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

	"github.com/joho/godotenv"
)

func main() {
	appEnv := config.GetAppEnv()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load environment: %v\n", err)
		os.Exit(1)
	}
	config, err := config.GetConfig("./config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger.Init(config.Logger, config.Env.App_Env)

	var conn *sql.DB

	if conn, err = db.ConnectDatabase(config.Database); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	db.TestConnection(conn)
	querier := sqlc.New(conn)

	server := server.New(querier, appEnv)
	go func() {
		if err := server.Start(config.Env.App_Host, config.Env.App_Port); err != nil {
			logger.Fatal().Err(err).Msg("An error occured starting the server`")
		}
	}()
	server.Shutdown()
}
