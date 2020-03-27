package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

// HasCreateRequest handles new request creation.
type HasCreateRequest interface {
	CreateRequest(ctx context.Context, requestItems []*models.RequestItem, applicantID string) (*models.Request, error)
}

// HasGetAllRequest handles requests retrieval.
type HasGetAllRequest interface {
	GetAllRequest(ctx context.Context) (*models.RequestSlice, error)
}

// RequestDatastore holds db information.
type RequestDatastore struct {
	*sql.DB
}

// CreateRequest handles new request creation with given request items and applicant id.
func (db *RequestDatastore) CreateRequest(
	ctx context.Context,
	requestItems []*models.RequestItem,
	applicantID string,
) (*models.Request, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	
	request := &models.Request{
		ID:                  ksuid.New().String(),
		Date:                time.Now(),
		DonationApplicantID: applicantID,
	}
	
	err = request.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range requestItems {
		item.ID = ksuid.New().String()
	}
	err = request.AddRequestItems(ctx, tx, true, requestItems...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return request, nil
}

// GetAllRequest gets all requests.
func (db *RequestDatastore) GetAllRequest(ctx context.Context) (*models.RequestSlice, error) {
	requests, err := models.Requests(
		Load(models.RequestRels.RequestItems),
		Load(models.RequestRels.DonationApplicant),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	
	return &requests, nil
}
