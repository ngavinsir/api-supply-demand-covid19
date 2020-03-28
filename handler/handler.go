package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource    *AuthResource
	stockResource   *StockResource
	requestResource *RequestResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	requestDatastore := &model.RequestDatastore{DB: db}
	stockDataStore := &model.StockDataStore{DB: db}

	authResource := &AuthResource{UserDatastore: userDatastore}
	stockResource := &StockResource{StockDataStore: stockDataStore}
	requestResource := &RequestResource{
		requestDatastore: requestDatastore,
		userDatastore:    userDatastore,
	}

	api := &API{
		authResource:    authResource,
		stockResource:   stockResource,
		requestResource: requestResource,
	}
	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/auth", api.authResource.router())
	r.Mount("/stocks", api.stockResource.router())
	r.Mount("/requests", api.requestResource.router())

	return r
}

type contextKey struct {
	name string
}
