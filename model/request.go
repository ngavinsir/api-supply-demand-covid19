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
	CreateRequest(ctx context.Context, requestItems []*models.RequestItem, applicantID string) (*RequestData, error)
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
) (*RequestData, error) {
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

	var items []*models.RequestItem
	resultChan := make(chan struct {
		*models.RequestItem
		error
	}) 

	for _, item := range requestItems {
		go func(item *models.RequestItem) {
			item.ID = ksuid.New().String()
			item.RequestID = request.ID
			
			if err := item.Insert(ctx, tx, boil.Infer()); err != nil {
				tx.Rollback()
				resultChan <- struct {
					*models.RequestItem
					error
				} { nil, err }
			}

			resultChan <- struct {
				*models.RequestItem
				error
			} { item, nil }
		}(item)
	}

	for i := 0; i < len(requestItems); i++ {
		result := <- resultChan
		if result.error != nil {
			tx.Rollback()
			return nil, result.error
		}
		items = append(items, result.RequestItem)
	}

	requestData := &RequestData{
		Request: request,
		Items:    items,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return requestData, nil
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

// RequestData struct
type RequestData struct {
	Request  *models.Request       `boil:"request" json:"request"`
	Items    []*models.RequestItem `boil:"items" json:"items"`
}
