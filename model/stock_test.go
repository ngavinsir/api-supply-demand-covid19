package model

import (
	"context"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"
)

const (
	testStockID     = "ID"
	testStockItemID = "ItemID"
	testStockUnitID = "UnitID"
	testStockLen    = 10
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
		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString("1.5")

		item := &models.Item{
			ID:   testStockItemID,
			Name: "name",
		}

		item.Insert(context.Background(), repo, boil.Infer())

		unit := &models.Unit{
			ID:   testStockUnitID,
			Name: "name",
		}

		unit.Insert(context.Background(), repo, boil.Infer())

		for i := 0; i < testStockLen; i++ {
			stock, err := repo.CreateNewStock(context.Background(), &models.Stock{
				ItemID:   testStockItemID,
				UnitID:   testStockUnitID,
				Quantity: types.NewDecimal(&decimal.Big{}),
			})

			if err != nil {
				t.Error(err)
			}

			if stock.ID == "" {
				t.Errorf("Want stock id assigned, got %s", stock.ID)
			}
		}

		t.Run("Get all stock", testGetAllStock(repo))
	}
}

func testGetAllStock(repo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		stocks, err := repo.GetAllStock(context.Background(), 1, 10)
		if err != nil {
			t.Error(err)
		}

		if len(stocks.Data) != testStockLen {
			t.Errorf("Want stock count %d, got %d", testStockLen, len(stocks.Data))
		}
	}
}
