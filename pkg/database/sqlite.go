// File: pkg/database/sqlite.go

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{}) 
	if err != nil {
		return nil, err
	}
	return db, nil
}