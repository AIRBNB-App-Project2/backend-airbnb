package booking

import "gorm.io/datatypes"

type BookingCreateResp struct {
	Booking_uid string         `json:"booking_uid"`
	Name        string         `name:"name"`
	Description string         `json:"description"`
	Start_date  datatypes.Date `json:"start_date"`
	End_date    datatypes.Date `json:"end_date"`
	Price       int            `json:"price"`
	Days        int            `json:"days"`
	Price_total int            `json:"price_total"`
}

type BookingGetByIdResp struct {
	Booking_uid string `json:"booking_uid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	Price       int    `json:"price"`
	Days        int    `json:"days"`
	Price_total int    `json:"price_total"`
	Status      string `json:"status"`
}
