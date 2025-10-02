// File: internal/driver/model.go

package driver

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Name          string  `gorm:"not null"`
	Email         string  `gorm:"unique;not null"`
	Password      string  `gorm:"not null"`
	LicenseNumber string  `gorm:"unique;not null"`
	IsAvailable   bool    `gorm:"default:true"`
}

type Claims struct {
	DriverID      uint   `json:"driver_id"`
	LicenseNumber string `json:"license_number"`
	jwt.RegisteredClaims
}