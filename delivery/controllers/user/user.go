package user

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
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
		user := UserCreateRequest{}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}
		v := validator.New()
		if err := v.Struct(user); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		res, err := uc.repo.Create(entities.User{Name: user.Name, Email: user.Email, Password: user.Password})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new user", err))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new user", res))
	}
}

func (uc *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {

		userUidParam := c.Param("uid")
		userUidToken := middlewares.ExtractTokenId(c)
		if userUidToken != userUidParam {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "sorry.. you are rejected", nil))
		}

		res, err := uc.repo.GetById(userUidParam)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "User not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get User", res))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUidParam := c.Param("uid")
		userUidToken := middlewares.ExtractTokenId(c)
		if userUidToken != userUidParam {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "sorry.. you are rejected", nil))
		}

		var newUser = UserUpdateRequest{}

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := uc.repo.Update(userUidParam, entities.User{Name: newUser.Name, Email: newUser.Email})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Update User", res))
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenId(c)

		if err := uc.repo.Delete(userUid); err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Delete User", nil))
	}
}
