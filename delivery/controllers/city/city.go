package city

import (
	"be/delivery/controllers/templates"
	"be/repository/database/city"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	repo city.City
}

func New(repo city.City) *CityController {
	return &CityController{
		repo: repo,
	}
}

func (cont *CityController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cont.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error in get all city data " + err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(nil, "success get all city data", res))
	}
}
