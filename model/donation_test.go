package model

import (
	"context"
	"sync"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
)

const (
	testDonationItemItemID  = "ItemID"
	testDonationItemItemID2 = "ItemID2"
	testDonationItemUnitID  = "UnitID"
	testDonationItemUnitID2 = "UnitID2"
	testDonationUserID      = "UserID"
	testDonationItemsLen    = 100
	testDonationCount       = 15
	testDonationStockCount  = "150.00"
	testDonationQuantity    = "1.5"
)

func TestDonation(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()

	donationDatastore := &DonationDataStore{DB: db}

	t.Run("Create", testCreateDonation(donationDatastore, &StockDataStore{DB: db}))
	t.Run("Get", testGetDonation(donationDatastore))
	t.Run("Get user", testGetUserDonation(donationDatastore))
}

func testCreateDonation(repo *DonationDataStore, stockRepo *StockDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString(testDonationQuantity)

		item := &models.Item{
			ID:   testDonationItemItemID,
			Name: "name",
		}

		item.Insert(context.Background(), repo, boil.Infer())

		unit := &models.Unit{
			ID:   testDonationItemUnitID,
			Name: "name",
		}

		unit.Insert(context.Background(), repo, boil.Infer())

		user := &models.User{
			ID:       testDonationUserID,
			Email:    "email@email.com",
			Password: "password",
			Name:     "name",
			Role:     "DONATOR",
		}

		user.Insert(context.Background(), repo, boil.Infer())

		var wg sync.WaitGroup
		var d string
		var mu sync.Mutex
		for i := 0; i < testDonationCount; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				donationItem := []*models.DonationItem{}
				for i := 0; i < testDonationItemsLen; i++ {
					item := &models.DonationItem{
						ItemID:   testDonationItemItemID,
						UnitID:   testDonationItemUnitID,
						Quantity: quantity,
					}
					donationItem = append(donationItem, item)
				}

				donation, err := repo.CreateOrUpdateDonation(context.Background(), donationItem, user.ID, CreateAction)

				if err != nil {
					t.Error(err)
				}

				if donation.ID == "" {
					t.Errorf("Want donation id assigned, got %s", donation.ID)
				}

				if got, want := len(donation.DonationItems), testDonationItemsLen; got != want {
					t.Errorf("Want donation items length %d, got %d", want, got)
				}

				for i := 0; i < testDonationItemsLen; i++ {
					if got, want := donation.DonationItems[i].DonationID, donation.ID; got != want {
						t.Errorf("Want donation item donation id %s, got %s", want, got)
					}
				}

				mu.Lock()
				if d == "" {
					d = donation.ID
				}
				mu.Unlock()
			}()
		}
		wg.Wait()

		t.Run("Update", testUpdateDonation(repo, testDonationUserID))
		t.Run("Accept", testAcceptDonation(repo, stockRepo, d))
	}
}

func testUpdateDonation(repo *DonationDataStore, userID string) func(t *testing.T) {
	return func(t *testing.T) {
		item := &models.Item{
			ID:   testDonationItemItemID2,
			Name: testDonationItemItemID2,
		}

		item.Insert(context.Background(), repo, boil.Infer())

		unit := &models.Unit{
			ID:   testDonationItemUnitID2,
			Name: testDonationItemUnitID2,
		}

		unit.Insert(context.Background(), repo, boil.Infer())

		donations, err := models.Donations(qm.Load(models.DonationRels.DonationItems)).All(context.Background(), repo)
		if err != nil {
			t.Error(err)
		}

		for _, donation := range donations {
			var quantity types.Decimal
			quantity.Big, _ = new(decimal.Big).SetString(testDonationQuantity)

			donationItems := donation.R.DonationItems
			for _, item := range donationItems {
				item.UnitID = testDonationItemUnitID2
				item.ItemID = testDonationItemItemID2
				item.Quantity = quantity
			}

			donationData, err := repo.UpdateDonation(
				context.Background(),
				donationItems,
				donation.ID,
			)
			if err != nil {
				t.Error(err)
			}

			for _, item := range donationData.DonationItems {
				if got, want := item.DonationID, donation.ID; got != want {
					t.Errorf("Want donation item id %s, got %s", want, got)
				}
				if got, want := item.Quantity, quantity; got != want {
					t.Errorf("Want donation item quantity %s, got %s", want, got)
				}
				if got, want := item.Item.Name, testDonationItemItemID2; got != want {
					t.Errorf("Want donation item item name %s, got %s", want, got)
				}
				if got, want := item.Unit.Name, testDonationItemUnitID2; got != want {
					t.Errorf("Want donation item unit name %s, got %s", want, got)
				}
			}
		}
	}
}

func testAcceptDonation(repo *DonationDataStore, stockRepo *StockDataStore, donationID string) func(t *testing.T) {
	return func(t *testing.T) {
		err := repo.AcceptDonation(context.Background(), donationID, stockRepo)
		if err != nil {
			t.Error(err)
		}

		donation, err := models.FindDonation(context.Background(), repo.DB, donationID)
		if err != nil {
			t.Error(err)
		}

		if got, want := donation.IsAccepted, true; got != want {
			t.Errorf("Want donation is accepted, got %v", got)
		}

		stock, err := models.Stocks(
			models.StockWhere.ItemID.EQ(testDonationItemItemID2),
		).One(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		if got, want := stock.Quantity.Big.String(), testDonationStockCount; got != want {
			t.Errorf("Want stock quantity %s, got %s", want, got)
		}
	}
}

func testGetDonation(repo *DonationDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		donations, err := models.Donations(
			qm.Load(models.DonationRels.Donator),
			qm.Load(models.DonationRels.DonationItems),
		).All(context.Background(), repo)
		if err != nil {
			t.Error(err)
		}

		donationID := donations[0].ID

		donation, err := repo.GetDonation(context.Background(), donationID)
		if err != nil {
			t.Error(err)
		}

		if got, want := donation.ID, donationID; got != want {
			t.Errorf("Want donation id %s, got %s", want, got)
		}

		if got, want := len(donation.DonationItems), len(donations[0].R.DonationItems); got != want {
			t.Errorf("Want donation items length %d, got %d", want, got)
		}

		// Check not found error
		_, err = repo.GetDonation(context.Background(), "randomDonationID")
		if err == nil {
			t.Errorf("Want error, got success")
		}
	}
}

func testGetUserDonation(repo *DonationDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		donations, err := repo.GetUserDonations(context.Background(), testDonationUserID, 0, testDonationCount)
		if err != nil {
			t.Error(err)
		}

		for _, donation := range donations {
			if got, want := donation.Donator.ID, testDonationUserID; got != want {
				t.Errorf("Want donator id %s, got %s", want, got)
			}
		}
	}
}
