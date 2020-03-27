package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/types"
)

// HasGetAllStock handles get stock data.
type HasGetAllStock interface {
	GetAllStock(ctx context.Context) ([]*StockData, error)
}

// HasCreateNewStock handles create stock data
type HasCreateNewStock interface {
	CreateAllStock(ctx context.Context, stock *models.Stock) (*models.Stock, error)
}

// StockDataStore holds db information.
type StockDataStore struct {
	*sql.DB
}

// GetAllStock returns stocks
func (db *StockDataStore) GetAllStock(ctx context.Context) ([]*StockData, error) {
	stocks, err := models.Stocks().All(ctx, db)
	stockData := []*StockData{}
	for _, stock := range stocks {
		item, err := stock.Item().One(ctx, db)
		if err != nil {
			return nil, err
		}

		unit, err := stock.Unit().One(ctx, db)
		if err != nil {
			return nil, err
		}

		data := &StockData{
			ID:       stock.ID,
			Name:     item.Name,
			Unit:     unit.Name,
			Quantity: stock.Quantity,
		}
		stockData = append(stockData, data)
	}

	return stockData, err
}

// CreateNewStock returns stock
func (db *StockDataStore) CreateNewStock(ctx context.Context, data *models.Stock) (*models.Stock, error) {
	stock := &models.Stock{
		ID:       ksuid.New().String(),
		ItemID:   data.ItemID,
		UnitID:   data.UnitID,
		Quantity: data.Quantity,
	}

	if err := stock.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return stock, nil
}

// StockData struct
type StockData struct {
	ID       string        `boil:"id" json:"id"`
	Name     string        `boil:"name" json:"name"`
	Unit     string        `boil:"unit" json:"unit"`
	Quantity types.Decimal `boil:"quantity" json:"int"`
}
