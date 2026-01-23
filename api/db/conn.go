package db

import (
	"database/sql"
	"log/slog"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "user:password@tcp(127.0.0.1:3306)/test"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	var maxOpenConnection int
	maxOpenConnectionStr := os.Getenv("MYSQL_MAX_CONNECTIONS")
	if maxOpenConnectionStr != "" {
		maxOpenConnection, _ = strconv.Atoi(maxOpenConnectionStr)
	}

	slog.Info("max open connections", "max", maxOpenConnection)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetMaxIdleConns(10)

	return db, nil
}
