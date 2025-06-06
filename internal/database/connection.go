package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type DB struct {
	*sql.DB
}

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	MaxConns int
	MaxIdle  int
}

var TestDB *DB

func ConnectMySQL(config *Config) (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 config.Username,
		Passwd:               config.Password,
		Net:                  "tcp",
		Addr:                 config.Host + ":" + fmt.Sprintf("%d", config.Port),
		DBName:               config.Database,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// Configure connection pool
	db.SetMaxOpenConns(config.MaxConns)
	db.SetMaxIdleConns(config.MaxIdle)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	logrus.Info("Database connection established successfully")

	return db, nil
}

// Close closes database connection
func (db *DB) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
