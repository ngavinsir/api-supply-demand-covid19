package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource *AuthResource
	unitResource *UnitResource
	requestResource *RequestResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	unitDatastore := &model.UnitDatastore{DB: db}
	requestDatastore := &model.RequestDatastore{DB: db}

	authResource := &AuthResource{UserDatastore: userDatastore}
	unitResource := &UnitResource{
		UnitDatastore: unitDatastore,
		UserDatastore: userDatastore,
	}
	requestResource := &RequestResource{
		requestDatastore: requestDatastore,
		userDatastore: userDatastore,
	}

	api := &API{
		authResource: authResource,
		unitResource: unitResource,
		requestResource: requestResource,
	}

	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/auth", api.authResource.router())
	r.Mount("/units", api.unitResource.router())
	r.Mount("/requests", api.requestResource.router())

	return r
}

type contextKey struct {
	name string
}