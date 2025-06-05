package db

import (
	"database/sql"
	"fmt"

	"github.com/Maxxxxxx-x/iris-swift/config"
	_ "modernc.org/sqlite"
)

func TestConnection(db *sql.DB) error {
	return db.Ping()
}

func ConnectDatabase(config config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("file:%s.db", config.Database_Name)
	db, err := sql.Open("sqlite", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.Max_Open_Connections)
	db.SetMaxIdleConns(config.Max_Idle_Connections)

	return db, nil
}
