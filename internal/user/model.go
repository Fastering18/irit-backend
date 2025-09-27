// File: internal/user/model.go

package user

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Role string
const (
	RoleMahasiswa Role = "mahasiswa"
	RoleDosen     Role = "dosen"
	RoleAdmin     Role = "admin"
	
)

type User struct {
	gorm.Model    
	Name          string `gorm:"not null"`
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"not null"` 
	Identifier    string `gorm:"unique;not null"` 
	Address       string
	Role          Role   `gorm:"type:varchar(20);not null"`
}

type Claims struct {
	UserID     uint   `json:"user_id"`
	Identifier string `json:"identifier"`
	Role       Role   `json:"role"`
	jwt.RegisteredClaims
}