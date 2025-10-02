// File: internal/booking/service.go

package booking

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateBooking(userID uint, input CreateBookingInput) (*Booking, *ResponseError)
	AcceptBooking(driverID, bookingID uint) (*Booking, *ResponseError)
	UpdateBookingStatus(driverID, bookingID uint, newStatus BookingStatus) *ResponseError
	GetBookingById(bookingID uint) (*Booking, *ResponseError)
	GetBookingDetails(userID, driverID uint, bookingID uint, isDriver bool) (*Booking, error)
	GetBookingHistoryForUser(userID uint) ([]*Booking, error)
	GetBookingHistoryForDriver(driverID uint) ([]*Booking, error)
	FindOpenBookings() ([]*Booking, error)
	GetDistanceToDriver(requesterID, bookingID uint, isDriver bool) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateBooking(userID uint, input CreateBookingInput) (*Booking, *ResponseError) {
	existingBooking, err := s.repo.FindActiveByUserID(userID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInternalServer
	}

	if existingBooking.Status == StatusMencariDriver {
		if err := s.repo.Delete(existingBooking); err != nil {
			return nil, ErrInternalServer
		}
	} else if existingBooking.ID != 0 {
		return nil, ErrUserOnActiveBooking
	}

	newBooking := &Booking{
		UserID:          userID,
		PickupLocation:  input.PickupLocation,
		DropoffLocation: input.DropoffLocation,
		VehicleType:     input.VehicleType,
		Status:          StatusMencariDriver,
	}

	if err := s.repo.Create(newBooking); err != nil {
		return nil, ErrInternalServer
	}

	createdBooking, err := s.repo.FindByID(newBooking.ID)
	if err != nil {
		return nil, ErrInternalServer
	}

	return createdBooking, nil
}

func (s *service) AcceptBooking(driverID, bookingID uint) (*Booking, *ResponseError) {
	_, err := s.repo.FindActiveByDriverID(driverID)
	if err == nil {
		return nil, ErrDriverBusy
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInternalServer
	}

	booking, err := s.repo.FindByID(bookingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookingNotFound
		}
		return nil, ErrInternalServer
	}

	if booking.DriverID != nil {
		return nil, ErrBookingNotAvailable
	}

	booking.DriverID = &driverID
	booking.Status = StatusMenungguDriver

	if err := s.repo.Update(booking); err != nil {
		return nil, ErrInternalServer
	}
	return booking, nil
}

func (s *service) GetBookingDetails(userID, driverID, bookingID uint, isDriver bool) (*Booking, error) {
	booking, err := s.repo.FindByID(bookingID)
	if err != nil {
		return nil, ErrBookingNotFound
	}

	if isDriver {
		if booking.DriverID == nil || *booking.DriverID != driverID {
			return nil, ErrForbidden
		}
	} else {
		if booking.UserID != userID {
			return nil, ErrForbidden
		}
	}
	return booking, nil
}

func (s *service) GetBookingHistoryForDriver(driverID uint) ([]*Booking, error) {
	return s.repo.FindByDriverID(driverID)
}

func (s *service) UpdateBookingStatus(driverID, bookingID uint, newStatus BookingStatus) *ResponseError {
	booking, err := s.repo.FindByID(bookingID)
	if err != nil {
		return ErrBookingNotFound
	}

	if booking.DriverID == nil || *booking.DriverID != driverID {
		return ErrForbidden
	}

	booking.Status = newStatus
	if err := s.repo.Update(booking); err != nil {
		return ErrInternalServer
	}
	return nil
}

func (s *service) GetBookingById(bookingID uint) (*Booking, *ResponseError) {
	booking, err := s.repo.FindByID(bookingID)
	if err != nil {
		return nil, ErrBookingNotFound
	}
	// if booking.UserID != userID {
	// 	return nil, ErrForbidden
	// }
	return booking, nil
}

func (s *service) GetBookingHistoryForUser(userID uint) ([]*Booking, error) {
	return s.repo.FindByUserID(userID)
}

func (s *service) FindOpenBookings() ([]*Booking, error) {
	return s.repo.FindAvailable()
}

func (s *service) GetDistanceToDriver(requesterID, bookingID uint, isDriver bool) (int, error) {
	booking, err := s.repo.FindByID(bookingID)
	if err != nil {
		return 0, ErrBookingNotFound
	}

	if isDriver {
		if booking.DriverID == nil || *booking.DriverID != requesterID {
			return 0, ErrForbidden
		}
	} else {
		if booking.UserID != requesterID {
			return 0, ErrForbidden
		}
	}

	switch booking.Status {
	case StatusMencariDriver:
		return 0, ErrDriverNotFound
	case StatusMenungguDriver:
		return 100, nil
	case StatusDiperjalanan:
		return 50, nil
	case StatusSelesai:
		return 0, nil
	default:
		return 0, ErrInternalServer
	}
}
