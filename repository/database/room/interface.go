package room

import "be/entities"

type Room interface {
	GetAllRoom(length int, city, category, name, status string) ([]RoomGetAllResp, error)
	Create(room entities.Room) (RoomCreateResp, error)
	Update(user_uid string, room_uid string, upRoom entities.Room) (entities.Room, error)
	GetById(room_uid string) (RoomGetByIdResp, error)
	Delete(room_uid string) (entities.Room, error)
}
