package user

import "gorm.io/datatypes"

type UserCreateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RoomUserResp struct {
	Room_uid    string `json:"room_uid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Status      string `json:"status"`
}

type BookingUserResp struct {
	Booking_uid string         `json:"booking_uid"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Start_date  datatypes.Date `json:"start_date"`
	End_date    datatypes.Date `json:"end_date"`
	Price       int            `json:"price"`
	Days        int            `json:"days"`
	Price_total int            `json:"price_total"`
	Status      int            `json:"reservation"`
}

type GetByIdResponse struct {
	User_uid string            `json:"user_uid"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Rooms    []RoomUserResp    `json:"rooms"`
	Bookings []BookingUserResp `json:"bookings"`
}
