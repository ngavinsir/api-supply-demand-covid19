package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/segmentio/ksuid"
)

const (
	testUnitCount = 10
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
		for i := 0; i < testUnitCount; i++ {
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
		}

		t.Run("Get all", testGetAllUnit(repo))
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