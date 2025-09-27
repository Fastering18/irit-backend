// File: internal/booking/repository.go

package booking

import "gorm.io/gorm"

type Repository interface {
	Create(booking *Booking) error
	FindByID(id uint) (*Booking, error)
	FindAvailable() ([]*Booking, error)
	FindByUserID(userID uint) ([]*Booking, error)
	FindByDriverID(driverID uint) ([]*Booking, error)
	FindActiveByDriverID(driverID uint) (*Booking, error)
	FindActiveByUserID(userID uint) (*Booking, error)
	Delete(booking *Booking) error
	Update(booking *Booking) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(booking *Booking) error {
	return r.db.Create(booking).Error
}

func (r *repository) FindByID(id uint) (*Booking, error) {
	var booking Booking
	err := r.db.Preload("User").Preload("Driver").First(&booking, id).Error
	return &booking, err
}

func (r *repository) FindAvailable() ([]*Booking, error) {
	var bookings []*Booking
	err := r.db.Preload("User").Preload("Driver").Where("status = ?", StatusMencariDriver).Find(&bookings).Error
	return bookings, err
}

func (r *repository) FindByUserID(userID uint) ([]*Booking, error) {
	var bookings []*Booking
	err := r.db.Preload("User").Preload("Driver").Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *repository) FindByDriverID(driverID uint) ([]*Booking, error) {
	var bookings []*Booking
	err := r.db.Preload("User").Preload("Driver").Where("driver_id = ?", driverID).Find(&bookings).Error
	return bookings, err
}

func (r *repository) FindActiveByDriverID(driverID uint) (*Booking, error) {
	var booking Booking
	err := r.db.Preload("User").Preload("Driver").Where("driver_id = ? AND status NOT IN (?, ?)", driverID, StatusSelesai, StatusDibatalkan).First(&booking).Error
	return &booking, err
}

func (r *repository) FindActiveByUserID(userID uint) (*Booking, error) {
	var booking Booking
	err := r.db.Preload("User").Preload("Driver").Where("user_id = ? AND status NOT IN (?, ?)", userID, StatusSelesai, StatusDibatalkan).First(&booking).Error
	return &booking, err
}

func (r *repository) Update(booking *Booking) error {
	return r.db.Save(booking).Error
}

func (r *repository) Delete(booking *Booking) error {
	return r.db.Delete(booking).Error
}