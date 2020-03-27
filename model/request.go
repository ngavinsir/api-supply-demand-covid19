package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// HasCreateRequest handles new request creation.
type HasCreateRequest interface {
	CreateRequest(ctx context.Context, requestItems []*models.RequestItem, applicantID string) (*models.Request, error)
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
