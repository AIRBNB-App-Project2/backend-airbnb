package booking

import (
	"gorm.io/datatypes"
)

type CreateBookingRequesFormat struct {
	User_uid      string
	Room_uid      string         `json:"room_uid" validate:"required"`
	Start_date    datatypes.Date `json:"start_date" validate:"required"`
	End_date      datatypes.Date `json:"end_date" validate:"required"`
	PaymentMethod string         `json:"payment_method" validate:"required"`
}
