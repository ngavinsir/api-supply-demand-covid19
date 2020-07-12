package model

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/types"
)

func TestRequestItemAllocation(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()

	requestItemID, err := setupTest(db)
	if err != nil {
		t.Error(err)
		return
	}

	requestItemAllocationDatastore := &RequestItemAllocationDatastore{DB: db}

	var requestItemAllocation *models.RequestItemAllocation
	t.Run("Create request item allocation", func(t *testing.T) {
		requestItemAllocation, err = requestItemAllocationDatastore.CreateRequestItemAllocation(
			context.Background(),
			time.Now(),
			requestItemID,
			"description",
		)
		if err != nil {
			t.Error(err)
		}

		if requestItemAllocation.ID == "" {
			t.Error("want request item id generated, got null")
		}

		if got, want := requestItemAllocation.RequestItemID, requestItemID; got != want {
			t.Errorf("want request item id %s, got %s", want, got)
		}

		if got, want := requestItemAllocation.Description.String, "description"; got != want {
			t.Errorf("want request item allocation description %s, got %s", want, got)
		}
	})

	t.Run("Edit request item allocation", func(t *testing.T) {
		requestItemAllocation, err := requestItemAllocationDatastore.EditRequestItemAllocation(
			context.Background(),
			requestItemAllocation.ID,
			time.Now(),
			"description2",
		)
		if err != nil {
			t.Error(err)
		}

		if got, want := requestItemAllocation.Description.String, "description2"; got != want {
			t.Errorf("want request item allocation description %s, got %s", want, got)
		}
	})

	t.Run("Delete request item allocation", func(t *testing.T) {
		if err = requestItemAllocationDatastore.DeleteRequestItemAllocation(
			context.Background(),
			requestItemAllocation.ID,
		); err != nil {
			t.Error(err)
		}

		requestItemAllocationExist, err := models.RequestItemAllocationExists(
			context.Background(),
			requestItemAllocationDatastore.DB,
			requestItemAllocation.ID,
		)
		if err != nil {
			t.Error(err)
		}

		if got, want := requestItemAllocationExist, false; got != want {
			t.Errorf("want request item allocation deleted, got still exist")
		}
	})
}

func setupTest(db *sql.DB) (string, error) {
	userDatastore := &UserDatastore{DB: db}
	user, err := userDatastore.CreateNewUser(
		context.Background(),
		&models.User{
			Email:    "email",
			Password: "pass",
			Name:     "name",
			Role:     "DONATOR",
		},
	)
	if err != nil {
		return "", err
	}

	itemDatastore := &ItemDatastore{DB: db}
	item, err := itemDatastore.CreateItem(
		context.Background(),
		"item",
	)
	if err != nil {
		return "", err
	}

	unitDatastore := &UnitDatastore{DB: db}
	unit, err := unitDatastore.CreateUnit(
		context.Background(),
		"unit",
	)
	if err != nil {
		return "", err
	}

	var quantity types.Decimal
	quantity.Big, _ = new(decimal.Big).SetString("15.5")

	requestDatastore := &RequestDatastore{DB: db}
	request, err := requestDatastore.CreateRequest(
		context.Background(),
		[]*models.RequestItem{
			{
				ItemID:   item.ID,
				UnitID:   unit.ID,
				Quantity: quantity,
			},
		},
		user.ID,
	)

	return request.RequestItems[0].ID, nil
}
