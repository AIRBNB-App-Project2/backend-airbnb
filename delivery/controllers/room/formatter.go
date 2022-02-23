package room

import (
	"be/repository/database/room"
)

type GetURoomByIdResponseFormat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    room.RoomGetByIdResp `json:"data"`
}
