package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// API provides application resources and handlers.
type API struct {
	authResource *AuthResource
	requestResource *RequestResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) *API {
	userDatastore := &model.UserDatastore{DB: db}
	requestDatastore := &model.RequestDatastore{DB: db}

	authResource := &AuthResource{UserDatastore: userDatastore}
	requestResource := &RequestResource{
		requestDatastore: requestDatastore,
		userDatastore: userDatastore,
	}

	api := &API{
		authResource: authResource,
		requestResource: requestResource,
	}
	return api
}

// Router provides application routes.
func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/auth", api.authResource.router())
	r.Mount("/requests", api.requestResource.router())

	return r
}

type contextKey struct {
	name string
}