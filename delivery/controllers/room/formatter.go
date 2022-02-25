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
	City_id     int    `form:"city_id" json:"city_id" validate:"required"`
	Address     string `form:"address" json:"address" validate:"required"`
	Name        string `form:"name" json:"name" validate:"required"`
	Category    string `form:"category" json:"category" validate:"required"`
	Status      string `form:"status" json:"status" validate:"required"`
	Price       int    `form:"price" json:"price" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
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

type CreateImageRequesFormat struct {
	Room_uid string `form:"room_uid"`
	Url      string
}

type ImageInput struct {
	Url string `json:"url"`
}

type ImageReq struct {
	Array []ImageInput `json:"array"`
}
