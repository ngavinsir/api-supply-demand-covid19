package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// HasCreateUnit handles unit creations.
type HasCreateUnit interface {
	CreateUnit(ctx context.Context, name string) (*models.Unit, error)
}

// HasGetAllUnit handles units retrieval.
type HasGetAllUnit interface {
	GetAllUnit(ctx context.Context) (*models.UnitSlice, error)
}

// HasDeleteUnit handles item deletion.
type HasDeleteUnit interface {
	DeleteUnit(ctx context.Context, unitID string) error
}

// UnitDatastore holds db information.
type UnitDatastore struct {
	*sql.DB
}

// CreateUnit creates a new unit with given unique name.
func (db *UnitDatastore) CreateUnit(ctx context.Context, name string) (*models.Unit, error) {
	unit := &models.Unit{
		ID: ksuid.New().String(),
		Name: name,
	}
	err := unit.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return unit, nil
}

// GetAllUnit gets all unit.
func (db *UnitDatastore) GetAllUnit(ctx context.Context) (*models.UnitSlice, error) {
	units, err := models.Units().All(ctx, db)
	if err != nil {
		return nil, err
	}

	return &units, nil
}

// DeleteUnit deletes unit by given id.
func (db *UnitDatastore) DeleteUnit(ctx context.Context, unitID string) error {
	_, err := models.Units(models.UnitWhere.ID.EQ(unitID)).DeleteAll(ctx, db)
	
	return err
}