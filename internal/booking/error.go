// File: internal/booking/error.go

package booking

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
	ErrBookingNotFound     = NewResponseError(http.StatusNotFound, "Booking tidak ditemukan")
	ErrBookingNotAvailable = NewResponseError(http.StatusConflict, "Booking sedang diambil oleh driver lain")
	ErrDriverBusy          = NewResponseError(http.StatusConflict, "Driver sedang mengantar active booking")
	ErrDriverNotFound      = NewResponseError(http.StatusNotFound, "Driver belum ditemukan untuk booking ini")
	ErrForbidden           = NewResponseError(http.StatusForbidden, "Anda tidak boleh melakukan aksi ini")
	ErrInternalServer      = NewResponseError(http.StatusInternalServerError, "an internal server error occurred")
	ErrUserOnActiveBooking = NewResponseError(http.StatusBadRequest, "Anda sudah memiliki pesanan yang sedang aktif")
)