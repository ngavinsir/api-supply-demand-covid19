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
	testStockItemID1 = "ItemID1"
	testStockUnitID1 = "UnitID1"
	testStockItemID2 = "ItemID2"
	testStockUnitID2 = "UnitID2"
	testStock1Len    = 10
	testStock2Len    = 12
)

var testQuantity types.Decimal

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

	testQuantity.Big, _ = new(decimal.Big).SetString("1.2")

	t.Run("Create & Update", testCreateAndUpdateStock(&StockDataStore{DB: db}))
}

func testCreateAndUpdateStock(repo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString("1.5")

		item := &models.Item{
			ID:   testStockItemID1,
			Name: "name1",
		}

		item.Insert(context.Background(), repo, boil.Infer())

		unit := &models.Unit{
			ID:   testStockUnitID1,
			Name: "name1",
		}

		unit.Insert(context.Background(), repo, boil.Infer())

		item = &models.Item{
			ID:   testStockItemID2,
			Name: "name2",
		}

		item.Insert(context.Background(), repo, boil.Infer())

		unit = &models.Unit{
			ID:   testStockUnitID2,
			Name: "name2",
		}

		unit.Insert(context.Background(), repo, boil.Infer())

		for i := 0; i < testStock1Len; i++ {
			stock, err := repo.CreateOrUpdateStock(context.Background(), &models.Stock{
				ItemID:   testStockItemID1,
				UnitID:   testStockUnitID1,
				Quantity: testQuantity,
			}, nil)

			if err != nil {
				t.Error(err)
			}

			if stock.ID == "" {
				t.Errorf("Want stock id assigned, got %s", stock.ID)
			}
		}

		for i := 0; i < testStock2Len; i++ {
			stock, err := repo.CreateOrUpdateStock(context.Background(), &models.Stock{
				ItemID:   testStockItemID2,
				UnitID:   testStockUnitID2,
				Quantity: testQuantity,
			}, nil)

			if err != nil {
				t.Error(err)
			}

			if stock.ID == "" {
				t.Errorf("Want stock id assigned, got %s", stock.ID)
			}
		}

		t.Run("Get all stock", testGetAllStock(repo))
		t.Run("Is available", testIsAvailableStock(repo))
	}
}

func testGetAllStock(repo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		stocks, _, err := repo.GetAllStock(context.Background(), 0, 10)
		if err != nil {
			t.Error(err)
		}

		if len(stocks) != 2 {
			t.Errorf("Want stock count %d, got %d", 2, len(stocks))
		}

		stock1, err := models.Stocks(
			models.StockWhere.ItemID.EQ(testStockItemID1),
		).One(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		if got, want := stock1.Quantity.Big.String(), "12.00"; got != want {
			t.Errorf("Want stock 1 quantity %s, got %s", want, got)
		}

		stock2, err := models.Stocks(
			models.StockWhere.ItemID.EQ(testStockItemID2),
		).One(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		if got, want := stock2.Quantity.Big.String(), "14.40"; got != want {
			t.Errorf("Want stock 2 quantity %s, got %s", want, got)
		}
	}
}

func testIsAvailableStock(repo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		var testAvailableQuantity types.Decimal
		testAvailableQuantity.Big, _ = new(decimal.Big).SetString("13.00")

		stock1, err := models.Stocks(
			models.StockWhere.ItemID.EQ(testStockItemID1),
		).One(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		isAvailable, err := repo.IsStockAvailable(
			context.Background(),
			stock1.ItemID,
			stock1.UnitID,
			testAvailableQuantity,
		)
		if err != nil {
			t.Error(err)
		}

		if got, want := isAvailable, false; got != want {
			t.Errorf("Want stock 1 not to be available, got %v", got)
		}

		stock2, err := models.Stocks(
			models.StockWhere.ItemID.EQ(testStockItemID2),
		).One(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		isAvailable, err = repo.IsStockAvailable(
			context.Background(),
			stock2.ItemID,
			stock2.UnitID,
			testAvailableQuantity,
		)
		if err != nil {
			t.Error(err)
		}

		if got, want := isAvailable, true; got != want {
			t.Errorf("Want stock 2 to be available, got %v", got)
		}
	}
}
