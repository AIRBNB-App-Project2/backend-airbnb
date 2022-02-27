package room

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/entities"
	"be/repository/database/image"
	"be/repository/database/room"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin entities.User) (entities.User, error) {
	return entities.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

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
func (repo *MockRoomRepo) GetAllRoom(length int, city, category string) ([]room.RoomGetAllResp, error) {
	return []room.RoomGetAllResp{}, nil
}

type MockFailRoomRepo struct{}

func (repo *MockFailRoomRepo) Create(roomInput entities.Room) (room.RoomCreateResp, error) {
	return room.RoomCreateResp{}, errors.New("")
}

func (repo *MockFailRoomRepo) Update(user_uid string, room_uid string, upRoom entities.Room) (entities.Room, error) {
	return entities.Room{}, errors.New("")
}

func (repo *MockFailRoomRepo) GetById(room_uid string) (room.RoomGetByIdResp, error) {
	return room.RoomGetByIdResp{}, errors.New("")
}
func (repo *MockFailRoomRepo) Delete(room_uid string) (entities.Room, error) {
	return entities.Room{}, errors.New("")
}
func (repo *MockFailRoomRepo) GetAllRoom(length int, city, category string) ([]room.RoomGetAllResp, error) {
	return []room.RoomGetAllResp{}, errors.New("")
}

type MockImage struct{}

func (repo *MockImage) Create(room_uid string, image image.ImageReq) error {
	return nil
}

type MockFailImage struct{}

func (repo *MockFailImage) Create(room_uid string, image image.ImageReq) error {
	return errors.New("")
}

func TestGetById(t *testing.T) {
	t.Run("success get by id", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/room/:room_uid")

		controller := New(&MockRoomRepo{}, &MockImage{})
		controller.GetById()(context)

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/room/:room_uid")

		controller := New(&MockFailRoomRepo{}, &MockImage{})
		controller.GetById()(context)

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success get all", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/room")
		context.QueryParams().Add("city", "malang")
		context.QueryParams().Add("category", "standart")
		context.QueryParams().Add("length", "10")
		context.QueryParams()

		controller := New(&MockRoomRepo{}, &MockImage{})
		controller.GetAll()(context)

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
	})

	t.Run("error atoi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/room")
		context.QueryParams().Add("city", "malang")
		context.QueryParams().Add("category", "standart")
		context.QueryParams().Add("length", "nothing")
		context.QueryParams()

		controller := New(&MockRoomRepo{}, &MockImage{})
		controller.GetAll()(context)

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/room")
		context.QueryParams().Add("city", "malang")
		context.QueryParams().Add("category", "standart")
		context.QueryParams().Add("length", "10")
		context.QueryParams()

		controller := New(&MockFailRoomRepo{}, &MockImage{})
		controller.GetAll()(context)

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
	})
}

func TestDelete(t *testing.T) {

	jwtToken := ""
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.New(&MockAuthLib{})
		authControl.Login()(context)

		responses := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Data["token"].(string)
		fmt.Println(jwtToken)
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("success delete", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/room/:room_uid")

		controller := New(&MockRoomRepo{}, &MockImage{})
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(controller.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/room/:room_uid")

		controller := New(&MockFailRoomRepo{}, &MockImage{})
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(controller.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetURoomByIdResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
	})
}
