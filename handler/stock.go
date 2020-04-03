package handler

import (
	"net/http"

	"github.com/ericlagergren/decimal"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// StockResource holds stock data store information.
type StockResource struct {
	*model.StockDataStore
}

func (store *StockResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(PaginationCtx)
	r.Get("/", GetAllStock(store))
	r.Post("/", CreateOrUpdateStock(store))

	return r
}

// GetAllStock return stocks
func GetAllStock(repo interface {
	model.HasGetAllStock
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paging, _ := r.Context().Value(PageCtxKey).(*Paging)

		stockData, totalCount, err := repo.GetAllStock(r.Context(), paging.Offset(), paging.Size)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		stockDataPage := &StockDataPage{
			Data:  stockData,
			Pages: paging.Pages(totalCount),
		}

		render.JSON(w, r, stockDataPage)
	}
}

// CreateOrUpdateStock creates or updates stock.
func CreateOrUpdateStock(repo interface{ model.HasCreateOrUpdateStock }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &StockRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		stock, err := repo.CreateOrUpdateStock(r.Context(), data.Stock, nil)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, stock)
	}
}

// StockDataPage struct
type StockDataPage struct {
	Data  []*model.StockData `boil:"data" json:"data"`
	Pages *Page              `boil:"pages" json:"pages"`
}

// StockRequest struct
type StockRequest struct {
	*models.Stock
}

// Bind StockRequest
func (req *StockRequest) Bind(r *http.Request) error {
	if req.Stock == nil || req.Stock.Quantity.Big.Cmp(&decimal.Big{}) == 0 || 
	(req.Stock.ID == "" && (req.Stock.ItemID == "" || req.Stock.UnitID == "")) {
		return ErrMissingReqFields
	}

	return nil
}
