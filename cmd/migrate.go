package main

import (
	"log"
	"irit-backend/internal/user"
	"irit-backend/internal/driver"
	"irit-backend/internal/booking"
	"irit-backend/pkg/config"
	"irit-backend/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Tidak bisa memuat konfigurasi: %v", err)
	}

	db, err := database.Initialize(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Gagal menginisialisasi database: %v", err)
	}
	log.Println("Koneksi database berhasil.")

	err = db.AutoMigrate(&user.User{}, &driver.Driver{}, &booking.Booking{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	log.Println("Migrasi database berhasil.")
}