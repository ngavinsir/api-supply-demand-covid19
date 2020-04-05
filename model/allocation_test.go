package model

import (
	"context"
	"fmt"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"
)

var testStockQuantity types.Decimal

func init() {
	testStockQuantity.Big, _ = new(decimal.Big).SetString("15")
}

func TestAllocation(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()
	database.ResetTestDB(db)

	t.Run("Create", testCreateAllocation(&AllocationDatastore{DB: db}, &StockDataStore{DB: db}))
}

func testCreateAllocation(repo *AllocationDatastore, stockRepo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		unit := &models.Unit{
			ID:   testUnitID,
			Name: testUnitName,
		}
		if err := unit.Insert(context.Background(), repo.DB, boil.Infer()); err != nil {
			panic(err)
		}

		item := &models.Item{
			ID:   testItemID,
			Name: testItemName,
		}
		if err := item.Insert(context.Background(), repo.DB, boil.Infer()); err != nil {
			panic(err)
		}

		stock := &models.Stock{
			ID:       "ID",
			ItemID:   testItemID,
			UnitID:   testUnitID,
			Quantity: testStockQuantity,
		}
		if err := stock.Insert(context.Background(), repo.DB, boil.Infer()); err != nil {
			panic(err)
		}

		userDatastore := &UserDatastore{DB: repo.DB}
		user, err := userDatastore.CreateNewUser(context.Background(), &models.User{
			Email:    testUserEmail,
			Name:     testUserName,
			Role:     RoleAdmin,
			Password: testUserPassword,
		})
		if err != nil {
			panic(err)
		}

		requestDatastore := &RequestDatastore{DB: repo.DB}
		request, err := requestDatastore.CreateRequest(
			context.Background(),
			[]*models.RequestItem{
				&models.RequestItem{
					ItemID:   testItemID,
					Quantity: testStockQuantity,
					UnitID:   testUnitID,
				},
			},
			user.ID,
		)
		if err != nil {
			panic(err)
		}

		allocationItems := []*models.AllocationItem{
			&models.AllocationItem{
				ItemID:   testItemID,
				Quantity: testStockQuantity,
				UnitID:   testUnitID,
			},
		}

		allocationData, err := repo.CreateAllocation(
			context.Background(),
			&models.Allocation{
				AdminID:   user.ID,
				RequestID: request.ID,
			},
			allocationItems,
			stockRepo,
		)
		if err != nil {
			t.Error(err)
		}

		if allocationData.ID == "" {
			t.Errorf("Want allocation id assigned, got %s", allocationData.ID)
		}

		if got, want := len(allocationData.AllocationItems), 1; got != want {
			t.Errorf("Want allocation items count %d, got %d", want, got)
		}

		for _, allocationItem := range allocationData.AllocationItems {
			if allocationItem.ID == "" {
				t.Errorf("Want allocation item id assigned, got %s", allocationItem.ID)
			}
		}

		_, err = repo.CreateAllocation(
			context.Background(),
			&models.Allocation{
				AdminID:   user.ID,
				RequestID: request.ID,
			},
			allocationItems,
			stockRepo,
		)
		if err == nil {
			t.Error("Want create allocation error, got no error")
		}

		if got, want := err.Error(), fmt.Sprintf("stock not available for item with id %s", allocationItems[0].ItemID); got != want {
			t.Errorf("Want create allocation error msg %s, got %s", want, got)
		}
	}
}
