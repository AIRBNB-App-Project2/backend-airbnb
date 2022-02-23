package room

import (
	"be/repository/database/room"
)

type GetURoomByIdResponseFormat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    room.RoomGetByIdResp `json:"data"`
}
type CreateRoomResponseFormat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    room.RoomGetByIdResp `json:"data"`
}

type CreateRoomRequesFormat struct {
	User_uid    string
	City_id     int    `json:"city_id" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type UpdateRoomRequesFormat struct {
	City_id     int    `json:"city_id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
