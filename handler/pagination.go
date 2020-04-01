package handler

import (
	"context"
	"math"
	"net/http"
	"strconv"
)

// Paging struct
type Paging struct {
	Page int
	Size int
}

// PaginationCtx middleware is used to exctract page and size query param
func PaginationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}

		size, err := strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil || size < 1 {
			size = 10
		}
		
		paging := &Paging{
			Page: page,
			Size: size,
		}

		ctx := context.WithValue(r.Context(), PageCtxKey, paging)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Offset returns offset.
func (p *Paging) Offset() int {
	return (p.Page - 1) * p.Size
}

// Pages returns page.
func (p *Paging) Pages(totalCount int64) *Page {
	isLast := (int(totalCount) - (p.Size * p.Page)) < p.Size
	isFirst := p.Page == 1
	totalPages := int(math.Ceil(float64(totalCount) / float64(p.Size)))
	if totalPages == 0 {
		totalPages = 1
	}

	return &Page{
		Current: p.Page,
		Total:   totalPages,
		First:   isFirst,
		Last:    isLast,
	}
}

// Page struct
type Page struct {
	Current int  `boil:"current" json:"current"`
	Total   int  `boil:"total" json:"total"`
	First   bool `boil:"first" json:"first"`
	Last    bool `boil:"last" json:"last"`
}