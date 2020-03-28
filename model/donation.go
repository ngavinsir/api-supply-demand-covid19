package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// HasCreateDonation handles get stock data.
type HasCreateDonation interface {
	CreateDonation(ctx context.Context, data []*models.DonationItem, userID string) (*DonationData, error)
}

// DonationDataStore holds db information.
type DonationDataStore struct {
	*sql.DB
}

// CreateDonation handles create new donation
func (db *DonationDataStore) CreateDonation(
	ctx context.Context, data []*models.DonationItem, userID string) (*DonationData, error) {
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

	items := []*models.DonationItem{}
	for _, item := range data {
		item := &models.DonationItem{
			ID:         ksuid.New().String(),
			DonationID: donation.ID,
			ItemID:     item.ItemID,
			UnitID:     item.UnitID,
			Quantity:   item.Quantity,
		}
		items = append(items, item)
		if err := item.Insert(ctx, tx, boil.Infer()); err != nil {
			tx.Rollback()
			return nil, err
		}
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
