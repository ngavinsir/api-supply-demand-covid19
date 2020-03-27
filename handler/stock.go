package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// StockResource holds stock data store information.
type StockResource struct {
	*model.StockDataStore
}

func (store *StockResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", GetAllStock(store))

	return r
}

// GetAllStock return stocks
func GetAllStock(repo interface {
	model.HasGetAllStock
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stocks, err := repo.GetAllStock(r.Context())

		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		stockResponse := &StockResponse{
			Success: true,
			Data:    stocks,
		}

		render.JSON(w, r, stockResponse)
	}
}

// StockResponse struct
type StockResponse struct {
	Success bool               `boil:"success" json:"success"`
	Data    []*model.StockData `boil:"data" json:"data"`
}
