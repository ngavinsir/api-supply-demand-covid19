package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// Const for create or update action.
const (
	CreateAction = "CREATE"
	UpdateAction = "UPDATE"
)

// HasCreateOrUpdateDonation handles get donation data.
type HasCreateOrUpdateDonation interface {
	CreateOrUpdateDonation(ctx context.Context, data []*models.DonationItem, userID string, action string) (*DonationData, error)
}

// DonationDataStore holds db information.
type DonationDataStore struct {
	*sql.DB
}

// CreateOrUpdateDonation handles create new donation
func (db *DonationDataStore) CreateOrUpdateDonation(
	ctx context.Context, data []*models.DonationItem, userID string, action string) (*DonationData, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	donation := &models.Donation{
		ID:         ksuid.New().String(),
		Date:       time.Now(),
		IsAccepted: false,
		IsDonated:  false,
		DonatorID:  userID,
	}

	if err := donation.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return nil, err
	}

	var items []*models.DonationItem
	resultChan := make(chan struct {
		*models.DonationItem
		error
	})

	for _, item := range data {
		go func(item *models.DonationItem) {
			var err error

			switch action {
			case CreateAction:
				item.ID = ksuid.New().String()
				item.DonationID = donation.ID

				err = item.Insert(ctx, tx, boil.Infer())
			case UpdateAction:
				item.DonationID = donation.ID
				_, err = item.Update(ctx, tx, boil.Infer())
			}

			if err != nil {
				tx.Rollback()
				resultChan <- struct {
					*models.DonationItem
					error
				}{nil, err}
			}

			resultChan <- struct {
				*models.DonationItem
				error
			}{item, nil}
		}(item)
	}

	for i := 0; i < len(data); i++ {
		result := <-resultChan
		if result.error != nil {
			tx.Rollback()
			return nil, result.error
		}
		items = append(items, result.DonationItem)
	}

	donationData := &DonationData{
		Donation: donation,
		Items:    items,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return donationData, nil
}

// DonationData struct
type DonationData struct {
	Donation *models.Donation       `boil:"donation" json:"donation"`
	Items    []*models.DonationItem `boil:"items" json:"items"`
}
