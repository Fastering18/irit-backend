// File: internal/driver/error.go

package driver

import "net/http"

type ResponseError struct {
	Code    int
	Message string
}

func (e *ResponseError) Error() string {
	return e.Message
}

func NewResponseError(code int, message string) *ResponseError {
	return &ResponseError{Code: code, Message: message}
}

var (
	ErrEmailExists        = NewResponseError(http.StatusConflict, "email already exists")
	ErrLicenseExists      = NewResponseError(http.StatusConflict, "license number already exists")
	ErrInvalidCredentials = NewResponseError(http.StatusUnauthorized, "invalid email or password")
	ErrDriverNotFound     = NewResponseError(http.StatusNotFound, "driver not found")
	ErrInternalServer     = NewResponseError(http.StatusInternalServerError, "an internal server error occurred")
)