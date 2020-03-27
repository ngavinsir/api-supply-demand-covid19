package handler

import (
	"net/http"
	"time"

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
	r.Get("/", GetAllRequest(res.requestDatastore))
	
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
		requests, err := repo.GetAllRequest(r.Context())
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		var requestResponses []*RequestResponse
		for _, r := range *requests {
			donationApplicant := &LoginResponse{
				ID: r.R.DonationApplicant.ID,
				Email: r.R.DonationApplicant.Email,
				Name: r.R.DonationApplicant.Name,
				Role: r.R.DonationApplicant.Role,
				ContactNumber: r.R.DonationApplicant.ContactNumber.String,
				ContactPerson: r.R.DonationApplicant.ContactPerson.String,
			}			

			requestResponses = append(requestResponses, &RequestResponse{
				ID: r.ID,
				Date: r.Date,
				IsFulfilled: r.IsFulfilled,
				RequestItems: &r.R.RequestItems,
				DonationApplicant: donationApplicant,
			})
		}

		render.JSON(w, r, requestResponses)
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

// RequestResponse struct
type RequestResponse struct {
	ID string `json:"id"`
	Date time.Time `json:"date"`
	IsFulfilled	bool `json:"isFulfilled"`
	DonationApplicant *LoginResponse `json:"donationApplicant"`
	RequestItems *models.RequestItemSlice `json:"requestItems"`
}