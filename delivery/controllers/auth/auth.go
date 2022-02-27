package auth

import (
	"be/entities"
	"be/repository/database/auth"
	"net/http"

	"be/delivery/controllers/templates"
	"be/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		Userlogin := Userlogin{}

		if err := c.Bind(&Userlogin); err != nil || Userlogin.Email == "" || Userlogin.Password == "" {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for login user ", err))
		}

		checkedUser, err := ac.repo.Login(entities.User{Email: Userlogin.Email, Password: Userlogin.Password})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for login user "+err.Error(), err))
		}
		token, err := middlewares.GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, templates.BadRequest(http.StatusNotAcceptable, "error in process token "+err.Error(), err))
		}

		return c.JSON(http.StatusOK, templates.Success(nil, "success login", map[string]interface{}{
			"name":  checkedUser.Name,
			"token": token,
		}))
	}
}
