package room

type Images struct {
	Url string
}

type RoomGetByIdResp struct {
	Room_uid    string   `json:"room_uid"`
	Name        string   `json:"name"`
	Image       []Images `json:"Image"`
	Address     string   `json:"address"`
	Owner_room  string   `json:"owner_room"`
	City        string   `json:"city"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Category    string   `json:"category"`
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
