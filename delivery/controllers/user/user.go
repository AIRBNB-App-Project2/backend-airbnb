package user

import (
	"be/delivery/controllers/templates"
	"be/entities"
	"be/repository/database/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repo user.User) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := UserRequest{}

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for create new user", err))
		}
		v := validator.New()
		if err := v.Struct(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for create new user", nil))
		}
		res, err := uc.repo.Create(entities.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new user", err))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new user", res))
	}
}
