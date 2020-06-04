package handler

import (
	"net/http"
	"time"

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
	*model.RequestDatastore
}

func (res *AllocationResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Use(UserCtx(res.UserDatastore))
	r.Post("/", CreateAllocation(res))
	r.With(PaginationCtx).Get("/", GetAllAllocations(res))

	return r
}

// CreateAllocation creates new allocation.
func CreateAllocation(
	repo interface {
		model.HasCreateAllocation
		model.HasIsStockAvailable
		model.HasCreateOrUpdateStock
		model.HasGetRequest
	},
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleAdmin && user.Role != model.RoleDonator {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		data := &CreateAllocationRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		allocation, err := repo.CreateAllocation(
			r.Context(),
			&models.Allocation{
				AllocatorID: user.ID,
				PhotoURL:    null.StringFrom(data.PhotoURL),
				RequestID:   data.RequestID,
				Date:        data.Date,
			},
			data.AllocationItems,
			repo,
			repo,
		)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, allocation)
	}
}

// GetAllAllocations gets all allocations.
func GetAllAllocations(
	repo interface {
		model.HasGetAllAllocations
		model.HasGetTotalAllocationCount
		model.HasGetRequest
	},
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		paging, _ := r.Context().Value(PageCtxKey).(*Paging)

		allocationData, err := repo.GetAllAllocations(r.Context(), paging.Offset(), paging.Size, repo)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		totalAllocationCount, err := repo.GetTotalAllocationCount(r.Context())

		allocationDataPage := &AllocationDataPage{
			Data:  allocationData,
			Pages: paging.Pages(totalAllocationCount),
		}

		render.JSON(w, r, allocationDataPage)
	}
}

// CreateAllocationRequest struct
type CreateAllocationRequest struct {
	RequestID       string                     `json:"requestID"`
	Date            time.Time                  `json:"date"`
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

// AllocationDataPage struct
type AllocationDataPage struct {
	Data  []*model.AllocationData `boil:"data" json:"data"`
	Pages *Page                   `boil:"pages" json:"pages"`
}
