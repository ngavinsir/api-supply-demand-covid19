package model

import (
	"context"
	"sync"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/segmentio/ksuid"
)

const (
	testUnitCount = 100
)

func TestUnit(t *testing.T) {
	db, err := database.InitTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		database.ResetTestDB(db)
		db.Close()
	}()

	t.Run("Create", testCreateUnit(&UnitDatastore{DB: db}))
}

func testCreateUnit(repo *UnitDatastore) func(t *testing.T) {
	return func(t *testing.T) {
		var unitIDs []string
		var wg sync.WaitGroup
		var mu sync.Mutex

		for i := 0; i < testUnitCount; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				testName := ksuid.New().String()
				unit, err := repo.CreateUnit(context.Background(), testName)
				if err != nil {
					t.Error(err)
				}

				if unit.ID == "" {
					t.Errorf("Want unit id assigned, got %s", unit.ID)
				}
				if got, want := unit.Name, testName; got != want {
					t.Errorf("Want unit name %s, got %s", want, got)
				}

				mu.Lock()
				unitIDs = append(unitIDs, unit.ID)
				mu.Unlock()
			}()
		}
		wg.Wait()

		t.Run("Get all", testGetAllUnit(repo))
		t.Run("Delete", testDeleteUnit(repo, unitIDs))
	}
}

func testGetAllUnit(repo *UnitDatastore) func(t *testing.T) {
	return func(t *testing.T) {
		units, err := repo.GetAllUnit(context.Background())
		if err != nil {
			t.Error(err)
		}

		if got, want := len(*units), testUnitCount; got != want {
			t.Errorf("Want units count %d, got %d", want, got)
		}
	}
}

func testDeleteUnit(repo *UnitDatastore, unitIDs []string) func(t *testing.T) {
	return func(t *testing.T) {
		var wg sync.WaitGroup

		for _, unitID := range unitIDs {
			wg.Add(1)

			go func(unitID string) {
				defer wg.Done()

				err := repo.DeleteUnit(context.Background(), unitID)
				if err != nil {
					t.Error(err)
				}
			}(unitID)
		}
		wg.Wait()

		units, err := repo.GetAllUnit(context.Background())
		if err != nil {
			t.Error(err)
		}

		if got, want := len(*units), 0; got != want {
			t.Errorf("Want units count %d, got %d", want, got)
		}
	}
}
