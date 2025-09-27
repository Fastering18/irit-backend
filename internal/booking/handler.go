// File: internal/booking/handler.go

package booking

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	bookingService Service
}

func NewHandler(bookingService Service) *Handler {
	return &Handler{bookingService: bookingService}
}

func (h *Handler) CreateBooking(c *gin.Context) {
	userID, _ := c.Get("userID")
	var input CreateBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	booking, err := h.bookingService.CreateBooking(userID.(uint), input)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully", "data": booking})
}

func (h *Handler) GetBookingDetails(c *gin.Context) {
	bookingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	userID, userExists := c.Get("userID")
	driverID, driverExists := c.Get("driverID")
	
	var booking *Booking
	
	if userExists {
		booking, err = h.bookingService.GetBookingDetails(userID.(uint), 0, uint(bookingID), false)
	} else if driverExists {
		booking, err = h.bookingService.GetBookingDetails(0, driverID.(uint), uint(bookingID), true)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	
	if err != nil {
		if resErr, ok := err.(*ResponseError); ok {
			c.JSON(resErr.Code, gin.H{"error": resErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "an unexpected error occurred"})
		}
		return
	}
	c.JSON(http.StatusOK, booking)
}

func (h *Handler) GetHistoryList(c *gin.Context) {
	var bookings []*Booking
	var err error

	if userID, exists := c.Get("userID"); exists {
		bookings, err = h.bookingService.GetBookingHistoryForUser(userID.(uint))
	} else if driverID, exists := c.Get("driverID"); exists {
		bookings, err = h.bookingService.GetBookingHistoryForDriver(driverID.(uint))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve history"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (h *Handler) FindOpenBookings(c *gin.Context) {
	bookings, err := h.bookingService.FindOpenBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve open bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (h *Handler) AcceptBooking(c *gin.Context) {
	driverID, _ := c.Get("driverID")
	var input AcceptBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format, requires 'booking_id'"})
		return
	}

	_, err := h.bookingService.AcceptBooking(driverID.(uint), input.BookingID)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking accepted successfully"})
}

func (h *Handler) UpdateStatusByDriver(c *gin.Context) {
	driverID, _ := c.Get("driverID")
	var input DriverUpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format, requires 'booking_id' and 'status'"})
		return
	}

	err := h.bookingService.UpdateBookingStatus(driverID.(uint), input.BookingID, input.Status)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking status updated successfully"})
}

func (h *Handler) GetDistance(c *gin.Context) {
	bookingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	var distance int
	if userID, exists := c.Get("userID"); exists {
		distance, err = h.bookingService.GetDistanceToDriver(userID.(uint), uint(bookingID), false)
	} else if driverID, exists := c.Get("driverID"); exists {
		distance, err = h.bookingService.GetDistanceToDriver(driverID.(uint), uint(bookingID), true)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	if err != nil {
		if resErr, ok := err.(*ResponseError); ok {
			c.JSON(resErr.Code, gin.H{"error": resErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "an unexpected error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"distance_in_meters": distance})
}

func RegisterRoutes(router *gin.Engine, service Service, userAuth gin.HandlerFunc, driverAuth gin.HandlerFunc, combinedAuth gin.HandlerFunc) {
	handler := NewHandler(service)
	
	api := router.Group("/api/v1/bookings")
	
	// User routes
	api.POST("/book", userAuth, handler.CreateBooking)
	
	// Driver routes
	api.GET("/orders", driverAuth, handler.FindOpenBookings)
	api.POST("/orders", driverAuth, handler.AcceptBooking)
	api.POST("/update", driverAuth, handler.UpdateStatusByDriver)
	
	// Combined routes
	api.GET("/:id", combinedAuth, handler.GetBookingDetails)
	api.GET("/:id/distance", combinedAuth, handler.GetDistance)
	api.GET("/list", combinedAuth, handler.GetHistoryList)
}