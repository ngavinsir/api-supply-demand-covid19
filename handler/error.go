package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

// common error message
var (
	ErrInvalidUserID    = errors.New("INVALID_USER_ID")
	ErrMissingReqFields = errors.New("MISSING_REQUEST_FIELDS")
	ErrInvalidRole		= errors.New("INVALID_ROLE")
)

// ErrResponse contains err, http_status_code, status_text, app_code, error_text
type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

// Render error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest to render invalid request error
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 200,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

// ErrUnauthorized to render unathorized error
func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 200,
		StatusText:     "Unauthorized",
		ErrorText:      err.Error(),
	}
}

// ErrRender to render common error
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 200,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}