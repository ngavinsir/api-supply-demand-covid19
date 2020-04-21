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

// Const for create or update action.
const (
	CreateAction = "CREATE"
	UpdateAction = "UPDATE"
)

// HasCreateOrUpdateDonation handles get donation data.
type HasCreateOrUpdateDonation interface {
	CreateOrUpdateDonation(ctx context.Context, data []*models.DonationItem, userID string, action string) (*DonationData, error)
}

// HasAcceptDonation accepts donation by given id.
type HasAcceptDonation interface {
	AcceptDonation(ctx context.Context, donationID string, stockRepo interface{ HasCreateOrUpdateStock }) error
}

// HasUpdateDonation handles update existing donation.
type HasUpdateDonation interface {
	UpdateDonation(ctx context.Context, donationItems []*models.DonationItem, donationID string) (*DonationData, error)
}

// HasGetDonation handles get donation detail
type HasGetDonation interface {
	GetDonation(ctx context.Context, donationID string) (*DonationData, error)
}

// HasGetUserDonations handles get user donations
type HasGetUserDonations interface {
	GetUserDonations(ctx context.Context, userID string, offset int, limit int) ([]*DonationData, error)
}

// HasGetAllDonations handles domains retrieval.
type HasGetAllDonations interface {
	GetAllDonations(ctx context.Context, offset int, limit int) ([]*DonationData, error)
}

// HasGetTotalDonationCount handles donations count retrieval.
type HasGetTotalDonationCount interface {
	GetTotalDonationCount(ctx context.Context) (int64, error)
}

