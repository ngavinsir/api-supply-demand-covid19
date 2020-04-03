package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// DonationResource struct
type DonationResource struct {
	*model.DonationDataStore
	*model.UserDatastore
	*model.StockDataStore
}

func (res *DonationResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Use(UserCtx(res.UserDatastore))

	r.Post("/", CreateOrUpdateDonation(res.DonationDataStore, model.CreateAction))
	r.Put("/", CreateOrUpdateDonation(res.DonationDataStore, model.UpdateAction))
	r.Get("/{donationID}", GetDonation(res.DonationDataStore))
	r.Put("/{donationID}/accept", AcceptDonation(res.DonationDataStore, res.StockDataStore))

	return r
}

// CreateOrUpdateDonation return donations
func CreateOrUpdateDonation(repo interface {
	model.HasCreateOrUpdate
}, action string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &CreateDonationRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleDonator && user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		request, err := repo.CreateOrUpdateDonation(r.Context(), data.DonationItems, user.ID, action)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, request)
	}
}

// AcceptDonation accepts donation by given id.
func AcceptDonation(
	donationRepo interface{ model.HasAcceptDonation },
	stockRepo interface{ model.HasCreateOrUpdateStock },
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		donationID := chi.URLParam(r, "donationID")
		if donationID == "" {
			render.Render(w, r, ErrInvalidRequest(ErrMissingReqFields))
			return
		}

		if err := donationRepo.AcceptDonation(r.Context(), donationID, stockRepo); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
}

// GetDonation handles get donation detail given id.
func GetDonation(donationRepo interface{ model.HasGetDonation }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		donationID := chi.URLParam(r, "donationID")
		if donationID == "" {
			render.Render(w, r, ErrInvalidRequest(ErrMissingReqFields))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleDonator && user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		response, err := donationRepo.GetDonation(r.Context(), donationID, user)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, response)
	}
}

// CreateDonationRequest struct
type CreateDonationRequest struct {
	DonationItems []*models.DonationItem `json:"donationItems"`
}

// Bind RegisterRequest ([]RequestItem) [Required]
func (req *CreateDonationRequest) Bind(r *http.Request) error {
	if req.DonationItems == nil || len(req.DonationItems) == 0 {
		return ErrMissingReqFields
	}

	return nil
}
