// File: internal/driver/service.go

package driver

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	RegisterDriver(input RegisterDriverInput) (*Driver,  *ResponseError)
	LoginDriver(input LoginDriverInput) (string,  *ResponseError)
	UpdateLocation(driverID uint, input UpdateLocationInput)  *ResponseError
	GetDriverByID(id uint) (*Driver,  *ResponseError)
}

type service struct {
	repo      Repository
	jwtSecret []byte
}

func NewService(repo Repository, jwtSecret string) Service {
	return &service{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *service) RegisterDriver(input RegisterDriverInput) (*Driver, *ResponseError) {
	_, err := s.repo.FindByEmail(input.Email)
	if err == nil {
		return nil, ErrEmailExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInternalServer
	}

	_, err = s.repo.FindByLicenseNumber(input.LicenseNumber)
	if err == nil {
		return nil, ErrLicenseExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInternalServer
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternalServer
	}

	newDriver := &Driver{
		Name:          input.Name,
		Email:         input.Email,
		Password:      string(hashedPassword),
		LicenseNumber: input.LicenseNumber,
	}

	if err := s.repo.Create(newDriver); err != nil {
		return nil, ErrInternalServer
	}
	return newDriver, nil
}

func (s *service) LoginDriver(input LoginDriverInput) (string, *ResponseError) {
	driver, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", ErrInternalServer
	}

	if err := bcrypt.CompareHashAndPassword([]byte(driver.Password), []byte(input.Password)); err != nil {
		return "", ErrInvalidCredentials
	}

	claims := &Claims{
		DriverID:      driver.ID,
		LicenseNumber: driver.LicenseNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", ErrInternalServer
	}
	return signedToken, nil
}

func (s *service) UpdateLocation(driverID uint, input UpdateLocationInput) *ResponseError {
	driver, err := s.repo.FindByID(driverID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrDriverNotFound
		}
		return ErrInternalServer
	}

	driver.Latitude = input.Latitude
	driver.Longitude = input.Longitude

	if err := s.repo.Update(driver); err != nil {
		return ErrInternalServer
	}
	return nil
}

func (s *service) GetDriverByID(id uint) (*Driver, *ResponseError) {
	driver, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDriverNotFound
		}
		return nil, ErrInternalServer
	}
	return driver, nil
}