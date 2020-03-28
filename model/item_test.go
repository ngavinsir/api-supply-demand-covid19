package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/segmentio/ksuid"
)

const (
	testItemCount = 10
)

func TestItem(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()

	t.Run("Create", testCreateItem(&ItemDatastore{DB: db}))
}

func testCreateItem(repo *ItemDatastore) func(t *testing.T) {
	return func(t *testing.T) {
		for i := 0; i < testItemCount; i++ {
			testName := ksuid.New().String()
			item, err := repo.CreateItem(context.Background(), testName)
			if err != nil {
				t.Error(err)
			}

			if item.ID == "" {
				t.Errorf("Want item id assigned, got %s", item.ID)
			}
			if got, want := item.Name, testName; got != want {
				t.Errorf("Want item name %s, got %s", want, got)
			}
		}

		t.Run("Get all", testGetAllItem(repo))
	}
}

func testGetAllItem(repo *ItemDatastore) func(t *testing.T) {
	return func(t *testing.T) {
		items, err := repo.GetAllItem(context.Background())
		if err != nil {
			t.Error(err)
		}

		if got, want := len(*items), testItemCount; got != want {
			t.Errorf("Want items count %d, got %d", want, got)
		}
	}
}