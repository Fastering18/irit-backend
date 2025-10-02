// File: pkg/database/mysql.go

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Initialize membuat dan mengembalikan koneksi database GORM.
func Initialize(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err != nil {
		return nil, err
	}
	return db, nil
}