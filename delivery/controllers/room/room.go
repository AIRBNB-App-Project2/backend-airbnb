package room

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/entities"
	"be/repository/database/room"
	"net/http"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
)

type RoomController struct {
	repo room.Room
}

func New(repo room.Room) *RoomController {
	return &RoomController{
		repo: repo,
	}
}

func (cont *RoomController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		room_uid := c.Param("room_uid")

		res, err := cont.repo.GetById(room_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Room not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get Room", res))
	}
}
func (cont *RoomController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		s := c.QueryParam("s")
		city := c.QueryParam("s")
		category := c.QueryParam("s")
		name := c.QueryParam("s")
		length := c.QueryParam("s")
		status := c.QueryParam("s")

		res, err := cont.repo.GetAll(s, city, category, name, length, status)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Room not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get Room", res))
	}
}
func (cont *RoomController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		var room CreateRoomRequesFormat

		if err := c.Bind(&room); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		if err := v.Struct(room); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		room.User_uid = middlewares.ExtractTokenId(c)

		res, err := cont.repo.Create(entities.Room{User_uid: room.User_uid, City_id: room.City_id, Address: room.Address, Name: room.Name, Category: room.Category, Status: room.Status, Price: room.Price, Description: room.Description})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Room not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get Room", res))
	}
}
func (cont *RoomController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomParam := c.Param("room_uid")
		var room UpdateRoomRequesFormat

		if err := c.Bind(&room); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}
		user_uid := middlewares.ExtractTokenId(c)

		res, err := cont.repo.Update(user_uid, roomParam, entities.Room{City_id: room.City_id, Address: room.Address, Name: room.Name, Category: room.Category, Status: room.Status, Price: room.Price, Description: room.Description})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Room not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get Room", res))
	}
}