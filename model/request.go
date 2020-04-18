package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

// HasCreateRequest handles new request creation.
type HasCreateRequest interface {
	CreateRequest(ctx context.Context, requestItems []*models.RequestItem, applicantID string) (*RequestData, error)
}

// HasGetAllRequest handles requests retrieval.
type HasGetAllRequest interface {
	GetAllRequest(ctx context.Context, offset int, limit int) ([]*RequestData, error)
}

// HasGetUserRequests handles get user requests
type HasGetUserRequests interface {
	GetUserRequests(ctx context.Context, userID string, offset int, limit int) ([]*RequestData, error)
}

// HasGetTotalRequestCount handles requests count retrieval.
type HasGetTotalRequestCount interface {
	GetTotalRequestCount(ctx context.Context) (int64, error)
}

// HasGetTotalUserRequestCount handles requests count retrieval.
type HasGetTotalUserRequestCount interface {
	GetTotalUserRequestCount(ctx context.Context, userID string) (int64, error)
}

// HasGetRequest handles requests retrieval.
type HasGetRequest interface {
	GetRequest(ctx context.Context, requestID string) (*RequestData, error)
}

// HasUpdateRequest handles update existing request.
type HasUpdateRequest interface {
	UpdateRequest(ctx context.Context, requestItems []*models.RequestItem, requestID string) (*RequestData, error)
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

	var items []*RequestItemData
	resultChan := make(chan struct {
		*models.RequestItem
		error
	})

	for _, item := range requestItems {
		go func(item *models.RequestItem) {
			item.ID = ksuid.New().String()
			item.RequestID = request.ID

			item.L.LoadItem(ctx, db, true, item, nil)
			item.L.LoadUnit(ctx, db, true, item, nil)

			if err := item.Insert(ctx, tx, boil.Infer()); err != nil {
				tx.Rollback()
				resultChan <- struct {
					*models.RequestItem
					error
				}{nil, err}
			}

			resultChan <- struct {
				*models.RequestItem
				error
			}{item, nil}
		}(item)
	}

	for i := 0; i < len(requestItems); i++ {
		result := <-resultChan
		if result.error != nil {
			tx.Rollback()
			return nil, result.error
		}
		items = append(items, &RequestItemData{
			ID:       result.RequestItem.ID,
			Item:     result.RequestItem.R.Item.Name,
			Unit:     result.RequestItem.R.Unit.Name,
			Quantity: result.RequestItem.Quantity,
		})
	}

	requestData := &RequestData{
		ID:           request.ID,
		Date:         request.Date,
		IsFulfilled:  request.IsFulfilled,
		RequestItems: items,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return requestData, nil
}

// UpdateRequest handles new request creation with given request items and request id.
func (db *RequestDatastore) UpdateRequest(
	ctx context.Context,
	requestItems []*models.RequestItem,
	requestID string,
) (*RequestData, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	request, err := models.FindRequest(ctx, tx, requestID)
	if err != nil {
		return nil, fmt.Errorf("can't find request with id: %s", requestID)
	}

	if request.IsFulfilled {
		return nil, errors.New("request is already fulfilled")
	}

	var items []*RequestItemData
	resultChan := make(chan struct {
		*models.RequestItem
		error
	})

	for _, item := range requestItems {
		go func(item *models.RequestItem, requestID string) {
			if _, err := models.RequestItems(
				models.RequestItemWhere.ID.EQ(item.ID),
				models.RequestItemWhere.RequestID.EQ(request.ID),
			).One(ctx, db); err != nil {
				resultChan <- struct {
					*models.RequestItem
					error
				}{nil, fmt.Errorf("can't find request item with id: %s", item.ID)}
			}

			item.RequestID = requestID
			item.L.LoadItem(ctx, db, true, item, nil)
			item.L.LoadUnit(ctx, db, true, item, nil)

			if _, err := item.Update(ctx, tx, boil.Infer()); err != nil {
				resultChan <- struct {
					*models.RequestItem
					error
				}{nil, err}
			}

			resultChan <- struct {
				*models.RequestItem
				error
			}{item, nil}
		}(item, requestID)
	}

	for i := 0; i < len(requestItems); i++ {
		result := <-resultChan
		if result.error != nil {
			tx.Rollback()
			return nil, result.error
		}
		items = append(items, &RequestItemData{
			ID:        result.RequestItem.ID,
			RequestID: result.RequestItem.RequestID,
			Item:      result.RequestItem.R.Item.Name,
			Unit:      result.RequestItem.R.Unit.Name,
			Quantity:  result.RequestItem.Quantity,
		})
	}

	requestData := &RequestData{
		ID:           request.ID,
		Date:         request.Date,
		IsFulfilled:  request.IsFulfilled,
		RequestItems: items,
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return requestData, nil
}

// GetAllRequest gets all requests.
func (db *RequestDatastore) GetAllRequest(ctx context.Context, offset int, limit int) ([]*RequestData, error) {
	requests, err := models.Requests(
		Load("RequestItems.Item"),
		Load("RequestItems.Unit"),
		Load(models.RequestRels.DonationApplicant),
		Offset(offset),
		Limit(limit),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var requestData []*RequestData
	for _, r := range requests {
		r.R.DonationApplicant.Password = ""

		var requestItems []*RequestItemData
		for _, item := range r.R.RequestItems {
			requestItems = append(requestItems, &RequestItemData{
				ID:       item.ID,
				Item:     item.R.Item.Name,
				Unit:     item.R.Unit.Name,
				Quantity: item.Quantity,
			})
		}

		requestData = append(requestData, &RequestData{
			ID:                r.ID,
			Date:              r.Date,
			IsFulfilled:       r.IsFulfilled,
			RequestItems:      requestItems,
			DonationApplicant: r.R.DonationApplicant,
		})
	}

	return requestData, nil
}

// GetUserRequests gets all user requests.
func (db *RequestDatastore) GetUserRequests(ctx context.Context, userID string, offset int, limit int) ([]*RequestData, error) {
	requests, err := models.Requests(
		models.RequestWhere.DonationApplicantID.EQ(userID),
		Load("RequestItems.Item"),
		Load("RequestItems.Unit"),
		Load(models.RequestRels.DonationApplicant),
		Offset(offset),
		Limit(limit),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var requestData []*RequestData
	for _, r := range requests {
		r.R.DonationApplicant.Password = ""

		var requestItems []*RequestItemData
		for _, item := range r.R.RequestItems {
			requestItems = append(requestItems, &RequestItemData{
				ID:       item.ID,
				Item:     item.R.Item.Name,
				Unit:     item.R.Unit.Name,
				Quantity: item.Quantity,
			})
		}

		requestData = append(requestData, &RequestData{
			ID:                r.ID,
			Date:              r.Date,
			IsFulfilled:       r.IsFulfilled,
			RequestItems:      requestItems,
			DonationApplicant: r.R.DonationApplicant,
		})
	}

	return requestData, nil
}

// GetTotalRequestCount returns total request count.
func (db *RequestDatastore) GetTotalRequestCount(ctx context.Context) (int64, error) {
	totalRequestCount, err := models.Requests().Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return totalRequestCount, nil
}

// GetTotalUserRequestCount returns total user request count.
func (db *RequestDatastore) GetTotalUserRequestCount(ctx context.Context, userID string) (int64, error) {
	totalRequestCount, err := models.Requests(
		models.RequestWhere.DonationApplicantID.EQ(userID),
	).Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return totalRequestCount, nil
}

// GetRequest handles get request detail given request id and applicant id.
func (db *RequestDatastore) GetRequest(
	ctx context.Context,
	requestID string,
) (*RequestData, error) {
	request, err := models.Requests(
		models.RequestWhere.ID.EQ(requestID),
		Load(models.RequestRels.DonationApplicant),
		Load(models.RequestRels.RequestItems),
	).One(ctx, db)
	if err != nil {
		return nil, errors.New("request id not found")
	}

	var requestItems []*RequestItemData
	for _, item := range request.R.RequestItems {
		item.L.LoadItem(ctx, db, true, item, nil)
		item.L.LoadUnit(ctx, db, true, item, nil)

		requestItemData := &RequestItemData{
			ID:       item.ID,
			Item:     item.R.Item.Name,
			Unit:     item.R.Unit.Name,
			Quantity: item.Quantity,
		}

		requestItems = append(requestItems, requestItemData)
	}

	requestData := &RequestData{
		ID:                request.ID,
		Date:              request.Date,
		IsFulfilled:       request.IsFulfilled,
		DonationApplicant: request.R.DonationApplicant,
		RequestItems:      requestItems,
	}

	return requestData, nil
}

// RequestData struct
type RequestData struct {
	ID                string             `json:"id"`
	Date              time.Time          `json:"date"`
	IsFulfilled       bool               `json:"isFulfilled"`
	DonationApplicant *models.User       `json:"donationApplicant,omitempty"`
	RequestItems      []*RequestItemData `json:"requestItems"`
}

type RequestItemData struct {
	ID        string        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Item      string        `boil:"item" json:"item" toml:"item" yaml:"item"`
	Unit      string        `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	Quantity  types.Decimal `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	RequestID string        `boil:"request_id" json:"request_id,omitempty" toml:"request_id" yaml:"request_id"`
}
