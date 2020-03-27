package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource *AuthResource
	unitResource *UnitResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	unitDatastore := &model.UnitDatastore{DB: db}

	authResource := &AuthResource{UserDatastore: userDatastore}
	unitResource := &UnitResource{UnitDatastore: unitDatastore}

	api := &API{
		authResource: authResource,
		unitResource: unitResource,
	}
	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/auth", api.authResource.router())
	r.Mount("/units", api.unitResource.router())

	return r
}

type contextKey struct {
	name string
}