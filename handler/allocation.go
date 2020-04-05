package handler

import (
	"net/http"

	"github.com/ericlagergren/decimal"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/volatiletech/null"
)

// AllocationResource struct
type AllocationResource struct {
	*model.AllocationDatastore
	*model.StockDataStore
	*model.UserDatastore
}

func (res *AllocationResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Use(UserCtx(res.UserDatastore))
	r.Post("/", CreateAllocation(res.AllocationDatastore, res.StockDataStore))

	return r
}

// CreateAllocation creates new allocation.
func CreateAllocation(
	allocationRepo interface{ model.HasCreateAllocation },
	stockRepo interface {
		model.HasIsStockAvailable
		model.HasCreateOrUpdateStock
	},
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		data := &CreateAllocationRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		allocation, err := allocationRepo.CreateAllocation(
			r.Context(),
			&models.Allocation{
				AdminID:   user.ID,
				PhotoURL:  null.StringFrom(data.PhotoURL),
				RequestID: data.RequestID,
			},
			data.AllocationItems,
			stockRepo,
		)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, allocation)
	}
}

// CreateAllocationRequest struct
type CreateAllocationRequest struct {
	RequestID       string                     `json:"requestID"`
	PhotoURL        string                     `json:"photoURL"`
	AllocationItems models.AllocationItemSlice `json:"items"`
}

// Bind CreateAllocationRequest ([]AllocationItem) [Required]
func (req *CreateAllocationRequest) Bind(r *http.Request) error {
	if req.AllocationItems == nil || len(req.AllocationItems) == 0 || req.RequestID == "" {
		return ErrMissingReqFields
	}
	zeroBig := &decimal.Big{}
	for _, allocationItem := range req.AllocationItems {
		if allocationItem.ItemID == "" || allocationItem.UnitID == "" || allocationItem.Quantity.Big.Cmp(zeroBig) == 0 {
			return ErrMissingReqFields
		}
	}

	return nil
}
