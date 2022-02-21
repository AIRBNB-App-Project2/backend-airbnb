package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/user"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func RoutesPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/user", uc.Create())
	e.POST("/login", ac.Login())
}
