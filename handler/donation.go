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
	r.Post("/", CreateDonation(res.DonationDataStore))
	r.Put("/{donationID}/accept", AcceptDonation(res.DonationDataStore, res.StockDataStore))

	return r
}

// CreateDonation return donations
func CreateDonation(repo interface {
	model.HasCreateDonation
}) http.HandlerFunc {
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

		request, err := repo.CreateDonation(r.Context(), data.DonationItems, user.ID)
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