// HasGetTotalUserDonationCount handles user donations count retrieval.
type HasGetTotalUserDonationCount interface {
	GetTotalUserDonationCount(ctx context.Context, userID string) (int64, error)
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

	var items []*DonationItemData
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

			item.L.LoadItem(ctx, db, true, item, nil)
			item.L.LoadUnit(ctx, db, true, item, nil)

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
		items = append(items, &DonationItemData{
			ID:         result.DonationItem.ID,
			Item:       result.DonationItem.R.Item,
			Unit:       result.DonationItem.R.Unit,
			Quantity:   result.DonationItem.Quantity,
			DonationID: result.DonationItem.DonationID,
		})
	}

	donationData := &DonationData{
		ID:            donation.ID,
		Date:          donation.Date,
		IsAccepted:    donation.IsAccepted,
		IsDonated:     donation.IsDonated,
		DonationItems: items,
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

// UpdateDonation updates donation with given donation items and donation id.
func (db *DonationDataStore) UpdateDonation(
	ctx context.Context,
	donationItems []*models.DonationItem,
	donationID string,
) (*DonationData, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	donation, err := models.FindDonation(ctx, tx, donationID)
	if err != nil {
		return nil, fmt.Errorf("can't find donation with id: %s", donationID)
	}

	if donation.IsAccepted {
		return nil, errors.New("donation has already been accepted")
	}

	var items []*DonationItemData
	resultChan := make(chan struct {
		*models.DonationItem
		error
	})

	for _, item := range donationItems {
		go func(item *models.DonationItem, donationID string) {
			if _, err := models.DonationItems(
				models.DonationItemWhere.ID.EQ(item.ID),
				models.DonationItemWhere.DonationID.EQ(donation.ID),
			).One(ctx, db); err != nil {
				resultChan <- struct {
					*models.DonationItem
					error
				}{nil, fmt.Errorf("can't find donation item with id: %s", item.ID)}
			}

			item.DonationID = donationID
			item.L.LoadItem(ctx, db, true, item, nil)
			item.L.LoadUnit(ctx, db, true, item, nil)

			if _, err := item.Update(ctx, tx, boil.Infer()); err != nil {
				resultChan <- struct {
					*models.DonationItem
					error
				}{nil, err}
			}

			resultChan <- struct {
				*models.DonationItem
				error
			}{item, nil}
		}(item, donationID)
	}

	for i := 0; i < len(donationItems); i++ {
		result := <-resultChan
		if result.error != nil {
			tx.Rollback()
			return nil, result.error
		}
		items = append(items, &DonationItemData{
			ID:         result.DonationItem.ID,
			DonationID: result.DonationItem.DonationID,
			Item:       result.DonationItem.R.Item,
			Unit:       result.DonationItem.R.Unit,
			Quantity:   result.DonationItem.Quantity,
		})
	}

	donationData := &DonationData{
		ID:            donation.ID,
		Date:          donation.Date,
		IsAccepted:    donation.IsAccepted,
		IsDonated:     donation.IsDonated,
		DonationItems: items,
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return donationData, nil
}

// GetDonation handles get donation detail by given id
func (db *DonationDataStore) GetDonation(
	ctx context.Context,
	donationID string,
) (*DonationData, error) {
	donation, err := models.Donations(
		models.DonationWhere.ID.EQ(donationID),
		Load("DonationItems.Item"),
		Load("DonationItems.Unit"),
		Load(models.DonationRels.Donator),
	).One(ctx, db)
	if err != nil {
		return nil, errors.New("donation id not found")
	}

	donation.R.Donator.Password = ""

	var donationItems []*DonationItemData
	for _, item := range donation.R.DonationItems {
		donationItems = append(donationItems, &DonationItemData{
			ID:       item.ID,
			Item:     item.R.Item,
			Unit:     item.R.Unit,
			Quantity: item.Quantity,
		})
	}

	donationData := &DonationData{
		ID:            donation.ID,
		Date:          donation.Date,
		Donator:       donation.R.Donator,
		IsAccepted:    donation.IsAccepted,
		IsDonated:     donation.IsDonated,
		DonationItems: donationItems,
	}

	return donationData, nil
}

// GetUserDonations gets all user donations.
func (db *DonationDataStore) GetUserDonations(ctx context.Context, userID string, offset int, limit int) ([]*DonationData, error) {
	donations, err := models.Donations(
		models.DonationWhere.DonatorID.EQ(userID),
		Load("DonationItems.Item"),
		Load("DonationItems.Unit"),
		Load(models.DonationRels.Donator),
		Offset(offset),
		Limit(limit),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var donationData []*DonationData
	for _, d := range donations {
		d.R.Donator.Password = ""

		var donationItems []*DonationItemData
		for _, item := range d.R.DonationItems {
			donationItems = append(donationItems, &DonationItemData{
				ID:       item.ID,
				Item:     item.R.Item,
				Unit:     item.R.Unit,
				Quantity: item.Quantity,
			})
		}

		donationData = append(donationData, &DonationData{
			ID:            d.ID,
			Date:          d.Date,
			IsAccepted:    d.IsAccepted,
			IsDonated:     d.IsDonated,
			DonationItems: donationItems,
			Donator:       d.R.Donator,
		})
	}

	return donationData, nil
}

// GetAllDonations gets all donations.
func (db *DonationDataStore) GetAllDonations(ctx context.Context, offset int, limit int) ([]*DonationData, error) {
	donations, err := models.Donations(
		Load("DonationItems.Item"),
		Load("DonationItems.Unit"),
		Load(models.DonationRels.Donator),
		Offset(offset),
		Limit(limit),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var donationData []*DonationData
	for _, d := range donations {
		d.R.Donator.Password = ""

		var donationItems []*DonationItemData
		for _, item := range d.R.DonationItems {
			donationItems = append(donationItems, &DonationItemData{
				ID:       item.ID,
				Item:     item.R.Item,
				Unit:     item.R.Unit,
				Quantity: item.Quantity,
			})
		}

		donationData = append(donationData, &DonationData{
			ID:            d.ID,
			Date:          d.Date,
			IsAccepted:    d.IsAccepted,
			IsDonated:     d.IsDonated,
			DonationItems: donationItems,
			Donator:       d.R.Donator,
		})
	}

	return donationData, nil
}

// GetTotalDonationCount returns total donation count.
func (db *DonationDataStore) GetTotalDonationCount(ctx context.Context) (int64, error) {
	totalDonationCount, err := models.Donations().Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return totalDonationCount, nil
}

// GetTotalUserDonationCount returns total user donation count.
func (db *DonationDataStore) GetTotalUserDonationCount(ctx context.Context, userID string) (int64, error) {
	totalDonationCount, err := models.Donations(
		models.DonationWhere.DonatorID.EQ(userID),
	).Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return totalDonationCount, nil
}

// DonationData struct
type DonationData struct {
	ID            string              `json:"id"`
	Date          time.Time           `json:"date"`
	IsAccepted    bool                `json:"isAccepted"`
	IsDonated     bool                `json:"isDonated"`
	Donator       *models.User        `json:"donator,omitempty"`
	DonationItems []*DonationItemData `json:"donationItems"`
}

// DonationItemData struct
type DonationItemData struct {
	ID         string        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Item       *models.Item  `boil:"item" json:"item" toml:"item" yaml:"item"`
	Unit       *models.Unit  `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	Quantity   types.Decimal `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	DonationID string        `boil:"donation_id" json:"donation_id,omitempty" toml:"donation_id" yaml:"donation_id"`
}
