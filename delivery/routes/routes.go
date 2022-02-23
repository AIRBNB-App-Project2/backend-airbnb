package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/user"
	"be/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutesPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/login", ac.Login())
	e.POST("/user", uc.Create())
	e.GET("/user/:uid", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/user/:uid", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/user/:uid", uc.Delete(), middlewares.JwtMiddleware())
}
