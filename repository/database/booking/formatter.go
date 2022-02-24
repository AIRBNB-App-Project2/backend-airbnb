package booking

import (
	"time"
)

type BookingCreateResp struct {
	Booking_uid string    `json:"booking_uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Start_date  time.Time `json:"start_date"`
	End_date    time.Time `json:"end_date"`
	Price       int       `json:"price"`
	Days        int       `json:"days"`
	Price_total int       `json:"price_total"`
}
