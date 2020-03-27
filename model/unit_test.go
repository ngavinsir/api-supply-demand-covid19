package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/api-supply-demand-covid19/database"
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
		unit, err := repo.CreateUnit(context.Background(), "TEST_UNIT")
		if err != nil {
			t.Error(err)
		}

		if unit.ID == "" {
			t.Errorf("Want unit id assigned, got %s", unit.ID)
		}
		if got, want := unit.Name, "TEST_UNIT"; got != want {
			t.Errorf("Want unit name %s, got %s", want, got)
		}
	}
}