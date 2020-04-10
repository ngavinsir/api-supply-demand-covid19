package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

const (
	CreateAction = "CREATE"
	UpdateAction = "UPDATE"
)

// HasCreateOrUpdate handles get donation data.
type HasCreateOrUpdate interface {
	CreateOrUpdateDonation(ctx context.Context, data []*models.DonationItem, userID string, action string) (*DonationData, error)
}

// HasAcceptDonation accepts donation by given id.
type HasAcceptDonation interface {
	AcceptDonation(ctx context.Context, donationID string, stockRepo interface{ HasCreateOrUpdateStock }) error
}

// HasGetDonation handles get donation detail
type HasGetDonation interface {
	GetDonation(ctx context.Context, donationID string) (*DonationData, error)
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

// AcceptDonation accepts donation by given id
func (db *DonationDataStore) AcceptDonation(
	ctx context.Context,
	donationID string,
	stockRepo interface{ HasCreateOrUpdateStock },
) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}

	donation, err := models.Donations(
		models.DonationWhere.ID.EQ(donationID),
		Load(models.DonationRels.DonationItems),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if donation.IsAccepted {
		tx.Rollback()
		return errors.New("donation has already been accepted")
	}

	donation.IsAccepted = true
	_, err = donation.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, donationItem := range donation.R.DonationItems {
		_, err := stockRepo.CreateOrUpdateStock(ctx, &models.Stock{
			ItemID:   donationItem.ItemID,
			UnitID:   donationItem.UnitID,
			Quantity: donationItem.Quantity,
		}, tx)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// GetDonation handles get donation detail by given id
func (db *DonationDataStore) GetDonation(
	ctx context.Context,
	donationID string,
) (*DonationData, error) {
	donation, err := models.Donations(
		models.DonationWhere.ID.EQ(donationID),
		Load(models.DonationRels.DonationItems),
	).One(ctx, db)
	if err != nil {
		return nil, errors.New("donation id not found")
	}

	donationData := &DonationData{
		Donation: donation,
		Items:    donation.R.DonationItems,
	}

	return donationData, nil
}

// DonationData struct
type DonationData struct {
	Donation *models.Donation       `boil:"donation" json:"donation"`
	Items    []*models.DonationItem `boil:"items" json:"items"`
}
