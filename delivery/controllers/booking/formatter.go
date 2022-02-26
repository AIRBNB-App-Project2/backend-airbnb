package booking

import (
	"be/entities"
)

type CreateBookingRequesFormat struct {
	User_uid   string
	Room_uid   string `json:"room_uid" validate:"required"`
	Start_date string `json:"start_date" validate:"required"`
	End_date   string `json:"end_date" validate:"required"`
}

type GetBookingResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Booking `json:"data"`
}
