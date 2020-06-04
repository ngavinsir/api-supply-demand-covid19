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
	. "github.com/volatiletech/sqlboiler/queries/qm"
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
		requestRepo interface{ HasGetRequest },
	) (*AllocationData, error)
}

// HasGetAllAllocations handles allocations retrieval.
type HasGetAllAllocations interface {
	GetAllAllocations(ctx context.Context, offset int, limit int, requestRepo interface{ HasGetRequest }) ([]*AllocationData, error)
}

// HasGetTotalAllocationCount handles allocations count retrieval.
type HasGetTotalAllocationCount interface {
	GetTotalAllocationCount(ctx context.Context) (int64, error)
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
	requestRepo interface{ HasGetRequest },
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
			Item:     result.AllocationItem.R.Item,
			Unit:     result.AllocationItem.R.Unit,
			Quantity: result.AllocationItem.Quantity,
		})
	}

	allocation.L.LoadAllocator(ctx, tx, true, allocation, nil)
	allocation.R.Allocator.Password = ""

	allocationRequest, err := requestRepo.GetRequest(ctx, allocation.RequestID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	allocationData := &AllocationData{
		ID:              allocation.ID,
		Date:            allocation.Date,
		Allocator:       allocation.R.Allocator,
		PhotoURL:        allocation.PhotoURL.String,
		Request:         allocationRequest,
		AllocationItems: items,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return allocationData, nil
}

// GetAllAllocations gets all allocations.
func (db *AllocationDatastore) GetAllAllocations(
	ctx context.Context,
	offset int,
	limit int,
	requestRepo interface{ HasGetRequest },
) ([]*AllocationData, error) {
	allocations, err := models.Allocations(
		Load("AllocationItems.Item"),
		Load("AllocationItems.Unit"),
		Load(models.AllocationRels.Allocator),
		Offset(offset),
		Limit(limit),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var allocationData []*AllocationData
	for _, allocation := range allocations {
		allocation.R.Allocator.Password = ""

		var allocationItems []*AllocationItemData
		for _, item := range allocation.R.AllocationItems {
			allocationItems = append(allocationItems, &AllocationItemData{
				ID:       item.ID,
				Item:     item.R.Item,
				Unit:     item.R.Unit,
				Quantity: item.Quantity,
			})
		}

		allocationRequest, err := requestRepo.GetRequest(ctx, allocation.RequestID)
		if err != nil {
			return nil, err
		}

		allocationData = append(allocationData, &AllocationData{
			ID:              allocation.ID,
			Date:            allocation.Date,
			AllocationItems: allocationItems,
			Allocator:       allocation.R.Allocator,
			PhotoURL:        allocation.PhotoURL.String,
			Request:         allocationRequest,
		})
	}

	return allocationData, nil
}

// GetTotalAllocationCount returns total allocation count.
func (db *AllocationDatastore) GetTotalAllocationCount(ctx context.Context) (int64, error) {
	totalAllocationCount, err := models.Allocations().Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return totalAllocationCount, nil
}

// AllocationData struct
type AllocationData struct {
	ID              string                `json:"id"`
	Date            time.Time             `json:"date"`
	Request         *RequestData          `json:"request"`
	Allocator       *models.User          `json:"allocator,omitempty"`
	PhotoURL        string                `json:"photoUrl"`
	AllocationItems []*AllocationItemData `json:"items"`
}

// AllocationItemData struct
type AllocationItemData struct {
	ID           string        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Item         *models.Item  `boil:"item" json:"item" toml:"item" yaml:"item"`
	Unit         *models.Unit  `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	Quantity     types.Decimal `boil:"quantity" json:"quantity"`
	AllocationID string        `boil:"allocation_id" json:"allocation_id,omitempty"`
}
