package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"
)

// HasCreateAllocation interface
type HasCreateAllocation interface {
	CreateAllocation(
		ctx context.Context,
		allocation *models.Allocation,
		allocationItems []*models.AllocationItem,
		stockRepo interface {
			HasIsStockAvailable
			HasCreateOrUpdateStock
		},
	) (*AllocationData, error)
}

// AllocationDatastore holds db information.
type AllocationDatastore struct {
	*sql.DB
}

// CreateAllocation creates new allocation.
func (db *AllocationDatastore) CreateAllocation(
	ctx context.Context,
	allocation *models.Allocation,
	allocationItems []*models.AllocationItem,
	stockRepo interface {
		HasIsStockAvailable
		HasCreateOrUpdateStock
	},
) (*AllocationData, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	allocation.ID = ksuid.New().String()
	if allocation.Date.IsZero() {
		allocation.Date = time.Now()
	}

	if err := allocation.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return nil, err
	}

	var items []*AllocationItemData
	resultChan := make(chan struct {
		*models.AllocationItem
		error
	})

	for _, item := range allocationItems {
		go func(item *models.AllocationItem) {
			isStockAvailable, err := stockRepo.IsStockAvailable(ctx, item.ItemID, item.UnitID, item.Quantity)
			if err != nil || !isStockAvailable {
				resultChan <- struct {
					*models.AllocationItem
					error
				}{nil, fmt.Errorf("stock not available for item with id %s", item.ItemID)}
				return
			}

			var negQuantity types.Decimal
			negQuantity.Big = &decimal.Big{}
			negQuantity.Big.Neg(item.Quantity.Big)
			_, err = stockRepo.CreateOrUpdateStock(ctx,
				&models.Stock{
					ItemID:   item.ItemID,
					UnitID:   item.UnitID,
					Quantity: negQuantity,
				},
				tx,
			)
			if err != nil {
				resultChan <- struct {
					*models.AllocationItem
					error
				}{nil, err}
				return
			}

			item.ID = ksuid.New().String()
			item.AllocationID = allocation.ID

			item.L.LoadItem(ctx, tx, true, item, nil)
			item.L.LoadUnit(ctx, tx, true, item, nil)

			if err := item.Insert(ctx, tx, boil.Infer()); err != nil {
				resultChan <- struct {
					*models.AllocationItem
					error
				}{nil, err}
				return
			}

			resultChan <- struct {
				*models.AllocationItem
				error
			}{item, nil}
		}(item)
	}

	for i := 0; i < len(allocationItems); i++ {
		result := <-resultChan
		if result.error != nil {
			tx.Rollback()
			return nil, result.error
		}
		items = append(items, &AllocationItemData{
			ID:       result.AllocationItem.ID,
			Item:     result.AllocationItem.R.Item.Name,
			Unit:     result.AllocationItem.R.Unit.Name,
			Quantity: result.AllocationItem.Quantity,
		})
	}

	allocationData := &AllocationData{
		ID:              allocation.ID,
		Date:            allocation.Date,
		AdminID:         allocation.AdminID,
		PhotoURL:        allocation.PhotoURL.String,
		RequestID:       allocation.RequestID,
		AllocationItems: items,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return allocationData, nil
}

// AllocationData struct
type AllocationData struct {
	ID              string                `json:"id"`
	Date            time.Time             `json:"date"`
	RequestID       string                `json:"requestID"`
	AdminID         string                `json:"adminID"`
	PhotoURL        string                `json:"photoUrl"`
	AllocationItems []*AllocationItemData `json:"items"`
}

// AllocationItemData struct
type AllocationItemData struct {
	ID           string        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Item         string        `boil:"item" json:"item" toml:"item" yaml:"item"`
	Unit         string        `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	Quantity     types.Decimal `boil:"quantity" json:"quantity"`
	AllocationID string        `boil:"allocation_id" json:"allocation_id,omitempty"`
}
