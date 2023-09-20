package model

type BookingAttention struct {
	ID        uint `json:"id"`
	UserID    uint `json:"userId"`
	BookingID uint `json:"bookingId"`
}
