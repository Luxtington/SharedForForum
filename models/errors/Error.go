package errors

import "errors"

//comm

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}

func NewErrorResponse(err string, code int) *ErrorResponse {
	return &ErrorResponse{
		Error:      err,
		StatusCode: code,
	}
}
