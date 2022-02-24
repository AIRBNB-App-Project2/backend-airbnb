package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/city"
	"be/delivery/controllers/image"
	"be/delivery/controllers/room"
	"be/delivery/controllers/user"
	"be/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutesPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController, ic *image.ImageController, cc *city.CityController, rc *room.RoomController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	// User ====================================

	e.POST("/login", ac.Login())
	e.POST("/user", uc.Create())

	g := e.Group("", middlewares.JwtMiddleware())
	g.GET("/user", uc.GetById())
	g.PUT("/user", uc.Update())
	g.DELETE("/user", uc.Delete())

	// Room =============================

	g.GET("/room", rc.GetAll())
	g.GET("/room/:room_uid", rc.GetById())
	g.POST("/room", rc.Create())
	g.PUT("/room/:room_uid", rc.Update())

	//Image
	e.POST("/image", ic.Create())

	//City
	g.GET("city", cc.GetAll())

}
