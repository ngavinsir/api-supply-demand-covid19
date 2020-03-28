package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// HasCreateItem handles item creations.
type HasCreateItem interface {
	CreateItem(ctx context.Context, name string) (*models.Item, error)
}

// HasGetAllItem handles items retrieval.
type HasGetAllItem interface {
	GetAllItem(ctx context.Context) (*models.ItemSlice, error)
}

// ItemDatastore holds db information.
type ItemDatastore struct {
	*sql.DB
}

// CreateItem creates a new item with given unique name.
func (db *ItemDatastore) CreateItem(ctx context.Context, name string) (*models.Item, error) {
	item := &models.Item{
		ID: ksuid.New().String(),
		Name: name,
	}
	err := item.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return item, nil
}

// GetAllItem gets all Item.
func (db *ItemDatastore) GetAllItem(ctx context.Context) (*models.ItemSlice, error) {
	items, err := models.Items().All(ctx, db)
	if err != nil {
		return nil, err
	}

	return &items, nil
}