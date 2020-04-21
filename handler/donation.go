package handler

import (
	"net/http"

	"github.com/ericlagergren/decimal"
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

	r.Get("/{donationID}", GetDonation(res.DonationDataStore))
	r.With(PaginationCtx).Get("/", GetAllDonations(res.DonationDataStore))
	r.With(PaginationCtx).Get("/user/{userID}", GetUserDonations(res.DonationDataStore))

	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Use(UserCtx(res.UserDatastore))

		r.Post("/", CreateOrUpdateDonation(res.DonationDataStore, model.CreateAction))
		r.Put("/{donationID}", UpdateDonation(res.DonationDataStore))
		r.Put("/{donationID}/accept", AcceptDonation(res.DonationDataStore, res.StockDataStore))
	})

	return r
}

// CreateOrUpdateDonation return donations
func CreateOrUpdateDonation(repo interface {
	model.HasCreateOrUpdateDonation
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

// UpdateDonation handles donation update
func UpdateDonation(repo interface{ model.HasUpdateDonation }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		donationID := chi.URLParam(r, "donationID")

		data := &UpdateDonationRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleDonator && user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		request, err := repo.UpdateDonation(r.Context(), data.DonationItems, donationID)
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

		response, err := donationRepo.GetDonation(r.Context(), donationID)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, response)
	}
}

// GetAllDonations gets all requests.
func GetAllDonations(
	repo interface {
		model.HasGetAllDonations
		model.HasGetTotalDonationCount
	},
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paging, _ := r.Context().Value(PageCtxKey).(*Paging)

		donationData, err := repo.GetAllDonations(r.Context(), paging.Offset(), paging.Size)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		totalDonationCount, err := repo.GetTotalDonationCount(r.Context())

		donationDataPage := &DonationDataPage{
			Data:  donationData,
			Pages: paging.Pages(totalDonationCount),
		}

		render.JSON(w, r, donationDataPage)
	}
}

// GetUserDonations gets all user donations.
func GetUserDonations(
	repo interface {
		model.HasGetUserDonations
		model.HasGetTotalUserDonationCount
	},
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paging, _ := r.Context().Value(PageCtxKey).(*Paging)

		userID := chi.URLParam(r, "userID")
		if userID == "" {
			render.Render(w, r, ErrInvalidRequest(ErrMissingReqFields))
			return
		}

		donationData, err := repo.GetUserDonations(r.Context(), userID, paging.Offset(), paging.Size)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		totalDonationCount, err := repo.GetTotalUserDonationCount(r.Context(), userID)

		donationDataPage := &DonationDataPage{
			Data:  donationData,
			Pages: paging.Pages(totalDonationCount),
		}

		render.JSON(w, r, donationDataPage)
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

// UpdateDonationRequest struct
type UpdateDonationRequest struct {
	DonationItems models.DonationItemSlice `json:"donationItems"`
}

// Bind UpdateDonation ([]DonationItem) [Required]
func (req *UpdateDonationRequest) Bind(r *http.Request) error {
	if req.DonationItems == nil || len(req.DonationItems) == 0 {
		return ErrMissingReqFields
	}
	zeroBig := &decimal.Big{}
	for _, donationItem := range req.DonationItems {
		if donationItem.ItemID == "" || donationItem.ID == "" || donationItem.UnitID == "" || donationItem.Quantity.Big.Cmp(zeroBig) == 0 {
			return ErrMissingReqFields
		}
	}

	return nil
}

// DonationDataPage struct
type DonationDataPage struct {
	Data  []*model.DonationData `boil:"data" json:"data"`
	Pages *Page                 `boil:"pages" json:"pages"`
}
