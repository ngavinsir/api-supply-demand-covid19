package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// ItemResource holds item data store information.
type ItemResource struct {
	*model.ItemDatastore
	*model.UserDatastore
}

func (res *ItemResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Get("/", GetAllItem(res.ItemDatastore))
	r.With(UserCtx(res.UserDatastore)).Post("/", CreateItem(res.ItemDatastore))

	return r
}

// GetAllItem gets all item.
func GetAllItem(repo interface {model.HasGetAllItem}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := repo.GetAllItem(r.Context())
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, items)
	}
}

// CreateItem creates new item.
func CreateItem(repo interface {model.HasCreateItem}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &CreateItemRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		item, err := repo.CreateItem(r.Context(), data.Name)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, item)
	}
}

// CreateItemRequest struct
type CreateItemRequest struct {
	*models.Item
}

// Bind CreateItemRequest (name) [Required]
func (req *CreateItemRequest) Bind(r *http.Request) error {
	if req.Item == nil || req.Name == "" {
		return ErrMissingReqFields
	}

	return nil
}