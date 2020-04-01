package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

// HasGetAllStock handles get stock data.
type HasGetAllStock interface {
	GetAllStock(ctx context.Context, offset int, limit int) ([]*StockData, int64, error)
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
func (db *StockDataStore) GetAllStock(ctx context.Context, offset int, limit int) ([]*StockData, int64, error) {
	stocksCount, err := models.Stocks().Count(ctx, db)
	if err != nil {
		return nil, 0, err
	}

	stocks, err := models.Stocks(
		Offset(offset),
		Limit(limit),
		Load(models.StockRels.Item),
		Load(models.StockRels.Unit),
	).All(ctx, db)
	if err != nil {
		return nil, 0, err
	}

	stockData := []*StockData{}
	for _, stock := range stocks {
		data := &StockData{
			ID:       stock.ID,
			Name:     stock.R.Item.Name,
			Unit:     stock.R.Unit.Name,
			Quantity: stock.Quantity,
		}
		stockData = append(stockData, data)
	}

	return stockData, stocksCount, err
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
	Quantity types.Decimal `boil:"quantity" json:"quantity"`
}

