// File: internal/booking/dto.go

package booking

type CreateBookingInput struct {
	PickupLocation  string      `json:"pickup_location" binding:"required"`
	DropoffLocation string      `json:"dropoff_location" binding:"required"`
	VehicleType     VehicleType `json:"vehicle_type" binding:"required"`
}

type AcceptBookingInput struct {
	BookingID uint `json:"booking_id" binding:"required"`
}

type DriverUpdateStatusInput struct {
	BookingID uint          `json:"booking_id" binding:"required"`
	Status    BookingStatus `json:"status" binding:"required"`
}