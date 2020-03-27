package database

import (
	"context"
	"database/sql"
	"os"

	// for test db
	_ "github.com/lib/pq"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// InitTestDB opens new test db connection.
func InitTestDB() (*sql.DB, error) {
	conn := "postgresql://postgres:postgres@localhost:7654/postgres?sslmode=disable"
	if envConn := os.Getenv("TEST_DATABASE_URL"); envConn != "" {
		conn = envConn
	}

	return connect(conn)
}

// ResetTestDB clears test db data.
func ResetTestDB(db *sql.DB) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	models.AllocationItems().DeleteAll(context.Background(), tx)
	models.Allocations().DeleteAll(context.Background(), tx)
	models.DonationItems().DeleteAll(context.Background(), tx)
	models.Stocks().DeleteAll(context.Background(), tx)
	models.RequestItems().DeleteAll(context.Background(), tx)
	models.Donations().DeleteAll(context.Background(), tx)
	models.Items().DeleteAll(context.Background(), tx)
	models.Units().DeleteAll(context.Background(), tx)
	models.Requests().DeleteAll(context.Background(), tx)
	models.Users().DeleteAll(context.Background(), tx)

	err = tx.Commit()
	return err
}