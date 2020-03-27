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
	testUnitID = "TEST_UNIT_ID"
	testUnitName = "TEST_UNIT_NAME"
	testItemID = "TEST_ITEM_ID"
	testItemName = "TEST_ITEM_NAME"
	testUserName = "TEST_USER_NAME"
	testUserRole = "TEST_USER_ROLE"
	testUserEmail = "TEST_USER_EMAIL"
	testUserPassword = "TEST_USER_PASSWORD"
)

func TestRequest(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()
	
	unit := &models.Unit{
		ID: testUnitID,
		Name: testUnitName,
	}
	err = unit.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	item := &models.Item{
		ID: testItemID,
		Name: testItemName,
	}
	err = item.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	userDatastore := &UserDatastore{DB: db}
	user, err := userDatastore.CreateNewUser(context.Background(), &models.User{
		Email: testUserEmail,
		Name: testUserName,
		Role: testUserRole,
		Password: testUserPassword,
	})
	if err != nil {
		panic(err)
	}

	t.Run("Create", testCreateRequest(&RequestDatastore{DB: db}, unit.ID, item.ID, user.ID))
}

func testCreateRequest(repo *RequestDatastore, unitID string, itemID string, userID string) func(t *testing.T) {
	return func(t *testing.T) {
		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString("15.5")

		requestItem1 := &models.RequestItem{
			ItemID: itemID,
			Quantity: quantity,
			UnitID: unitID,
		}
		requestItem2 := &models.RequestItem{
			ItemID: itemID,
			Quantity: quantity,
			UnitID: unitID,
		}
		request, err := repo.CreateRequest(context.Background(), []*models.RequestItem{requestItem1, requestItem2}, userID)
		if err != nil {
			t.Error(err)
		}

		if request.ID == "" {
			t.Errorf("Want request id assigned, got %s", request.ID)
		}
		
		requestItems, err := request.RequestItems().All(context.Background(), repo.DB)
		if err != nil {
			t.Error(err)
		}

		if got, want := len(requestItems), 2; got != want {
			t.Errorf("Want request items count %d, got %d", want, got)
		}
	}
}