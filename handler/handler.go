package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource  *AuthResource
	StockResource *StockResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	stockDataStore := &model.StockDataStore{DB: db}

	authResource := &AuthResource{UserDatastore: userDatastore}
	stockResource := &StockResource{StockDataStore: stockDataStore}

	api := &API{
		authResource:  authResource,
		StockResource: stockResource,
	}
	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/auth", api.authResource.router())
	r.Mount("/stocks", api.StockResource.router())

	return r
}

type contextKey struct {
	name string
}
