package image

import (
	"be/delivery/controllers/templates"
	"be/entities"
	"image"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ImageController struct {
	repo image.Image
}

func New(repo image.Image) *ImageController {
	return &ImageController{
		repo: repo,
	}
}

func (ic *ImageController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// user := UserCreateRequest{}
		image := entities.Image{}

		if err := c.Bind(&image); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}
		v := validator.New()
		if err := v.Struct(image); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		res, err := ic.repo.Create(entities.Image{Url: image.Url})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new image", err))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new image", res))
	}
}
