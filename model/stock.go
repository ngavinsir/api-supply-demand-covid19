package model

import (
	"context"
	"database/sql"
	"math"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

// HasGetAllStock handles get stock data.
type HasGetAllStock interface {
	GetAllStock(ctx context.Context, page int, size int) (*StockDataPage, error)
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
func (db *StockDataStore) GetAllStock(ctx context.Context, page int, size int) (*StockDataPage, error) {
	offset := (page - 1) * size
	limit := size

	stocksCount, err := models.Stocks().Count(ctx, db)
	if err != nil {
		return nil, err
	}
	isLast := (int(stocksCount) - (size * limit)) < size
	isFirst := page == 1
	totalPages := int(math.Ceil(float64(stocksCount) / float64(size)))
	if totalPages == 0 {
		totalPages = 1
	}

	pages := &Page{
		Current: page,
		Total:   totalPages,
		First:   isFirst,
		Last:    isLast,
	}

	stocks, err := models.Stocks(qm.Offset(offset), qm.Limit(limit)).All(ctx, db)
	if err != nil {
		return nil, err
	}

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
	result := &StockDataPage{
		Data:  stockData,
		Pages: pages,
	}

	return result, err
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

// StockDataPage struct
type StockDataPage struct {
	Data  []*StockData
	Pages *Page
}

// StockData struct
type StockData struct {
	ID       string        `boil:"id" json:"id"`
	Name     string        `boil:"name" json:"name"`
	Unit     string        `boil:"unit" json:"unit"`
	Quantity types.Decimal `boil:"quantity" json:"quantity"`
}

// Page struct
type Page struct {
	Current int  `boil:"current" json:"current"`
	Total   int  `boil:"total" json:"total"`
	First   bool `boil:"first" json:"first"`
	Last    bool `boil:"last" json:"last"`
}
