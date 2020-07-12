package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

// HasCreateRequestItemAllocation create RequestItemAllocation
type HasCreateRequestItemAllocation interface {
	CreateRequestItemAllocation(
		ctx context.Context,
		allocationDate time.Time,
		requestItemID,
		description string,
	) (*models.RequestItemAllocation, error)
}

// HasEditRequestItemAllocation edit RequestItemAllocation
type HasEditRequestItemAllocation interface {
	EditRequestItemAllocation(
		ctx context.Context,
		requestItemAllocationID string,
		allocationDate time.Time,
		description string,
	) (*models.RequestItemAllocation, error)
}

// HasDeleteRequestItemAllocation delete RequestItemAllocation
type HasDeleteRequestItemAllocation interface {
	DeleteRequestItemAllocation(ctx context.Context, requestItemAllocationID string) error
}

// RequestItemAllocationDatastore holds db information.
type RequestItemAllocationDatastore struct {
	*sql.DB
}

// CreateRequestItemAllocation create RequestItemAllocation
func (db *RequestItemAllocationDatastore) CreateRequestItemAllocation(
	ctx context.Context,
	allocationDate time.Time,
	requestItemID,
	description string,
) (*models.RequestItemAllocation, error) {
	requestItemExist, err := models.RequestItems(models.RequestItemWhere.ID.EQ(requestItemID)).Exists(ctx, db)
	if err != nil || !requestItemExist {
		return nil, errors.New("can't find request item")
	}

	requestItemAllocation := &models.RequestItemAllocation{
		ID:             ksuid.New().String(),
		AllocationDate: allocationDate,
		Description:    null.StringFrom(description),
		RequestItemID:  requestItemID,
	}
	if err := requestItemAllocation.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return requestItemAllocation, nil
}

// EditRequestItemAllocation edit RequestItemAllocation
func (db *RequestItemAllocationDatastore) EditRequestItemAllocation(
	ctx context.Context,
	requestItemAllocationID string,
	allocationDate time.Time,
	description string,
) (*models.RequestItemAllocation, error) {
	requestItemAllocation, err := models.FindRequestItemAllocation(ctx, db, requestItemAllocationID)
	if err != nil {
		return nil, errors.New("can't find request item allocation")
	}

	requestItemAllocation.Description = null.StringFrom(description)
	if !allocationDate.IsZero() {
		requestItemAllocation.AllocationDate = allocationDate
	}

	if _, err = requestItemAllocation.Update(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return requestItemAllocation, nil
}

// DeleteRequestItemAllocation delete RequestItemAllocation
func (db *RequestItemAllocationDatastore) DeleteRequestItemAllocation(ctx context.Context, requestItemAllocationID string) error {
	requestItemAllocation, err := models.FindRequestItemAllocation(ctx, db, requestItemAllocationID)
	if err != nil {
		return errors.New("can't find request item allocation")
	}

	if _, err = requestItemAllocation.Delete(ctx, db); err != nil {
		return err
	}

	return nil
}
