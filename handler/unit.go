package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
)

// UnitResource holds unit data store information.
type UnitResource struct {
	*model.UnitDatastore
}

func (res *UnitResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Get("/", GetAllUnit(res.UnitDatastore))

	return r
}

// GetAllUnit gets all unit.
func GetAllUnit(repo interface {model.HasGetAllUnit}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		units, err := repo.GetAllUnit(r.Context())
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, units)
	}
}