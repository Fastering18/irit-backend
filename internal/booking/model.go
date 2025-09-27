// File: internal/booking/model.go

package booking

import (
	"irit-backend/internal/driver"
	"irit-backend/internal/user"

	"gorm.io/gorm"
)

type BookingStatus string
type VehicleType string

const (
	StatusMencariDriver BookingStatus = "mencari driver"
	StatusMenungguDriver BookingStatus = "menunggu driver"
	StatusDiperjalanan  BookingStatus = "diperjalanan"
	StatusSelesai       BookingStatus = "selesai"
	StatusDibatalkan    BookingStatus = "dibatalkan"
)

const (
	TypeShuttle      VehicleType = "shuttle"
	TypeSepeda       VehicleType = "sepeda"
	TypeCarpool      VehicleType = "carpool"
	TypeSepedaMotor  VehicleType = "sepeda motor"
)

type Booking struct {
	gorm.Model
	UserID          uint
	DriverID        *uint
	PickupLocation  string `gorm:"not null"`
	DropoffLocation string `gorm:"not null"`
	Status          BookingStatus `gorm:"type:varchar(20);not null"`
	VehicleType     VehicleType   `gorm:"type:varchar(20);not null"`
	User            user.User     `gorm:"foreignKey:UserID"`
	Driver          driver.Driver `gorm:"foreignKey:DriverID"`
}