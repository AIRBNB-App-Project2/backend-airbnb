package room

type Images struct {
	Url string
}

type BookingResp struct {
	Booking_uid string `json:"booking_uid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	Price       int    `json:"price"`
	Days        int    `json:"days"`
	Price_total int    `json:"price_total"`
	Status      int    `json:"reservation"`
}

type RoomGetByIdResp struct {
	Room_uid    string        `json:"room_uid"`
	Name        string        `json:"name"`
	Image       []Images      `json:"Image"`
	Address     string        `json:"address"`
	Owner_room  string        `json:"owner_room"`
	City        string        `json:"city"`
	Price       int           `json:"price"`
	Description string        `json:"description"`
	Status      string        `json:"status"`
	Category    string        `json:"category"`
	Bookings    []BookingResp `json:"bookings"`
}

type RoomCreateResp struct {
	Room_uid    string `json:"room_uid"`
	Name_user   string `json:"name_user"`
	Name_room   string `json:"name_roon"`
	Category    string `json:"category"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type RoomGetAllResp struct {
	Room_uid    string `json:"room_uid"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
