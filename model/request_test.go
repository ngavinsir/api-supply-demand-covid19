package model

import (
	"context"
	"sync"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"
)

const (
	testRequestCount = 15
	testRequestItemCount = 1000
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
		var wg sync.WaitGroup

		for i := 0; i < testRequestCount; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				var quantity types.Decimal
				quantity.Big, _ = new(decimal.Big).SetString("15.5")

				var requestItems []*models.RequestItem
				for i := 0; i < testRequestItemCount; i++ {
					requestItem := &models.RequestItem{
						ItemID: itemID,
						Quantity: quantity,
						UnitID: unitID,
					}
					requestItems = append(requestItems, requestItem)
				}

				requestData, err := repo.CreateRequest(
					context.Background(), 
					requestItems, 
					userID,
				)
				if err != nil {
					t.Error(err)
				}

				if requestData.Request.ID == "" {
					t.Errorf("Want request id assigned, got %s", requestData.Request.ID)
				}

				if got, want := len(requestData.Items), testRequestItemCount; got != want {
					t.Errorf("Want request items count %d, got %d", want, got)
				}

				for _, item := range requestData.Items {
					if item.ID == "" {
						t.Errorf("Want request item id assigned, got %s", item.ID)
					}
				}
			}()
		}
		wg.Wait()
	}
}