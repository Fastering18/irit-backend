// File: internal/user/error.go

package user

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
	ErrEmailExists        = NewResponseError(http.StatusBadRequest, "email already exists")
	ErrIdentifierExists   = NewResponseError(http.StatusBadRequest, "identifier (NRP/NIP) already exists")
	ErrInvalidCredentials = NewResponseError(http.StatusUnauthorized, "invalid email or password")
	ErrUserNotFound       = NewResponseError(http.StatusNotFound, "user not found")
	ErrInternalServer     = NewResponseError(http.StatusInternalServerError, "an internal server error occurred")
)