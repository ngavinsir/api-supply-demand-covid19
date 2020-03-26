package database

import (
	"database/sql"
	"os"
)

// InitDB opens new db connection by url.
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
	return db, nil
}
