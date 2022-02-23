package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/image"
	"be/delivery/controllers/user"
	"be/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutesPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController, ic *image.ImageController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/login", ac.Login())
	e.POST("/user", uc.Create())
	e.GET("/user", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/user", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/user", uc.Delete(), middlewares.JwtMiddleware())

	//Image
	e.POST("/image", ic.Create())

}
