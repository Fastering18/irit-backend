// File: internal/driver/handler.go

package driver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	driverService Service
}

func NewHandler(driverService Service) *Handler {
	return &Handler{driverService: driverService}
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterDriverInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	driver, err := h.driverService.RegisterDriver(input)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Driver registered successfully",
		"data": gin.H{
			"id":            driver.ID,
			"name":          driver.Name,
			"email":         driver.Email,
			"licenseNumber": driver.LicenseNumber,
		},
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginDriverInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	token, err := h.driverService.LoginDriver(input)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Account(c *gin.Context) {
	driverID, exists := c.Get("driverID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get driver from token"})
		return
	}

	id, ok := driverID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid driver ID format in token"})
		return
	}

	user, err := h.driverService.GetDriverByID(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             user.ID,
		"name":           user.Name,
		"email":          user.Email,
		"license_number": user.LicenseNumber,
	})
}

func RegisterRoutes(router *gin.Engine, service Service, authMiddleware gin.HandlerFunc) {
	handler := NewHandler(service)

	api := router.Group("/api/v1/drivers")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)

		authorized := api.Group("/")
		authorized.Use(authMiddleware)
		{
			authorized.GET("/account", handler.Account)
		}
	}
}
