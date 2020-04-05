package database

import (
	"os"
	"time"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
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
	if err := execSchema(db); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func execSchema(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
    m, err := migrate.NewWithDatabaseInstance(
        "github://ngavinsir:5373136ec972e61a901abcb1e0a39e0b442638b7@ngavinsir/api-supply-demand-covid19/tree/staging/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

    if err := m.Up(); err != nil {
		return err
	}

	return nil
}
