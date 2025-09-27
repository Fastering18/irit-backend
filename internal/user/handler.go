// File: internal/user/handler.go

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService Service
}

func NewHandler(userService Service) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"identifier": user.Identifier,
			"role":       user.Role,
		},
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.LoginUser(input)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Account(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get user from token"})
		return
	}

	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format in token"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"identifier": user.Identifier,
		"address":    user.Address,
		"role":       user.Role,
	})
}

func RegisterRoutes(router *gin.Engine, service Service, authMiddleware gin.HandlerFunc) {
	handler := NewHandler(service)

	api := router.Group("/api/v1/users")
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