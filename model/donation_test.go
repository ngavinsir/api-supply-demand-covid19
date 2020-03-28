package model

import (
	"context"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
)

const (
	testDonationItemItemID = "ItemID"
	testDonationItemUnitID = "UnitID"
	testDonationUserID     = "UserID"
	testDonationItemsLen   = 10
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

	t.Run("Create", testCreateDonation(&DonationDataStore{DB: db}))
}

func testCreateDonation(repo *DonationDataStore) func(t *testing.T) {
	return func(t *testing.T) {
		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString("1.5")

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

		donationItem := []*models.DonationItem{}
		for i := 0; i < testDonationItemsLen; i++ {
			item := &models.DonationItem{
				ItemID:   testDonationItemItemID,
				UnitID:   testDonationItemUnitID,
				Quantity: quantity,
			}
			donationItem = append(donationItem, item)
		}

		donation, err := repo.CreateDonation(context.Background(), donationItem, user.ID)

		if err != nil {
			t.Error(err)
		}

		if donation.Donation.ID == "" {
			t.Errorf("Want donation id assigned, got %s", donation.Donation.ID)
		}

		if got, want := len(donation.Items), testDonationItemsLen; got != want {
			t.Errorf("Want donation items length %d, got %d", want, got)
		}

		if got, want := donation.Donation.DonatorID, testDonationUserID; got != want {
			t.Errorf("Want donation donator id %s, got %s", want, got)
		}

		for i := 0; i < testDonationItemsLen; i++ {
			if got, want := donation.Items[i].DonationID, donation.Donation.ID; got != want {
				t.Errorf("Want donation item donation id %s, got %s", want, got)
			}
		}

	}
}
