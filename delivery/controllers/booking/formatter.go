package booking

import (
	"be/entities"
)

type CreateBookingRequesFormat struct {
	User_uid      string
	Room_uid      string `json:"room_uid"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
	PaymentMethod string `json:"paymentmethod"`
	Status        string `json:"status"`
}

type GetBookingResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Booking `json:"data"`
}
