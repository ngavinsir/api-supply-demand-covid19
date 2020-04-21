package model

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

const (
	testRequestCount     = 20
	testRequestItemCount = 15
	testUnitID           = "TEST_UNIT_ID"
	testUnitID2          = "TEST_UNIT_ID_2"
	testUnitName         = "TEST_UNIT_NAME"
	testUnitName2        = "TEST_UNIT_NAME_2"
	testItemID           = "TEST_ITEM_ID"
	testItemID2          = "TEST_ITEM_ID_2"
	testItemName         = "TEST_ITEM_NAME"
	testItemName2        = "TEST_ITEM_NAME_2"
	testUserName         = "TEST_USER_NAME"
	testUserRole         = "TEST_USER_ROLE"
	testUserEmail        = "TEST_USER_EMAIL"
	testUserPassword     = "TEST_USER_PASSWORD"
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
		ID:   testUnitID,
		Name: testUnitName,
	}
	err = unit.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	unit2 := &models.Unit{
		ID:   testUnitID2,
		Name: testUnitName2,
	}
	err = unit2.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	item := &models.Item{
		ID:   testItemID,
		Name: testItemName,
	}
	err = item.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	item2 := &models.Item{
		ID:   testItemID2,
		Name: testItemName2,
	}
	err = item2.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	userDatastore := &UserDatastore{DB: db}
	user, err := userDatastore.CreateNewUser(context.Background(), &models.User{
		Email:    testUserEmail,
		Name:     testUserName,
		Role:     testUserRole,
		Password: testUserPassword,
	})
	if err != nil {
		panic(err)
	}

	t.Run("Create", testCreateRequest(&RequestDatastore{DB: db}, unit.ID, item.ID, user.ID))
	t.Run("Update", testUpdateRequest(&RequestDatastore{DB: db}, unit2.ID, item2.ID, user.ID))
	t.Run("UpdateWhenFulfilled", testUpdateRequestWhenFulfilled(&RequestDatastore{DB: db}, unit.ID, item.ID, user.ID))
	t.Run("Get", testGetRequest(&RequestDatastore{DB: db}, unit.ID, item.ID, user))
	t.Run("Get user", testGetUserRequests(&RequestDatastore{DB: db}, user.ID))
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
						ItemID:   itemID,
						Quantity: quantity,
						UnitID:   unitID,
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

				if requestData.ID == "" {
					t.Errorf("Want request id assigned, got %s", requestData.ID)
				}

				if got, want := len(requestData.RequestItems), testRequestItemCount; got != want {
					t.Errorf("Want request items count %d, got %d", want, got)
				}

				for _, item := range requestData.RequestItems {
					if item.ID == "" {
						t.Errorf("Want request item id assigned, got %s", item.ID)
					}
				}
			}()
		}
		wg.Wait()
	}
}

func testUpdateRequest(repo *RequestDatastore, unitID string, itemID string, userID string) func(t *testing.T) {
	return func(t *testing.T) {
		requests, err := models.Requests(qm.Load(models.RequestRels.RequestItems)).All(context.Background(), repo)
		if err != nil {
			t.Error(err)
		}

		for _, request := range requests {
			var quantity types.Decimal
			quantity.Big, _ = new(decimal.Big).SetString("25.5")

			requestItems := request.R.RequestItems
			for _, item := range requestItems {
				item.UnitID = unitID
				item.ItemID = itemID
				item.Quantity = quantity
			}

			requestData, err := repo.UpdateRequest(
				context.Background(),
				requestItems,
				request.ID,
			)
			if err != nil {
				t.Error(err)
			}

			for _, item := range requestData.RequestItems {
				if got, want := item.RequestID, request.ID; got != want {
					t.Errorf("Want request item id %s, got %s", want, got)
				}
				if got, want := item.Quantity, quantity; got != want {
					t.Errorf("Want request item quantity %s, got %s", want, got)
				}
				if got, want := item.Item.Name, testItemName2; got != want {
					t.Errorf("Want request item item name %s, got %s", want, got)
				}
				if got, want := item.Unit.Name, testUnitName2; got != want {
					t.Errorf("Want request item unit name %s, got %s", want, got)
				}
			}
		}
	}
}

func testUpdateRequestWhenFulfilled(repo *RequestDatastore, unitID string, itemID string, userID string) func(t *testing.T) {
	return func(t *testing.T) {

		request := &models.Request{
			ID:                  ksuid.New().String(),
			Date:                time.Now(),
			DonationApplicantID: userID,
			IsFulfilled:         true,
		}

		err := request.Insert(context.Background(), repo, boil.Infer())
		if err != nil {
			t.Error(err)
		}

		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString("25.5")

		item := &models.RequestItem{
			ID:        ksuid.New().String(),
			UnitID:    unitID,
			ItemID:    itemID,
			RequestID: request.ID,
			Quantity:  quantity,
		}

		err = item.Insert(context.Background(), repo, boil.Infer())
		if err != nil {
			t.Error(err)
		}
		item.UnitID = testUnitID2
		item.ItemID = testItemID2

		items := []*models.RequestItem{}
		items = append(items, item)

		_, err = repo.UpdateRequest(
			context.Background(),
			items,
			request.ID,
		)

		if err == nil {
			t.Errorf("Want error, got success")
		}

	}

}

func testGetRequest(repo *RequestDatastore, unitID string, itemID string, user *models.User) func(t *testing.T) {
	return func(t *testing.T) {
		var quantity types.Decimal
		quantity.Big, _ = new(decimal.Big).SetString("25.5")

		var requestItems []*models.RequestItem
		requestItems = append(requestItems, &models.RequestItem{
			ID:       ksuid.New().String(),
			UnitID:   unitID,
			ItemID:   itemID,
			Quantity: quantity,
		})

		requestData, err := repo.CreateRequest(
			context.Background(),
			requestItems,
			user.ID,
		)
		if err != nil {
			t.Error(err)
		}

		requestID := requestData.ID
		request, err := repo.GetRequest(context.Background(), requestID)
		if err != nil {
			t.Error(err)
		}

		if got, want := request.ID, requestID; got != want {
			t.Errorf("Want request id %s, got %s", want, got)
		}

		if got, want := request.DonationApplicant.ID, user.ID; got != want {
			t.Errorf("Want request applicant id %s, got %s", want, got)
		}

		if got, want := request.RequestItems[0].ID, requestItems[0].ID; got != want {
			t.Errorf("Want request item id %s, got %s", want, got)
		}

		_, err = repo.GetRequest(context.Background(), "randomUserID")
		if err == nil {
			t.Errorf("want error, got success")
		}
	}
}

func testGetUserRequests(repo *RequestDatastore, userID string) func(t *testing.T) {
	return func(t *testing.T) {
		requests, err := repo.GetUserRequests(context.Background(), userID, 0, testRequestCount)
		if err != nil {
			t.Error(err)
		}

		for _, request := range requests {
			if got, want := request.DonationApplicant.ID, userID; got != want {
				t.Errorf("Want request donation applicant id %s, got %s", want, got)
			}
		}
	}
}
