package room

import (
	"be/delivery/controllers/templates"
	"be/repository/database/room"
	"net/http"

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
