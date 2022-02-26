package room

import (
	"be/entities"
	"be/repository/database/room"
)

type MockRoomRepo struct{}

func (repo *MockRoomRepo) Create(roomInput entities.Room) (room.RoomCreateResp, error) {
	return room.RoomCreateResp{}, nil
}

func (repo *MockRoomRepo) Update(user_uid string, room_uid string, upRoom entities.Room) (entities.Room, error) {
	return entities.Room{}, nil
}

func (repo *MockRoomRepo) GetById(room_uid string) (room.RoomGetByIdResp, error) {
	return room.RoomGetByIdResp{}, nil
}
func (repo *MockRoomRepo) Delete(room_uid string) (entities.Room, error) {
	return entities.Room{}, nil
}
func (repo *MockRoomRepo) GetAllRoom(length int, city, category, name, status string) ([]room.RoomGetAllResp, error) {
	return []room.RoomGetAllResp{}, nil
}

// func TestGetById(t *testing.T) {
// 	t.Run("success get by id", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/room")

// 		userController := New(&MockRoomRepo{})
// 		userController.GetById()(context)

// 		response := GetURoomByIdResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 200, response.Code)
// 	})
// }
