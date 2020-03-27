package model

import (
	"context"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/types"
)

const (
	testStockID     = "ID"
	testStockItemID = "ItemID"
	testStockUnitID = "UnitID"
)

func TestStock(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()

	t.Run("Create", testCreateStock(&StockDataStore{DB: db}))
}

func testCreateStock(repo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		stock, err := repo.CreateNewStock(context.Background(), &models.Stock{
			ID:       testStockID,
			ItemID:   testStockItemID,
			UnitID:   testStockUnitID,
			Quantity: types.NewNullDecimal(decimal.New(0, 0)),
		})
		if err != nil {
			t.Error(err)
		}

		if stock.ID == "" {
			t.Errorf("Want stock id assigned, got %s", stock.ID)
		}

		t.Run("Get all stock", testGetAllStock(repo))
	}
}

func testGetAllStock(repo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		stocks, err := repo.GetAllStock(context.Background())
		if err != nil {
			t.Error(err)
		}

		if got, want := stocks[0].ID, testStockID; got != want {
			t.Errorf("Want stock id %s, got %s", want, got)
		}
	}
}
