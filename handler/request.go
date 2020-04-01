package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// RequestResource holds request data store.
type RequestResource struct {
	requestDatastore *model.RequestDatastore
	userDatastore *model.UserDatastore
}

func (res *RequestResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Use(UserCtx(res.userDatastore))
	r.Post("/", CreateRequest(res.requestDatastore))
	r.With(PaginationCtx).Get("/", GetAllRequest(res.requestDatastore))
	
	return r
}

// CreateRequest handles request creation
func CreateRequest(repo interface {model.HasCreateRequest}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &CreateRequestRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleApplicant && user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		request, err := repo.CreateRequest(r.Context(), *data.RequestItems, user.ID)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, request)
	}
}

// GetAllRequest gets all requests.
func GetAllRequest(repo interface { model.HasGetAllRequest }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paging, _ := r.Context().Value(PageCtxKey).(*Paging)

		requestData, totalCount, err := repo.GetAllRequest(r.Context(), paging.Offset(), paging.Size)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		requestDataPage := &RequestDataPage{
			Data: requestData,
			Pages: paging.Pages(totalCount),
		}

		render.JSON(w, r, requestDataPage)
	}
}

// CreateRequestRequest struct
type CreateRequestRequest struct {
	RequestItems *models.RequestItemSlice `json:"requestItems"`
}

// Bind RegisterRequest ([]RequestItem) [Required]
func (req *CreateRequestRequest) Bind(r *http.Request) error {
	if req.RequestItems == nil || len(*req.RequestItems) == 0 {
		return ErrMissingReqFields
	}

	return nil
}

// RequestDataPage struct
type RequestDataPage struct {
	Data  []*model.RequestData `boil:"data" json:"data"`
	Pages *Page        `boil:"pages" json:"pages"`
}