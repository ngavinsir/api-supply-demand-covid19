package database

import (
	"database/sql"
	"os"
	"time"
)

// InitDB opens new db connection.
func InitDB() (*sql.DB, error) {
	conn := "postgresql://postgres:postgres@localhost:6543/postgres?sslmode=disable"
	if envConn := os.Getenv("DATABASE_URL"); envConn != "" {
		conn = envConn
	}

	return connect(conn)
}

func connect(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
