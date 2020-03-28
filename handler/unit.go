package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// UnitResource holds unit data store information.
type UnitResource struct {
	*model.UnitDatastore
	*model.UserDatastore
}

func (res *UnitResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Get("/", GetAllUnit(res.UnitDatastore))

	r.Group(func(r chi.Router) {
		r.Use(UserCtx(res.UserDatastore))

		r.Post("/", CreateUnit(res.UnitDatastore))
		r.Delete("/{unitID}", DeleteUnit(res.UnitDatastore))
	})

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

// CreateUnit creates new unit.
func CreateUnit(repo interface {model.HasCreateUnit}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &CreateUnitRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		unit, err := repo.CreateUnit(r.Context(), data.Name)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, unit)
	}
}

// DeleteUnit deletes unit by given id.
func DeleteUnit(repo interface {model.HasDeleteUnit}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		unitID := chi.URLParam(r, "unitID")
		if unitID == "" {
			render.Render(w, r, ErrInvalidRequest(ErrMissingReqFields))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		err := repo.DeleteUnit(r.Context(), unitID)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
}

// CreateUnitRequest struct
type CreateUnitRequest struct {
	*models.Unit
}

// Bind CreateUnitRequest (name) [Required]
func (req *CreateUnitRequest) Bind(r *http.Request) error {
	if req.Unit == nil || req.Name == "" {
		return ErrMissingReqFields
	}

	return nil
}