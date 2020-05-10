package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource       *AuthResource
	unitResource       *UnitResource
	stockResource      *StockResource
	requestResource    *RequestResource
	itemResource       *ItemResource
	donationResource   *DonationResource
	allocationResource *AllocationResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	unitDatastore := &model.UnitDatastore{DB: db}
	requestDatastore := &model.RequestDatastore{DB: db}
	itemDatastore := &model.ItemDatastore{DB: db}
	stockDatastore := &model.StockDataStore{DB: db}
	donationDatastore := &model.DonationDataStore{DB: db}
	allocationDatastore := &model.AllocationDatastore{DB: db}
	passwordResetRequestDatastore := &model.PasswordResetRequestDatastore{DB: db}

	authResource := &AuthResource{
		UserDatastore:                 userDatastore,
		PasswordResetRequestDatastore: passwordResetRequestDatastore,
	}
	unitResource := &UnitResource{
		UnitDatastore: unitDatastore,
		UserDatastore: userDatastore,
	}
	stockResource := &StockResource{
		StockDataStore: stockDatastore,
		UserDatastore:  userDatastore,
	}
	requestResource := &RequestResource{
		requestDatastore: requestDatastore,
		userDatastore:    userDatastore,
	}
	itemResource := &ItemResource{
		ItemDatastore: itemDatastore,
		UserDatastore: userDatastore,
	}
	donationResource := &DonationResource{
		DonationDataStore: donationDatastore,
		UserDatastore:     userDatastore,
		StockDataStore:    stockDatastore,
	}
	allocationResource := &AllocationResource{
		AllocationDatastore: allocationDatastore,
		StockDataStore:      stockDatastore,
		UserDatastore:       userDatastore,
	}

	api := &API{
		authResource:       authResource,
		unitResource:       unitResource,
		stockResource:      stockResource,
		requestResource:    requestResource,
		itemResource:       itemResource,
		donationResource:   donationResource,
		allocationResource: allocationResource,
	}

	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/auth", api.authResource.router())
	r.Mount("/units", api.unitResource.router())
	r.Mount("/stocks", api.stockResource.router())
	r.Mount("/requests", api.requestResource.router())
	r.Mount("/items", api.itemResource.router())
	r.Mount("/donations", api.donationResource.router())
	r.Mount("/allocations", api.allocationResource.router())

	return r
}

// Cmd handles command from terminal.
func (api *API) Cmd(args []string) {
	api.authResource.cmd(args)
}

type contextKey struct {
	name string
}
