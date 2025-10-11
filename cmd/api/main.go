// File: cmd/api/main.go

package main

import (
	//"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"irit-backend/internal/auth"
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

	// User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository, cfg.JWT.Secret)

	// Driver
	driverRepository := driver.NewRepository(db)
	driverService := driver.NewService(driverRepository, cfg.JWT.Secret)

	// Booking
	bookingRepository := booking.NewRepository(db)
	bookingService := booking.NewService(bookingRepository)

	userAuthMiddleware := auth.UserMiddleware(cfg.JWT.Secret)
	driverAuthMiddleware := auth.DriverMiddleware(cfg.JWT.Secret)
	combinedAuthMiddleware := auth.CombinedAuthMiddleware(cfg.JWT.Secret)

	router := gin.Default()
	user.RegisterRoutes(router, userService, userAuthMiddleware)
	driver.RegisterRoutes(router, driverService, driverAuthMiddleware)
	booking.RegisterRoutes(router, bookingService, userAuthMiddleware, driverAuthMiddleware, combinedAuthMiddleware)
	log.Println("Routes telah didaftarkan.")

	serverAddress := cfg.Server.URL
	log.Printf("Server berjalan di %s", serverAddress)

	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Gagal menjalankan server Gin: %v", err.Error())
	}
}
