// File: internal/driver/repository.go

package driver

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(driver *Driver) error
	FindByEmail(email string) (*Driver, error)
	FindByLicenseNumber(licensenumber string) (*Driver, error)
	FindByID(id uint) (*Driver, error)
	Update(driver *Driver) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(driver *Driver) error {
	return r.db.Create(driver).Error
}

func (r *repository) FindByEmail(email string) (*Driver, error) {
	var driver Driver
	if err := r.db.Where("email = ?", email).First(&driver).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *repository) FindByLicenseNumber(licensenumber string) (*Driver, error) {
	var driver Driver
	if err := r.db.Where("license_number = ?", licensenumber).First(&driver).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *repository) FindByID(id uint) (*Driver, error) {
	var driver Driver
	if err := r.db.Where("id = ?", id).First(&driver).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *repository) Update(driver *Driver) error {
	return r.db.Save(driver).Error
}