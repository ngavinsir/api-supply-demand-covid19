package model

import (
	"context"
	"sync"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/segmentio/ksuid"
)

const (
	testItemCount = 100
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
		var itemIDs []string
		var wg sync.WaitGroup
		var mu sync.Mutex

		for i := 0; i < testItemCount; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

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

				mu.Lock()
				itemIDs = append(itemIDs, item.ID)
				mu.Unlock()
			}()
		}
		wg.Wait()

		t.Run("Get all", testGetAllItem(repo))
		t.Run("Delete", testDeleteItem(repo, itemIDs))
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

func testDeleteItem(repo *ItemDatastore, itemIDs []string) func(t *testing.T) {
	return func(t *testing.T) {
		var wg sync.WaitGroup
		
		for _, itemID := range itemIDs {
			wg.Add(1)
			go func(itemID string) {
				defer wg.Done()

				err := repo.DeleteItem(context.Background(), itemID)
				if err != nil {
					t.Error(err)
				}
			}(itemID)
		}
		wg.Wait()

		items, err := repo.GetAllItem(context.Background())
		if err != nil {
			t.Error(err)
		}

		if got, want := len(*items), 0; got != want {
			t.Errorf("Want items count %d, got %d", want, got)
		}
	}
}